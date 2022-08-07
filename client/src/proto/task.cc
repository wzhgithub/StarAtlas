#include "task.h"

#include <string.h>
#include <vector>
#include <fstream>
#include <iostream>

#include <arpa/inet.h>

using std::vector;

Task::Task() {
  memset(m_name, 0, sizeof(m_name));
  m_index = 0;
  m_type = 0;
  m_status = 0;
  m_exe_time = 0;
  m_ret_code = 2;
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
  p+=_LEN_TASK_NAME_;
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
  m_time = 0; // time allocate
  m_index = 0;
  m_cnt_reset= 0;
  m_vmc_idx= 0;
}

Partition::~Partition() {
}

void Partition::init(int idx, int vmc_idx, const char* name) {
  m_index = (uint8_t)idx;
  m_vmc_idx = (uint8_t)vmc_idx;
  if (!name) {
    snprintf(m_name, _LEN_PARTITION_NAME_, "part_%02d", idx);
  } else {
    snprintf(m_name, _LEN_PARTITION_NAME_, "%s", name);
  }
}

int Partition::pack(char* buf) {
  char* p = buf;
  memcpy(p, m_name, _LEN_PARTITION_NAME_); 
  p+=_LEN_PARTITION_NAME_;
  ((uint8_t*)p++)[0] = m_total_task;
  ((uint16_t*)p)[0] = htons(m_duration); p+=2;
  ((uint16_t*)p)[0] = htons(m_time); p+=2;

  ((uint8_t*)p++)[0] = m_index;
  ((uint8_t*)p++)[0] = m_cnt_reset;
  ((uint8_t*)p++)[0] = m_vmc_idx;
  for (size_t i=0; i<m_total_task; i++) {
    p += m_tasks[i].pack(p);
  }
  return (int)(p-buf);
}

bool Partition::parseTask(const rapidjson::Value& _document) {
  if (!_document.HasMember("tasks")) {
    std::cerr << "Empty tasks." << std::endl;
    return false;
  }

  const auto& _tasks = _document["tasks"].GetArray();
  for (const auto& _task: _tasks) {
    if (!_task.IsObject()) {
      std::cerr << "Invalid task: !_task.IsObject() " << std::endl;
      m_tasks.clear();
      return false;
    }

    if (!_task.HasMember("name") ||
        !_task.HasMember("index") ||
        !_task.HasMember("type") ||
        !_task.HasMember("status") ||
        !_task.HasMember("start_time")) {
      std::cerr << "Invalid task: <name|index|type|status|start_time> needed." << std::endl;
      m_tasks.clear();
      return false;
    }

    Task oTask;
    oTask.init(
      uint16_t(_task["index"].GetInt()),
      uint8_t(_task["type"].GetInt()),
      uint8_t(_task["status"].GetInt()),
      uint8_t(_task["start_time"].GetInt()),
      _task["name"].GetString());

    m_tasks.emplace_back(std::move(oTask));
  }
  m_total_task = (uint8_t)m_tasks.size();
  return true;
}
