#include "telemsg.h"
#include <cassert>
#include <dirent.h>
#include <unistd.h>
#include <linux/limits.h>
#include <sys/stat.h>
#include <vector>

using namespace std;

const char* szConfBasePath = "conf/topology";
const char* szConfArray[] = {
  "switch.json",
  "remote.json",
  "tasks.json",
};
const char* szVmcPrefix = "vmc_";

size_t get_vmc_conf(const char* dir_name, vector<string>& arr) {
  size_t sz = arr.size();
  struct stat s;
  lstat(dir_name, &s);
  assert (S_ISDIR(s.st_mode));
  DIR* dir;
  dir = opendir(dir_name);

  struct dirent* filename;
  while ((filename = readdir(dir))!=NULL) {
    if(strncmp(filename->d_name, "vmc", 3)!=0) {
      continue;
    }
    string s(dir_name);
    arr.emplace_back(s + filename->d_name);
  }

  return arr.size()-sz;
}

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

  char dir[PATH_MAX] = {0};
  int n = readlink("/proc/self/exe", dir, PATH_MAX);
  char* p = dir, *plast = dir;
  for (; ;) {
    for (; p[0] && p[0]!='/'; p++) {
    }
    if (p[0]) plast = ++p;
    else break;
  }
  n = snprintf(plast, PATH_MAX-(plast-dir), "%s/%s/", szConfBasePath, curConf);
  p = plast + n;
  constexpr size_t sz = sizeof(szConfArray)/sizeof(const char*);
  vector<string> arrConf;
  for (size_t h=0; h<sz; h++) {
    snprintf(p, PATH_MAX-(p-dir), "%s", szConfArray[h]);
    arrConf.push_back(dir);
  }
  p[0] = '\0';
  get_vmc_conf(dir, arrConf);
  size_t h=0;
  for (; h<; h++) {
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
