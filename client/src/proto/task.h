#ifndef _PROTO_TASK_INCLUDED_
#define _PROTO_TASK_INCLUDED_

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
public:
# define _LEN_PARTITION_NAME_ 10
  char m_name[_LEN_PARTITION_NAME_];
  uint8_t m_total_task;
  uint16_t m_duration;
  uint16_t m_time;

  uint8_t m_index;
  uint8_t m_cnt_reset;
  uint8_t m_vmc_idx;

  vector<Task> m_task;

public:
  Partition();
  virtual ~Partition();

public:
  void ReSet(int idx, int max_task, int idx_vmc) {
    m_index = (uint8_t)idx;
    snprintf(m_name, 10, "part_%02d", idx);

    const uint8_t N_TASK_COUNT = 6;
    m_total_task = random()%max_task+1; // 
    if (m_total_task==0) {
      m_total_task = N_TASK_COUNT;
    }
    m_duration = 250;
    m_time = 100;
    m_cnt_reset = 0;
    m_vmc_idx  = idx_vmc;
    m_ptask = new Task[m_total_task];
    for (int i=1; i<=m_total_task; i++) {
      Task& tsk = m_ptask[i-1];

      snprintf(tsk.m_name, 2, "%d", i);
      tsk.m_index = uint16_t(i);
      tsk.m_type = random()%2;

      const uint8_t task_status[] = {
	0, 1, 2, 3
      }; 
      tsk.m_status = task_status[ random()%sizeof(task_status) ];
      tsk.m_exe_time = 10240;
      tsk.m_ret_code = random()%3;
      const uint8_t task_stime[] = {
	0, 10, 40, 60, 80, 90, 
      }; 
      tsk.m_start_time = i-1<sizeof(task_stime)?task_stime[i-1]:100;
    } 
  }

  int pack(char* buf) {
    char* p = buf;
    memcpy(p, m_name, 10); p+=10;
    ((uint8_t*)p++)[0] = m_total_task;
    ((uint16_t*)p)[0] = htons(m_duration); p+=2;
    ((uint16_t*)p)[0] = htons(m_time); p+=2;

    ((uint8_t*)p++)[0] = m_index;
    ((uint8_t*)p++)[0] = m_cnt_reset;
    ((uint8_t*)p++)[0] = m_vmc_idx;
    for (int i=0; i<m_total_task; i++) {
      p += m_ptask[i].pack(p);
    }
    return (int)(p-buf);
  }
};


#endif
