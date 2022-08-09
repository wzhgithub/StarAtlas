#ifndef _PROTO_TASK_INCLUDED_
#define _PROTO_TASK_INCLUDED_

#include <stdint.h>
#include <vector>

#include "rapidjson/schema.h"
#include "rapidjson/stringbuffer.h"

using std::vector;

constexpr int _LEN_TASK_NAME_ = 2;
constexpr int _LEN_PARTITION_NAME_ = 10;

class Partition;
class Task {
public:
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

public:
  friend class Partition;
};

class Partition {
private:
  char m_name[_LEN_PARTITION_NAME_];
  uint8_t m_total_task;
  uint16_t m_duration;
  uint16_t m_time;

  uint8_t m_index;
  uint8_t m_cnt_reset;
  uint8_t m_vmc_idx;

  vector<Task> m_tasks;

public:
  uint8_t task_count() const {return m_total_task;}

public:
  Partition();
  virtual ~Partition();

public:
  void init(int idx, int vmc_idx, const char* name=nullptr);
  int pack(char* buf);

public:
  bool parseTask(const rapidjson::Value& _document);
  void updateTask(uint16_t _interval = 100);
};

#endif
