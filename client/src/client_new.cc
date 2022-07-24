#include "proto/message.h"
#include "proto/parser.h"
#include "common/utils.h"

#include <functional>
#include <map>
#include <string>

using std::map;
using std::string;

const char* szConfBasePath = "conf/topology";
const char* szConfArray[] = {
  "switch.json",
  "remote.json",
  "tasks.json",
};
const char* szVmcPrefix = "vmc_";
constexpr size_t nConf = sizeof(szConfArray)/sizeof(const char*);

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
    snprintf(p, PATH_MAX-(p-dir), "%s", szConfArray[h]);
    _path_map[szConfArray[h]] = dir;
  }
  p[0] = '\0';

  vector<string> _vmc;
  get_vmc_conf(dir, _vmc);

  bool _bflag = false;
  for (size_t h=0; h<_vmc.size(); h++) {
    TeleMessage _msg;

    rapidjson::Document _doc;
    if (!parse(_vmc[h].c_str(), _doc) ||
        !_msg.parseVmc(_doc)) { // init
      exit(-1);
    }
    
    
  }
  exit(0);

  /*
  TeleMessage msg(
    idx,
    idx_exch,
    cnt_cpu,
    cnt_dsp,
    cnt_gpu,
    cnt_fpga,
    cnt_exchange,
    cnt_remote,
    cnt_block,
    cnt_max_task
  );

  char* buf = new char[msg.m_size+16];
  memset(buf, 0, msg.m_size+16);
  int sz_out = msg.pack(buf);
  ofstream binaryio(argv[1], ios::binary);
  if (!binaryio) {
    cerr<<"open "<<argv[1]<<" failed."<<endl;
    exit(0);
  }
  binaryio.write(buf,msg.m_size);
  delete [] buf;
  */
}
