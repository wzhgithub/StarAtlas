#ifndef _PROTO_DEVICE_INCLUDED_
#define _PROTO_DEVICE_INCLUDED_

#include <arpa/inet.h>

typedef enum {
  eCPU = 0x0,
  eDSP,
  eGPU,
  eFPGA,
  eREMOTE,
  eEXCHNAGE,
  eInvalid
} EDeviceType;

uint16_t g_device_tag[] = {
  0xeba0, 0xebaa, // CPU
  0xebb0, 0xebbb, // DSP 
  0xebc0, 0xebcc, // GPU
  0xebd0, 0xebdd, // FPGA
  0xebe0, 0xeedd, // remote
  0xebf0, 0xefdd, // exchange
  0x0, 0x0
};

char gaszDevNameFmt[][10] = {
  "CPU_%04d",
  "DSP_%04d",
  "GPU_%04d",
  "FPGA_%04d",
  "REMO_%04d",
  "EXCH_%04d",
  ""
};

class Device {
private:
  uint8_t m_dev_type;

  // all device
  uint16_t m_tag_head;
  uint16_t m_tag_tail;

  char m_device_name[10];
  uint8_t m_device_index;
  uint8_t m_device_type;
  uint8_t m_connect_to; // exchange && remote
  
  // cpu & dsp & gpu
  uint8_t m_cnt_core;
  uint16_t m_iops;   // only cpu & dsp
  uint16_t m_flops;
  uint16_t m_mem;
  uint8_t m_mem_rate;
  uint8_t m_xpu_rate;

public:
  Device();
  virtual ~Device();

public:
  bool parseSwitch(const char* filename);
  bool parseRemote(const char* filename);

  void set(uint8_t typ, int idx, const char* name=nullptr, 
    uint8_t sub_type =0xFF, int connect_to = 0);
  void update();
  int pack(char* buf);  
};

#endif
