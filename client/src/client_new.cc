#include "proto/message.h"
#include "proto/parser.h"
#include "common/utils.h"

#include <fstream>
#include <functional>
#include <iostream>
#include <map>
#include <string>

#include <sys/socket.h>
#include <sys/types.h>
#include <unistd.h>
#include <arpa/inet.h>

using std::cerr;
using std::cout;
using std::endl;
using std::map;
using std::string;
using std::function;

const char*  szConfBasePath = "conf/topology";
const char* pszTaskConf = "tasks.json";

typedef function<int (rapidjson::Document&, vector<Device>&)> ParseFunc;

class DeviceParser {
public:
  string m_fname;
  int m_type;
  ParseFunc m_func;

public:
  DeviceParser(const char* _name, int _typ, ParseFunc _parse_func = parseDefault):
    m_fname(_name), m_type(_typ), m_func(_parse_func) {
  }
};

const DeviceParser gdev_parser[] {
  DeviceParser("remote.json", eREMOTE, parseRemote),
  DeviceParser("switch.json", eEXCHNAGE, parseSwitch),
};

const char szVmcPrefix[] = "vmc";
constexpr size_t nConf = sizeof(gdev_parser)/sizeof(DeviceParser);

int main(int argc, char* argv[]) {
  if (argc==1) {
    fprintf(stderr, "Usage: %s [conf:random|demo|fault|parallel] [ip:port] [loop] [interval] [dump 4 debug]\n"
                    "  etc: %s random\n", argv[0], argv[0]);
    exit(0);
  }

  const char* curConf = "random";
  if (argc>1) {
    curConf = argv[1];
  }

  const char* ip = "127.0.0.1";
  unsigned short port = 9191;
# define LEN_BUF 128
  char szBuf[LEN_BUF] = {0};
  if (argc>2) {
    size_t len = strlen(argv[2]); 
    if (len>=LEN_BUF) {
      cerr<<"Invalid ip address: "<<argv[2]<<endl;
      exit(-1);
    }
    strncpy(szBuf, argv[2], LEN_BUF);
    char* p = szBuf;
    for (; p[0] && p[0]!=':'; p++) {}
    if (!p[0]) {
      cerr<<"Invalid ip address: "<<argv[2]<<endl;
      exit(-1);
    }
    *p++='\0';
    ip = szBuf;
    port = atoi(p);
  }
  cout<<"ip: "<<ip<<"; port:"<<port<<"."<<endl;

  map<string, string> _path_map;
  char dir[PATH_MAX] = {0};
  char *p = nullptr, *plast = get_cur_dir(dir, PATH_MAX);
  int n = snprintf(plast, PATH_MAX-(plast-dir), "%s/%s/", szConfBasePath, curConf);
  p = plast + n;
  string sCurPath(dir);
  for (size_t h=0; h<nConf; h++) {
    const string& _name = gdev_parser[h].m_fname;
    snprintf(p, PATH_MAX-(p-dir), "%s", _name.c_str());
    _path_map[_name] = dir;
  }
  // partition
  snprintf(p, PATH_MAX-(p-dir), "%s", pszTaskConf);
  string _filename_task(dir);
  p[0] = '\0';

  vector<string> _vmc;
  get_vmc_conf(dir, _vmc, szVmcPrefix, strlen(szVmcPrefix));

  bool _bflag = false;
  //cout << "hello: " << _vmc.size() << endl;

  vector<TeleMessage> _msg_arr;
  srand((unsigned)time(NULL));
  for (size_t h=0; h<_vmc.size(); h++) {
    TeleMessage _msg;
    cout << "parse vmc: " << h << "; name: " << _vmc[h] << endl;

    rapidjson::Document _doc;
    if (!parse(_vmc[h].c_str(), _doc) ||
        !_msg.parseVmc(_doc)) { // init
      exit(-1);
    }

    // parse remote & switch & task
    for (size_t hh=0; hh<nConf; hh++) {
      const string& _name = gdev_parser[hh].m_fname;
      int _typ = gdev_parser[hh].m_type;
      ParseFunc _func = gdev_parser[hh].m_func;
      auto it = _path_map.find(_name);
      rapidjson::Document _ddoc;
      if (it==_path_map.end() || 
          !parse(it->second.c_str(), _ddoc)) {
        cerr << "parse "<< _name << " failed." << endl;
        exit(-1);
      }

      int _n_dev = _func(_ddoc, _msg.getDevice());
      if (!_n_dev) {
        cerr << "warning: empty device, device type: "<< _typ <<endl;
      }
      _msg.setTotalDevice(_n_dev, _typ);
    }
    uint8_t nGlobalDev = _msg.getDevice().size();

    rapidjson::Document _xpu_doc;
    string _filename_xpu = (_doc.HasMember("xpu")?
      sCurPath + string(_doc["xpu"].GetString()):
    "");
    if (!parse(_filename_xpu.c_str(), _xpu_doc)) {
      cerr << "parse xpu failed." << endl;
      exit(-1);
    }

    const char* _xpu_name[] = {
      "cpu",
      "dsp",
      "gpu",
      "fpga",
    };
    for (int hh=0; hh<sizeof(_xpu_name)/sizeof(const char*); hh++) {
      if (!_xpu_doc.HasMember(_xpu_name[hh])) continue;
      int _n_dev = parseXpu(_xpu_doc[_xpu_name[hh]], _msg.getDevice(), hh, _msg.getBaseIndex(), nGlobalDev);
      if (!_n_dev) {
        cerr << "warning: empty device, device type: "<< hh <<endl;
      }
      _msg.setTotalDevice(_n_dev, hh);
    }

    // parse partition
    string _local_filename_task = (_doc.HasMember("tasks")?
      sCurPath + string(_doc["tasks"].GetString()):
    _filename_task);
    rapidjson::Document _tdoc;
    if (!parse(_local_filename_task.c_str(), _tdoc)) {
      exit(-1);
    }
    int _n_part = parsePartition(_tdoc, _msg.getPartition());
    if (!_n_part) {
      cerr << "warning: empty tasks." <<endl;
    }
    _msg.setTotalPartition(_n_part);

    _msg_arr.emplace_back(_msg); 
  }

  // send
  int sockfd;
  socklen_t len = sizeof(struct sockaddr_in);

  struct sockaddr_in serveraddr;
  memset(&serveraddr, 0, len);

  serveraddr.sin_family = AF_INET;
  serveraddr.sin_port = htons(port);
  serveraddr.sin_addr.s_addr = inet_addr(ip);

  sockfd = socket(AF_INET, SOCK_DGRAM, 0);
  if (sockfd < 0)
  {
    cerr << "fail to create socker" << endl;
    exit(-1);
  } 

  size_t _buff_size = 1024;
  char* _buff = new char[_buff_size];

  size_t max_loop = 1;
  if (argc>3) {
    max_loop = atoi(argv[3]);
  }
  int interval = 500;
  if (argc>4) {
    interval = atoi(argv[4]);
  }

  for (int _loop = 0; _loop < max_loop; _loop++) {
    for (size_t h=0; h<_msg_arr.size(); h++) {
      TeleMessage& _msg = _msg_arr[h];
       _msg.updateRandom();

       size_t sz = _msg.getSize();
       if (sz>_buff_size) {
         delete []_buff;
         _buff_size *= 2;
         _buff = new char[_buff_size];
       }
       memset(_buff, 0, _buff_size);
       int sz_out = _msg.pack(_buff);

       ssize_t ret = sendto(sockfd, _buff, sz_out, 0, (struct sockaddr*)&serveraddr, len);
       if (ret == -1) {
         cerr << "sendto failed, ret code: "<<ret<<endl;
         close(sockfd);
         delete [] _buff;
         exit(-1);
       }
       cout << "send message, total byte:"<<ret<<"."<<endl;

       char output[PATH_MAX] = {0};
       const char* _prefix_dump = "sample";
       if (argc>5) {
         _prefix_dump = argv[5];
       }
       snprintf(output, PATH_MAX, "%s_%d_%d.bin", _prefix_dump, h, _loop);
       ofstream binaryio(output, ios::binary);
       if (!binaryio) {
         cerr<<"open "<<output<<" failed."<<endl;
         exit(0);
       }
       binaryio.write(_buff,sz_out);
    }
    usleep(interval*1000);
  }
  delete []_buff;
  close(sockfd);
  exit(0);
}
