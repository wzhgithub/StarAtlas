#include "device.h"
#include "rapidjson/schema.h"
#include "rapidjson/stringbuffer.h"

#include <random>

#include <arpa/inet.h>

Device::Device() {
  m_device_type = eInvalid;
  m_tag_head = m_tag_tail = 0;
  memset(m_device_name, 0, sizeof(m_device_name));
  m_device_index = 0;
  m_device_subtype= 0;
  m_connect_to = 0;

  m_cnt_core = 0;
  m_inops = m_flops = 0;
  m_mem = 0;
  m_mem_rate = 0;
  m_xpu_rate = 0;
}

Device::~Device() {
}


bool Device::parseSwitch(const char* filename) {
}

bool Device::parseRemote(const char* filename) {
}

void Device::init(uint8_t typ, int idx, const char* name, 
  uint8_t sub_type, int connect_to) {
    m_tag_head  = g_device_tag[typ*2];
    m_tag_tail  = g_device_tag[typ*2+1];

    m_device_type = typ;
    m_device_index = idx;

    if (!name) {
      snprintf(m_device_name, _LEN_DEV_NAME_, gaszDevNameFmt[typ], idx);
    } else {
      snprintf(m_device_name, _LEN_DEV_NAME_, "%s", name);
    }

    m_device_subtype = random()%(gDevType[typ]+1);
    if (sub_type!=0xFF) m_device_subtype = sub_type;

    // remote & exchanger
    m_connect_to = connect_to;
    if (typ==eEXCHNAGE) {
      if (idx==0) {
        m_device_subtype = 0;
        m_connect_to = 0;
      } else {
        m_device_subtype = 1;
        m_connect_to = 0;
      }
    }
    if (typ==eREMOTE) {
      m_connect_to = 1;
    }
}

void Device::setBasic(uint8_t core, uint16_t inops, uint16_t flops, uint16_t mem) {
  m_cnt_core = core;
  m_inops = inops;
  m_flops = flops;
  m_mem = mem;
}

void Device::updateRate(uint8_t mem_rate, uint8_t xpu_rate) {
  m_mem_rate = mem_rate;
  m_xpu_rate = xpu_rate;
}

int Device::pack(char* buf) {
  char* p = buf;

  memcpy(p, m_device_name, _LEN_DEV_NAME_);
  p+=_LEN_DEV_NAME_;
  ((uint8_t*)p++)[0] = m_device_index;
  ((uint8_t*)p++)[0] = m_device_subtype;
  if (m_device_type==eEXCHNAGE || m_device_type==eREMOTE) {
    ((uint8_t*)p++)[0] = m_connect_to;
    return int(p-buf);
  }

  // cpu && dsp && gpu
  if (m_device_type!=eFPGA) {
    ((uint8_t*)p++)[0] = m_cnt_core;
    if (m_device_type!=eGPU) {
      ((uint16_t*)p)[0] = htons(m_inops); p+=2;
    }
    ((uint16_t*)p)[0] = htons(m_flops); p+=2;
    ((uint16_t*)p)[0] = htons(m_mem); p+=2;
    ((uint8_t*)p++)[0] = m_mem_rate;
    ((uint8_t*)p++)[0] = m_xpu_rate;
  }
  return (int)(p-buf);
}
