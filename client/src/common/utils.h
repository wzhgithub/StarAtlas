#ifndef _COMMON_UTILS_INCLUDED_
#define _COMMON_UTILS_INCLUDED_

#include <cassert>
#include <string>
#include <vector>

#include <dirent.h>
#include <linux/limits.h>
#include <sys/stat.h>
#include <unistd.h>


using std::string;
using std::vector;

char* get_cur_dir(char* buf, size_t sz);
size_t get_vmc_conf(const char* dir_name, vector<string>& arr);
uint8_t crc_calculate(uint8_t* p, size_t sz); 

#endif
