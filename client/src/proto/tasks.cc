#include "tasks.h"

#include <string.h>
#include <vector>

#include <arpa/inet.h>

using std::vector;

Task::Task() {
  memset(m_name, 0, sizeof(m_name));
  m_index = 0;
  m_type = 0;
  m_status = 0;
  m_exe_time = 0;
  m_ret_code = 0;
  m_start_time = 0;
}

Task::~Task() {
}

void Task::init(uint16_t idx, uint8_t typ, ) {
}

int Task::pack(char *buf) {
  char* p = buf;
  memcpy(p, m_name, _LEN_TASK_NAME); 
  p+=_LEN_TASK_NAME;
  ((uint16_t*)p)[0] = htons(m_index); p+=2;
  ((uint8_t*)p++)[0] = m_type;
  ((uint8_t*)p++)[0] = m_status;
  ((uint32_t*)p)[0] = htonl(m_exe_time); p+=4;
  ((uint8_t*)p++)[0] = m_ret_code;
  ((uint8_t*)p++)[0] = m_start_time;
  return (int)(p-buf);
}
