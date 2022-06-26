

class Package {
private:
  // proto
  uint_8 m_flag;   // head: 1
  ushort_16 m_len; // len:  2
  uint_8 m_type;   //       1
  
  // 
  char m_mvc_name[10]; //   10
  uint_8 m_idx;    //       1

  // device statistic
  uint_8 m_cpu_cnt;//       1
  uint_8 m_dsp_cnt;//       1
  uint_8 m_gpu_cnt;//       1
  uint_8 m_fpga_cnt;//      1
  uint_8 m_reserver; //     1
  uint_8 m_idx_changer;//   1
  ushort_16 m_total_mem;//  2
  ushort_16 m_total_disk; //2
  uint_8 m_mem_rate;      //1
  uint_8 m_total_cpu_rate;//1
  uint_8 m_total_dsp_rate;//1
  uint_8 m_total_gpu_rate;//1

  // cpu == equal dsp
  ushort_16 m_cpu_flag;   // 2 -> 0xeba0

  char m_cpu_name[10];          // 10
  uint8_t m_cpu_idx;            // 1
  uint8_t m_cpu_type;           // 1
  uint8_t m_cnt_cpu_core;       // 1
  ushort16_t m_int_cal_power;   // 2
  ushort16_t m_float_cal_power; // 2
  ushort16_t m_mem;             // 2
  ushort16_t m_mem_rate;        // 1
  ushort16_t m_cpu_rate;        // 1  -> sum: 21
  
  ushort_16 m_cpu_tail;   // 2 -> 0xebaa
  
  // dsp
  ushort_16 m_dsp_flag;   // 2 -> 0xebb0
  char m_dsp_name[10];
  uint8_t m_dsp_idx;
  uint8_t m_dsp_type;
  uint8_t m_cnt_dsp_core;
  ushort16_t m_int_cal_power;   // 2
  ushort16_t m_float_cal_power; // 2
  ushort16_t m_mem;             // 2
  ushort16_t m_mem_rate;        // 1
  ushort16_t m_cpu_rate;        // 1  -> sum: 21

  ushort_16 m_dsp_tail;   // 2 -> 0xebbb


  // gpu
  ushort_16 m_gpu_flag;   // 2 -> 0xebc0
  char m_gpu_name[10];
  uint8_t m_gpu_idx;
  uint8_t m_gpu_type;
  uint8_t m_cnt_gpu_core;
  ushort16_t m_cal_power; // 2
  ushort16_t m_mem;             // 2
  ushort16_t m_mem_rate;        // 1
  ushort16_t m_cpu_rate;        // 1  -> sum: 19
  ushort_16 m_gpu_tail;   // 2 -> 0xebcc
  
  // fpga
  ushort_16 m_fpga_flag;   // 2 -> 0xebc0
  char m_fpga_name[10];
  uint8_t m_fpga_idx;
  uint8_t m_fpga_type;     // 1 -> sum:12
  ushort_16 m_fpga_tail;   // 2 -> 0xebcc

  // task
  uint8_t m_cnt_block; // 1
  char m_block_name[10];  // 10
  uint8_t m_cnt_task;     // 1
  ushort16_t m_duration_task; // 2
  ushort16_t m_time_block; // 2

  // loop
  // task 1...n
  char m_task_name[10];  13*6
  uint8_t m_task_type;
  uint8_t m_task_status;
  uint8_t m_task_time;
  
  // task end
  // loop end
};
