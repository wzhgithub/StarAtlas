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

void Task::init(uint16_t idx, uint8_t typ, uint8_t status, uint8_t stt, const char* name) {
  m_index = idx;
  m_type = typ;
  m_status = status;
  m_start_time = stt;
  if (name) {
    snprintf(m_name, _LEN_TASK_NAME_, "%s", name);
  } else {
    snprintf(m_name, _LEN_TASK_NAME_, "%c", random()%26+'a');
  }
}

void Task::updateTask(uint32_t exe_time, uint8_t ret_code) {
  m_exe_time = exe_time;
  m_ret_code = ret_code;
}

int Task::pack(char *buf) {
  char* p = buf;
  memcpy(p, m_name, _LEN_TASK_NAME_); 
  p+=_LEN_TASK_NAME;
  ((uint16_t*)p)[0] = htons(m_index); p+=2;
  ((uint8_t*)p++)[0] = m_type;
  ((uint8_t*)p++)[0] = m_status;
  ((uint32_t*)p)[0] = htonl(m_exe_time); p+=4;
  ((uint8_t*)p++)[0] = m_ret_code;
  ((uint8_t*)p++)[0] = m_start_time;
  return (int)(p-buf);
}

Partition::Partition() {
      memset(m_name, 0, sizeof(m_name));
    m_total_task = 0;
    m_duration = 0;
    m_time = 0;
    m_index = 0;
    m_cnt_reset= 0;
    m_vmc_idx= 0;
    m_ptask = 0;

}

Partition::~Partition() {
}
