#ifndef _TELEMETERING_INCLUDED_
#define _TELEMETERING_INCLUDED_

#include "message.h"

TeleMessage::TeleMessage(uint8_t idx, uint8_t idx_exchange, const char* name) {
  m_tag = 0xeb;
  m_size = 0; // set later
  m_type = 0x55; // message type
  m_index = idx;
  memset(m_name, 0, sizeof(m_name));
  if (!name) {
    snprintf(m_name, _LEN_VMC_NAME_, "vmc_%02d", (int)m_index);
  } else {
    snprintf(m_name, _LEN_VMC_NAME_, "%s", name);
  }
  m_total_cpu = m_total_dsp = m_total_gpu = m_total_fpga = 0;
  m_exchange_idx = idx_exchange;
  m_total_mem = m_total_disk = 0;
  m_mem_rate = m_cpu_rate = m_dsp_rate = m_gpu_rate = m_disk_rate = 0;
  m_cnt_exchange = m_cnt_remote = 0;
}

int idx, int idx_exchange, int cnt_cpu,
    int cnt_exchange,
    int cnt_remote,
    int cnt_block,
    int cnt_max_task

    m_total_cpu = cnt_cpu;
    m_total_dsp = cnt_dsp;
    m_total_gpu = cnt_gpu;
    m_total_fpga = cnt_fpga;
    m_total_reserved = 0;

    m_exchange_idx = idx_exchange;

    m_total_mem = random()%65535;
    m_total_disk = random()%65535;

    m_mem_rate = random()%100;
    m_cpu_rate = random()%100;
    m_dsp_rate = random()%100;
    m_gpu_rate = random()%100;
    m_disk_rate = random()%100;

    m_cnt_exchange = cnt_exchange;
    m_cnt_remote = cnt_remote;

    // 20220707

    int n_total = cnt_remote + cnt_exchange + cnt_cpu + cnt_dsp + cnt_gpu + cnt_fpga, n_idx = 0;
    m_size += ((cnt_remote+cnt_exchange)*13 + (cnt_cpu+cnt_dsp)*21 + cnt_gpu*19 + cnt_fpga*12);
    if (cnt_remote) m_size+=3;
    if (cnt_exchange) m_size+=3;
    if (cnt_cpu) m_size+=3;
    if (cnt_dsp) m_size+=3;
    if (cnt_gpu) m_size+=3;
    if (cnt_fpga) m_size+=3;

    // device
    //cout<<"total device: "<<(int)n_total<<endl;
    m_pdevices = new Device[n_total];
    int anDevice[] = {
      cnt_remote,
      cnt_exchange,
      cnt_cpu,
      cnt_dsp,
      cnt_gpu,
      cnt_fpga,
    };

    uint8_t aDevType[] = {
      eREMOTE,
      eEXCHNAGE,
      eCPU,
      eDSP,
      eGPU,
      eFPGA,
    };

    for (int h=0; h<sizeof(anDevice)/sizeof(int); h++) {
      for (int i=0; i<anDevice[h]; i++, n_idx++) {
        Device& rd = m_pdevices[n_idx];

        int n_conn = 0;
        if (aDevType[h]==eREMOTE || aDevType[h]==eEXCHNAGE) n_conn = random()%256;
        SetXPU(rd, aDevType[h], i, n_conn);
      }
    }

    // block
    m_size++; // count of block
    m_total_block = cnt_block;
    m_pblock = new Block[m_total_block];
    for (int i=0; i<m_total_block; i++) {
      m_pblock[i].ReSet(i, cnt_max_task, m_index);

      m_size += (18+12*m_pblock[i].m_total_task);
    }
    m_size++;
    //cout<<"m_size: "<<m_size<<endl;
  }

  virtual ~TeleMessage() {
    delete []m_pdevices;
    delete []m_pblock;
  }

  int pack(char* buf) {
    char* p = buf;

    ((uint8_t*)p++)[0] = m_tag;
    ((uint16_t*)p)[0] = htons(m_size); p+=2;
    ((uint8_t*)p++)[0] = m_type;

    memcpy(p, m_name, 10); p+=10;
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

    //cout<<"remote: "<<m_cnt_remote<<endl;
    //cout<<"exchange: "<<m_cnt_exchange<<endl;
    for (int h=0; h<sizeof(anDevice)/sizeof(int); h++) {
      if (anDevice[h]==0) continue;
      //cout<<"cur offset: "<< (int)(p-buf) <<endl;

      ((uint16_t*)p)[0] = htons(g_device_tag[m_pdevices[i].m_dev_type*2]); p+=2;
      //cout<<"h: "<<h<<"; device size: "<<anSize[h]*anDevice[h]<<"; count:"<<anDevice[h]<<"; pack:"<<anSize[h]<<endl;
      ((uint8_t*)p++)[0] = anSize[h]*anDevice[h];
      //cout<<"h: "<<h<<"; device size: "<<anSize[h]*anDevice[h]<<"; count:"<<anDevice[h]<<"; pack:"<<anSize[h]<<endl;

      for (int j=0; j<anDevice[h]; i++,j++) {
        p+=m_pdevices[i].pack(p);
      }
      //cout<<"cur offset: "<< (int)(p-buf) <<endl;
    }

    ((uint8_t*)p++)[0] = m_total_block;
    for (int i=0; i<m_total_block; i++) {
      //cout<<"block start: "<< (int)(p-buf) <<"; "<<int(m_pblock[i].m_total_task)<<endl;
      p+=m_pblock[i].pack(p);
      //cout<<"block end: "<<int(p-buf)<<endl;
    }

    ((uint8_t*)p++)[0] = 0; //crc
    //cout<<"m_cur_size: "<<int(p-buf)<<endl;
    return (int)(p-buf);
  }
};

// control message
class FaultMsg {
public:
  uint8_t m_action;
  uint8_t m_taskType;
  uint8_t m_idxPart;
  uint8_t m_status; //  4 bytes
  char m_szTaskName[2]; // 2 bytes
  uint16_t m_size;

public:
  FaultMsg() {
    m_action = m_taskType = m_idxPart = m_status =0;
    m_szTaskName[0] = m_szTaskName[1] = '\0';
    m_size = 6;
  }
  int pack(char* buf) {
    char* p = buf;
    ((uint8_t*)p++)[0] = m_action;
    memcpy(p, m_szTaskName, sizeof(m_szTaskName)); p+=sizeof(m_szTaskName);
    ((uint8_t*)p++)[0] = m_taskType;
    ((uint8_t*)p++)[0] = m_idxPart;
    ((uint8_t*)p++)[0] = m_status;
    return int(p-buf);
  }
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
  ControlMessage() {
    m_tag = 0xeb;
    m_size = 7;
    m_size += m_msg.m_size; // 13 bytes;
    m_type = 0xAA;
  }

  int pack(char* buf) {
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
};


#endif
