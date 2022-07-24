#ifndef _PROTO_TASK_INCLUDED_
#define _PROTO_TASK_INCLUDED_

#include "rapidjson/schema.h"
#include "rapidjson/stringbuffer.h"

class Task {
public:
# define _LEN_TASK_NAME_ 2
  char m_name[_LEN_TASK_NAME_];

  uint16_t m_index;
  uint8_t m_type;
  uint8_t m_status;
  uint32_t m_exe_time;
  uint8_t m_ret_code;
  uint8_t m_start_time;

public:
  Task();
  virtual ~Task();

public:
  void init(uint16_t idx, uint8_t typ, uint8_t status, uint8_t stt, const char* name=nullptr);
  void updateTask(uint32_t exe_time, uint8_t ret_code);
  int pack(char* buf);
};

class Partition {
private:
# define _LEN_PARTITION_NAME_ 10
  char m_name[_LEN_PARTITION_NAME_];
  uint8_t m_total_task;
  uint16_t m_duration;
  uint16_t m_time;

  uint8_t m_index;
  uint8_t m_cnt_reset;
  uint8_t m_vmc_idx;

  vector<Task> m_tasks;

public:
  Partition();
  virtual ~Partition();

public:
  void init(int idx, const char* name=nullptr);
  int pack(char* buf);

public:
  bool parseTask(rapidjson::Document& _document);
  void updateTask();
  void reset();
};

#endif
