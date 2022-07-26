#include "proto/message.h"
#include "proto/parser.h"
#include "common/utils.h"

#include <fstream>
#include <functional>
#include <iostream>
#include <map>
#include <string>

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
    fprintf(stderr, "Usage: %s [conf:random|demo|fault|parallel]\n"
                    "  etc: %s random\n", argv[0], argv[0]);
    exit(0);
  }

  const char* curConf = "random";
  if (argc>1) {
    curConf = argv[1];
  }

  map<string, string> _path_map;
  char dir[PATH_MAX] = {0};
  char *p = nullptr, *plast = get_cur_dir(dir, PATH_MAX);
  int n = snprintf(plast, PATH_MAX-(plast-dir), "%s/%s/", szConfBasePath, curConf);
  p = plast + n;
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
  get_vmc_conf(dir, _vmc, szVmcPrefix, sizeof(szVmcPrefix));

  bool _bflag = false;
  for (size_t h=0; h<_vmc.size(); h++) {
    TeleMessage _msg;

    rapidjson::Document _doc;
    if (!parse(_vmc[h].c_str(), _doc) ||
        !_msg.parseVmc(_doc)) { // init
      exit(-1);
    }

    // parse remote & switch & task
    for (size_t h=0; h<nConf; h++) {
      const string& _name = gdev_parser[h].m_fname;
      int _typ = gdev_parser[h].m_type;
      ParseFunc _func = gdev_parser[h].m_func;
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

    const char* _xpu_name[] = {
      "cpu",
      "dsp",
      "gpu",
      "fpga",
    };
    for (int h=0; h<sizeof(_xpu_name)/sizeof(const char*); h++) {
      if (!_doc.HasMember(_xpu_name[h])) continue;
      int _n_dev = parseXpu(_doc[_xpu_name[h]], _msg.getDevice(), h);
      if (!_n_dev) {
        cerr << "warning: empty device, device type: "<< h <<endl;
      }
      _msg.setTotalDevice(_n_dev, h);
    }

    // parse partition
    rapidjson::Document _tdoc;
    if (!parse(_filename_task.c_str(), _tdoc)) {
      exit(-1);
    }
    int _n_part = parsePartition(_tdoc, _msg.getPartition());
    if (!_n_part) {
      cerr << "warning: empty tasks." <<endl;
    }
    _msg.setTotalPartition(_n_part);

    // 
    size_t sz = _msg.getSize();
    char* buf = new char[sz+16];
    memset(buf, 0, sz+16);
    int sz_out = _msg.pack(buf);
    ofstream binaryio(argv[2], ios::binary);
    if (!binaryio) {
      cerr<<"open "<<argv[2]<<" failed."<<endl;
      exit(0);
    }
    binaryio.write(buf,sz);
    delete [] buf;
  }
  exit(0);
}
