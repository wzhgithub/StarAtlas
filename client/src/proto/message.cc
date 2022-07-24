#include "message.h"
#include "common/utils.h"

TeleMessage::TeleMessage(uint8_t idx, uint8_t idx_exchange, const char* name) {
  m_tag = 0xeb;
  m_size = 0; // set later
  m_type = 0x55; // message type
  memset(m_name, 0, sizeof(m_name));
  m_index = 0;
  m_exchange_idx = 0;
  m_total_cpu = m_total_dsp = m_total_gpu = m_total_fpga = 0;
  m_total_mem = m_total_disk = 0;
  m_mem_rate = m_cpu_rate = m_dsp_rate = m_gpu_rate = m_disk_rate = 0;
  m_cnt_exchange = m_cnt_remote = 0;
  m_total_partition = 0;
}

TeleMessage::~TeleMessage() {
}

void TeleMessage::init(uint8_t idx, uint8_t idx_exchange, const char* name) {
  m_index = idx;
  m_exchange_idx = idx_exchange;
  if (!name) {
    snprintf(m_name, _LEN_VMC_NAME_, "vmc_%02d", (int)m_index);
  } else {
    snprintf(m_name, _LEN_VMC_NAME_, "%s", name);
  }
}

uint16_t TeleMessage::getSize() {
  m_size = 31; // init
  int n_total = cnt_remote + cnt_exchange + cnt_cpu + cnt_dsp + cnt_gpu + cnt_fpga, n_idx = 0;
  m_size += ((cnt_remote+cnt_exchange)*13 + (cnt_cpu+cnt_dsp)*21 + cnt_gpu*19 + cnt_fpga*12);
  if (cnt_remote) m_size+=3;
  if (cnt_exchange) m_size+=3;
  if (cnt_cpu) m_size+=3;
  if (cnt_dsp) m_size+=3;
  if (cnt_gpu) m_size+=3;
  if (cnt_fpga) m_size+=3;
  
  // partition
  m_size += 1;
  for (size_t h=0; h<m_partition.size(); h++) {
    m_size += (18+12*m_partition[h].m_total_task);
  } 
  m_size += 1; // crc
  return m_size;
}

int TeleMessage::pack(char* buf) {
  // calculate m_size
  m_size = getSize();
  char* p = buf;

  ((uint8_t*)p++)[0] = m_tag;
  ((uint16_t*)p)[0] = htons(m_size); p+=2;
  ((uint8_t*)p++)[0] = m_type;

  memcpy(p, m_name, _LEN_VMC_NAME_); 
  p+=_LEN_VMC_NAME_;
  ((uint8_t*)p++)[0] = m_index;

  ((uint8_t*)p++)[0] = m_total_cpu;
  ((uint8_t*)p++)[0] = m_total_dsp;
  ((uint8_t*)p++)[0] = m_total_gpu;
  ((uint8_t*)p++)[0] = m_total_fpga;

  ((uint8_t*)p++)[0] = m_exchange_idx;

  ((uint16_t*)p)[0] = htons(m_total_mem); p+=2;
  ((uint16_t*)p)[0] = htons(m_total_disk); p+=2;

  ((uint8_t*)p++)[0] = m_mem_rate;
  ((uint8_t*)p++)[0] = m_cpu_rate;
  ((uint8_t*)p++)[0] = m_dsp_rate;
  ((uint8_t*)p++)[0] = m_gpu_rate;
  ((uint8_t*)p++)[0] = m_disk_rate;

  ((uint8_t*)p++)[0] = m_cnt_exchange;
  ((uint8_t*)p++)[0] = m_cnt_remote;

  int total_dev = m_cnt_remote + m_cnt_exchange + m_total_cpu + m_total_dsp + m_total_gpu + m_total_fpga, i = 0;
  int anDevice[] = {
    m_cnt_remote,
    m_cnt_exchange,
    m_total_cpu,
    m_total_dsp,
    m_total_gpu,
    m_total_fpga,
  };
  int anSize[] = {
    13,
    13,
    21,
    21,
    19,
    12
  };

  for (int h=0; h<sizeof(anDevice)/sizeof(int); h++) {
    if (anDevice[h]==0) continue;
    ((uint16_t*)p)[0] = htons(m_devices[i].getTagHead()); p+=2;
    ((uint8_t*)p++)[0] = anSize[h]*anDevice[h];
    for (int j=0; j<anDevice[h]; i++,j++) {
      p+=m_devices[i].pack(p);
    }
  }

  ((uint8_t*)p++)[0] = m_total_partition;
  for (int i=0; i<m_total_block; i++) {
    p+=m_partition[i].pack(p);
  }
  ((uint8_t*)p++)[0] = crc_calculate(buf, m_size-1); //crc
  return (int)(p-buf);
}

void TeleMessage::reset() {
  m_size = 0;
  memset(m_name, 0, sizeof(m_name));
  m_index = 0;
  m_total_cpu = m_total_dsp = m_total_gpu = m_total_fpga = 0;
  m_exchange_idx = 0;
  m_total_mem = m_total_disk = 0;
  m_mem_rate = m_cpu_rate = m_dsp_rate = m_gpu_rate = m_disk_rate = 0;

  // keep switch & remote
  m_devices.erase(m_devices.begin()+(m_cnt_exchange+m_cnt_remote), m_devices.end());
}

bool TeleMessage::parseVmc(rapidjson::Document& _document) {
  if (!_document.HasMember("index") ||
      !_document.HasMember("connect_to") ||
      !_document.HasMember("name")) {
    cerr << "Invalid vmc. " << endl;
    return false;
  }
  
  init(uint8_t(_document["index"].GetInt()),
       uint8_t(_document["connect_to"].GetInt()),
       _document["name"].GetString().c_str());
  return true;
}

FaultMsg::FaultMsg() {
  m_action = m_taskType = m_idxPart = m_status =0;
  m_szTaskName[0] = m_szTaskName[1] = '\0';
  m_size = 6;
}

FaultMsg::~FaultMsg() {
}

int FaultMsg::pack(char* buf) {
  char* p = buf;
  ((uint8_t*)p++)[0] = m_action;
  memcpy(p, m_szTaskName, sizeof(m_szTaskName)); p+=sizeof(m_szTaskName);
  ((uint8_t*)p++)[0] = m_taskType;
  ((uint8_t*)p++)[0] = m_idxPart;
  ((uint8_t*)p++)[0] = m_status;
  return int(p-buf);
}

ControlMessage::ControlMessage() {
  m_tag = 0xeb;
  m_size = 7;
  m_size += m_msg.m_size; // 13 bytes;
  m_type = 0xAA;
}

ControlMessage::~ControlMessage() {
}

int ControlMessage::pack(char* buf) {
  char* p = buf;
  ((uint8_t*)p++)[0] = m_tag;
  ((uint16_t*)p)[0] = htons(m_size); p+=2;
  ((uint8_t*)p++)[0] = m_type;
  ((uint8_t*)p++)[0] = m_index;
  ((uint8_t*)p++)[0] = m_cpu_index;
  p+=m_msg.pack(p);

  m_crc = crc_calculate((uint8_t*)buf, m_size-1);
  ((uint8_t*)p++)[0] = m_crc;
  return int(p-buf);
}
