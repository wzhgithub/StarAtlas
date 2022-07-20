#include "device.h"
#include "rapidjson/schema.h"
#include "rapidjson/stringbuffer.h"

#include <random>

Device::Device() {
  m_dev_type = eInvalid;
  m_tag_head = m_tag_tail = 0;
  memset(m_device_name, 0, sizeof(m_device_name));
  m_device_index = 0;
  m_device_type= 0;
  m_connect_to = 0;
}

Device::~Device() {
}


bool Device::parseSwitch(const char* filename) {
}

bool Device::parseRemote(const char* filename) {
}

void Device::set(uint8_t typ, int idx, const char* name, 
  uint8_t sub_type, int connect_to) {
    m_tag_head  = g_device_tag[typ*2];
    m_tag_tail  = g_device_tag[typ*2+1];

    m_dev_type = typ;
    m_device_index = idx;

    if (!name) {
      snprintf(m_device_name, 10, gaszDevNameFmt[typ], idx);
    } else {
      snprintf(m_device_name, 10, "%s", name);
    }

    m_device_type = random()%(gDevType[typ]+1);
    if (sub_type!=0xFF) m_device_type = sub_type;

    // remote & exchanger
    rd.m_connect_to = connect_to;
    if (typ==eEXCHNAGE) {
      if (idx==0) {
        rd.m_device_type = 0;
        rd.m_connect_to = 0;
      } else {
        rd.m_device_type = 1;
        rd.m_connect_to = 0;
      }
    }
    if (typ==eREMOTE) {
      rd.m_connect_to = 1;
    }

    rd.m_cnt_core = random()%255;
    rd.m_iops = random()%65535;
    rd.m_flops = random()%65535;
    rd.m_mem= random()%65535;
    rd.m_mem_rate = random()%100;
    rd.m_xpu_rate = random()%100;
}



int Device::pack(char* buf) {
  char* p = buf;

  memcpy(p, m_device_name, 10); p+=10;
  ((uint8_t*)p++)[0] = m_device_index;
  ((uint8_t*)p++)[0] = m_device_type;
  if (m_dev_type==eEXCHNAGE || m_dev_type==eREMOTE) {
    ((uint8_t*)p++)[0] = m_connect_to;
    return int(p-buf);
  }

  // cpu && dsp && gpu
  if (m_dev_type!=eFPGA) {
    ((uint8_t*)p++)[0] = m_cnt_core;
    if (m_dev_type!=eGPU) {
      ((uint16_t*)p)[0] = htons(m_iops); p+=2;
    }
    ((uint16_t*)p)[0] = htons(m_flops); p+=2;
    ((uint16_t*)p)[0] = htons(m_mem); p+=2;
    ((uint8_t*)p++)[0] = m_mem_rate;
    ((uint8_t*)p++)[0] = m_xpu_rate;
  }
  return (int)(p-buf);
}
