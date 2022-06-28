#include <iostream>
#include <string>

#include <arpa/inet.h> 

using namespace std;

tyepdef enum {
  eCPU = 0x0,
  eDSP,
  eGPU,
  eFPGA
} EDeviceType;

uint16_t g_device_tag[] = {
  0xeba0, 0xebaa, // CPU
  0xebb0, 0xebbb, // DSP 
  0xebc0, 0xebcc, // GPU
  0xebd0, 0xebdd, // FPGA
  0x0, 0x0
};

class Device {
private:
  uint8_t m_device_type;

  // all device
  uint16_t m_tag_head;
  uint16_t m_tag_tail;

  char m_device_name[10];
  uint8_t m_device_index;
  uint8_t m_device_type;
  
  // cpu & dsp & gpu
  uint8_t m_cnt_core;
  uint16_t m_iops;   // only cpu & dsp
  uint16_t m_flops;
  uint16_t m_mem;
  uint8_t m_mem_rate;
  uint8_t m_cpu_rate;

public:
  Device(uint8_t type, uint8_t );


public:
  size_t pack(char* buf)  {
    char* p = buf;
    *((uint16_t*))p=htons( g_device_tag[m_device_type*2] );
    p+=2;

    switch( m_device_type) {
    case eCPU:
    case eDSP:
      {
        ;
        break;
      }
    }
    *((uint16_t*))p=htons( g_device_tag[m_device_type*2+1] );
    p+=2;
    return (p-buf);
  }
};

class Task {
public:
  char m_name[10];
  uint8_t m_type;
  uint8_t m_status;
  uint8_t m_time;

public:
  int pack(char* buf, size_t sz) {
    return 0;
  }
};

class Block {
public:
  char m_name[10];
  uint8_t m_total_task;
  uint16_t m_duration;
  uint16_t m_time;

  Task* m_ptask; // 
};


class Message {
public:
  uint8_t m_tag; //0xeb
  uint16_t m_size;
  uint8_t m_type;

  char m_name[10];
  uint8_t m_index;

  // device statistic
  uint8_t m_total_cpu;
  uint8_t m_total_dsp;
  uint8_t m_total_gpu;
  uint8_t m_total_fpga;
  uint8_t m_total_reserved;

  // 
  uint8_t m_exchange_idx;
  
  // performence
  uint16_t m_total_mem;
  uint16_t m_total_disk;

  uint8_t m_mem_rate;
  uint8_t m_cpu_rate;
  uint8_t m_dsp_rate;
  uint8_t m_gpu_rate;

  // device related
  Device* m_pdevices;
  
  // block
  uint8_t m_total_block;
  Block* m_pblock;

public:
};
