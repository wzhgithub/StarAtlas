#ifndef _PROTO_DEVICE_INCLUDED_
#define _PROTO_DEVICE_INCLUDED_

#include <stdint.h>

typedef enum {
  eCPU = 0x0,
  eDSP,
  eGPU,
  eFPGA,
  eREMOTE,
  eEXCHNAGE,
  eInvalid
} EDeviceType;

static uint16_t g_device_tag[] = {
  0xeba0, 0xebaa, // CPU
  0xebb0, 0xebbb, // DSP 
  0xebc0, 0xebcc, // GPU
  0xebd0, 0xebdd, // FPGA
  0xebe0, 0xeedd, // remote
  0xebf0, 0xefdd, // exchange
  0x0, 0x0
};

static char gaszDevNameFmt[][10] = {
  "CPU_%04d",
  "DSP_%04d",
  "GPU_%04d",
  "FPGA_%04d",
  "REMO_%04d",
  "EXCH_%04d",
  ""
};

static uint8_t gDevType[] = {
  4, // ARM:0,RISC_V:1,SPARC:2,PPC:3,MIPS:4
  2, // 6701:0,6678:1,8024:2
  0, // NVIDIA AGX:0
  0, // FPGA, reserved
  2, // sensor：0；executor：1；load：2
  1, // central：0；access：1
};

class Device {
private:
  uint8_t m_device_type;

  // all device
  uint16_t m_tag_head;
  uint16_t m_tag_tail;

# define _LEN_DEV_NAME_ 10
  char m_device_name[_LEN_DEV_NAME_];
  uint8_t m_device_index;
  uint8_t m_device_subtype;
  uint8_t m_connect_to; // exchange && remote
  
  // cpu & dsp & gpu
  uint8_t m_cnt_core;
  uint16_t m_inops;   // only cpu & dsp
  uint16_t m_flops;
  uint16_t m_mem;
  uint8_t m_mem_rate;
  uint8_t m_xpu_rate;

public:
  Device();
  virtual ~Device();

public:
  uint16_t getTagHead() const { return m_tag_head; }
  uint16_t getTagTail() const { return m_tag_tail; }

public:
  void init(uint8_t typ, int idx, const char* name=nullptr, 
    uint8_t sub_type =0xFF, int connect_to = 0);
  void setBasic(uint8_t core, uint16_t m_inops, uint16_t m_flops, uint16_t mem); // fix propery?

  void updateRate(uint8_t mem_rate, uint8_t xpu_rate);
  int pack(char* buf);  
};

#endif
