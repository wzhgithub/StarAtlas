#include "utils.h"

#include <cstring>

char* get_cur_dir(char* dir, size_t sz) {
  int n = readlink("/proc/self/exe", dir, sz);
  char* p = dir, *plast = dir;
  for (; ;) {
    for (; p[0] && p[0]!='/'; p++) {
    }
    if (p[0]) plast = ++p;
    else break;
  }
  return plast;
}

size_t get_vmc_conf(const char* dir_name, vector<string>& arr, const char* _prefix, size_t _len) {
  size_t sz = arr.size();
  struct stat s;
  lstat(dir_name, &s);
  assert (S_ISDIR(s.st_mode));
  DIR* dir;
  dir = opendir(dir_name);

  struct dirent* filename;
  while ((filename = readdir(dir))!=NULL) {
    if(strncmp(filename->d_name, _prefix, _len)!=0) {
      continue;
    }
    string s(dir_name);
    arr.emplace_back(s + filename->d_name);
  }

  return arr.size()-sz;
}

uint8_t crc_calculate(uint8_t* p, size_t sz) {
  uint8_t value = 0;
  for (size_t h=0; h<sz; h++) {
    value += p[h];
  }
  return ~value;
}
