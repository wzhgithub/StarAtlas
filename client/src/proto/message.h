#ifndef _TELEMETERING_INCLUDED_
#define _TELEMETERING_INCLUDED_

#include "device.h"
#include "task.h"

#include <memory>

using namespace std;

class TeleMessage {
public:
  uint8_t m_tag; //0xeb
  uint16_t m_size;
  uint8_t m_type;

# define _LEN_VMC_NAME_ 10
  char m_name[_LEN_VMC_NAME_];
  uint8_t m_index;

  // device statistic
  uint8_t m_total_cpu;
  uint8_t m_total_dsp;
  uint8_t m_total_gpu;
  uint8_t m_total_fpga;

  // 
  uint8_t m_exchange_idx;
  
  // performence
  uint16_t m_total_mem;
  uint16_t m_total_disk;

  uint8_t m_mem_rate;
  uint8_t m_cpu_rate;
  uint8_t m_dsp_rate;
  uint8_t m_gpu_rate;
  uint8_t m_disk_rate;

  // 0707
  uint8_t m_cnt_exchange;
  uint8_t m_cnt_remote;

  // device related
  vector<Device> m_devices;
  
  // block
  std::weak_ptr< vector<Partition> > m_partitions;

  // patch 08/28/2022
  uint32_t m_runtime;
  uint8_t m_time_unit;

public:
  TeleMessage();
  ~TeleMessage();

public:
  uint8_t getSwitchIndex() const { return m_exchange_idx; }
  uint8_t getBaseIndex() const { return m_index; }
  vector<Device>& getDevice() { return m_devices; }

  void setTotalDevice (uint8_t _n, int typ) {
    switch (typ) {
      case eREMOTE: m_cnt_remote = _n; break;
      case eEXCHNAGE: m_cnt_exchange= _n; break;
      case eCPU: m_total_cpu = _n; break;
      case eDSP: m_total_dsp= _n; break;
      case eGPU: m_total_gpu= _n; break;
      case eFPGA: m_total_fpga= _n; break;
      default: break;
    }
  }
  void setPartition(std::shared_ptr< vector<Partition> > _parts) { m_partitions = _parts; }

public:
  void init(uint8_t idx, uint8_t idx_exchange, uint16_t mem, uint16_t disk, const char* name = nullptr);
  uint16_t getSize(); 
  int pack(char* buf);  

public:
  void reset();
  bool parseVmc(rapidjson::Document& _document);

public:
  void updateRandom();
};

// control message
class FaultMsg {
private:
  uint8_t m_action;
  uint8_t m_taskType;
  uint8_t m_idxPart;
  uint8_t m_status; //  4 bytes
  char m_szTaskName[2]; // 2 bytes
  uint16_t m_size;

public:
  FaultMsg();
  virtual ~FaultMsg();

public:
  uint16_t size() const { return m_size; }

public:
  int pack(char* buf);
};

class ControlMessage{
private:
  uint8_t m_tag;
  uint16_t m_size;
  uint8_t m_type;  // 4 bytes
  uint8_t m_index;
  uint8_t m_cpu_index;
  uint8_t m_crc; // 7 bytes

  FaultMsg m_msg;
public:
  ControlMessage();
  virtual ~ControlMessage();

public:
  int pack(char* buf);
};
#endif
