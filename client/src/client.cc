#include <iostream>
#include <string>
#include <cstring>
#include <random>
#include <time.h>
#include <arpa/inet.h>
#include <fstream>

using namespace std;

typedef enum {
  eCPU = 0x0,
  eDSP,
  eGPU,
  eFPGA,
  eInvalid
} EDeviceType;

uint16_t g_device_tag[] = {
  0xeba0, 0xebaa, // CPU
  0xebb0, 0xebbb, // DSP 
  0xebc0, 0xebcc, // GPU
  0xebd0, 0xebdd, // FPGA
  0x0, 0x0
};

char gaszDevNameFmt[][10] = {
  "CPU_%02d",
  "DSP_%02d",
  "GPU_%02d",
  "FPGA_%02d",
  ""
};

uint8_t gDevType[] = {
  4,
  2,
  0,
  0,
  0
};

class Device {
public:
  uint8_t m_dev_type;

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
  uint8_t m_xpu_rate;

public:
  Device() {
    m_dev_type = eInvalid;
    m_tag_head = m_tag_tail = 0;
    m_device_name[0] = '\0';
    m_device_index = 0;
    m_device_type= 0;
  }

  int pack(char* buf)  {
    char* p = buf;
    ((uint16_t*)p)[0] = htons(m_tag_head);
    p+=2;

    memcpy(p, m_device_name, 10);
    p+=10;
    ((uint8_t*)p++)[0] = m_device_index;
    ((uint8_t*)p++)[0] = m_device_type;

    // cpu && dsp && gpu
    if (m_dev_type!=eFPGA) {
      ((uint8_t*)p++)[0] = m_cnt_core;
      if (m_dev_type!=eGPU) {
        ((uint16_t*)p)[0] = htons(m_iops);
	p+=2;
      }
      ((uint16_t*)p)[0] = htons(m_flops);
      p+=2;
      ((uint16_t*)p)[0] = htons(m_mem);
      p+=2;
    ((uint8_t*)p++)[0] = m_mem_rate;
    ((uint8_t*)p++)[0] = m_xpu_rate;
    }

    ((uint16_t*)p)[0] = htons(m_tag_tail);
    p+=2;

    //cout<<"name:"<<m_device_name<<"len: "<<(int)(p-buf)<<endl;
    return (int)(p-buf);
  }
};

class Task {
public:
  char m_name[10];
  uint8_t m_type;
  uint8_t m_status;
  uint8_t m_time;

public:
  int pack(char* buf) {
    char* p = buf;
    memcpy(p, m_name, 10);
    p+=10;
    ((uint8_t*)p++)[0] = m_type;
    ((uint8_t*)p++)[0] = m_status;
    ((uint8_t*)p++)[0] = m_time;
    return (int)(p-buf);
  }
};

class Block {
public:
  char m_name[10];
  uint8_t m_total_task;
  uint16_t m_duration;
  uint16_t m_time;
  uint8_t m_cur_task;

  Task* m_ptask; // 

public:
  Block() {
    m_name[0] = '\0';
    m_total_task = 0;
    m_duration = 0;
    m_time = 0;
    m_cur_task = 0;
    m_ptask = 0;
  }
  virtual ~Block() {
    delete []m_ptask;
  }

public:
  void ReSet(int idx, int max_task) {
    snprintf(m_name, 10, "part_%02d", idx);
    m_total_task = random()%max_task+1;
    //cout<<"total task1:"<<(int)m_total_task<<endl;
    m_duration = 250;
    m_time = 100;

    //cout<<"total task2:"<<(int)m_total_task<<endl;
    m_ptask = new Task[m_total_task];
    for (int i=1; i<=m_total_task; i++) {
      Task& tsk = m_ptask[i-1];

      snprintf(tsk.m_name, 10, "task_%02d", i);
      tsk.m_type = random()%2;

      const uint8_t task_status[] = {
	0, 1, 2, 3, 0xff
      }; 
      tsk.m_status = task_status[ random()%sizeof(task_status) ];

      const uint8_t task_stime[] = {
	0, 10, 40, 60, 80, 90, 
      }; 
      tsk.m_time = i-1<sizeof(task_stime)?task_stime[i-1]:100;
    } 
    m_cur_task = random()%m_total_task+1;
  }

  int pack(char* buf) {
    char* p = buf;
    memcpy(p, m_name, 10);
    p+=10;

    ((uint8_t*)p++)[0] = m_total_task;
    ((uint16_t*)p)[0] = htons(m_duration);
    p+=2;
    ((uint16_t*)p)[0] = htons(m_time);
    p+=2;
    for (int i=0; i<m_total_task; i++) {
      p += m_ptask[i].pack(p);
    }
    ((uint8_t*)p++)[0] = m_cur_task;
    return (int)(p-buf);
  }
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
  uint8_t m_disk_rate;

  // device related
  Device* m_pdevices;
  
  // block
  uint8_t m_total_block;
  Block* m_pblock;


public:
  void SetXPU(Device& rd, uint8_t typ, int idx) {
    rd.m_dev_type = typ;
    rd.m_tag_head  = g_device_tag[typ*2];
    rd.m_tag_tail  = g_device_tag[typ*2+1];

    snprintf(rd.m_device_name, 10, gaszDevNameFmt[typ], idx);
    rd.m_device_index = idx;
    rd.m_device_type = gDevType[typ];

    rd.m_cnt_core = random()%255;
    rd.m_iops = random()%65535;
    rd.m_flops = random()%65535;
    rd.m_mem= random()%65535;
    rd.m_mem_rate = random()%100;
    rd.m_xpu_rate = random()%100;
  }

  Message(
    int idx, // vmc index
    int idx_exchange,
    int cnt_cpu,
    int cnt_dsp,
    int cnt_gpu,
    int cnt_fpga,
    int cnt_block,
    int cnt_max_task
  ) {
    m_tag = 0xeb;
    m_size = 30; // init
    m_type = 0x55;
    m_index = idx;
    snprintf(m_name, 10, "vmc_%02d", (int)m_index);

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

    int n_total = cnt_cpu + cnt_dsp + cnt_gpu + cnt_fpga, n_idx = 0;
    m_size += ((cnt_cpu+cnt_dsp)*25 + cnt_gpu*23 + cnt_fpga*16);

    // device
    //cout<<"total device: "<<(int)n_total<<endl;
    m_pdevices = new Device[n_total];
    for (int i=0; i<cnt_cpu; i++, n_idx++) {
      Device& rd = m_pdevices[n_idx];
      SetXPU(rd, eCPU, i);
    }

    for (int i=0; i<cnt_dsp; i++, n_idx++) {
      Device& rd = m_pdevices[n_idx];
      SetXPU(rd, eDSP, i);
    }

    for (int i=0; i<cnt_gpu; i++, n_idx++) {
      Device& rd = m_pdevices[n_idx];
      SetXPU(rd, eGPU, i);
    }

    for (int i=0; i<cnt_fpga; i++, n_idx++) {
      Device& rd = m_pdevices[n_idx];
      SetXPU(rd, eFPGA, i);
    }

    // block
    //cout<<"cnt block: "<<(int)cnt_block<<endl;
    m_size++; // count of block
    m_total_block = cnt_block;
    //cout<<"!total block: "<<(int)cnt_block<<endl;
    //cout<<"!!total block: "<<(int)m_total_block<<endl;
    m_pblock = new Block[m_total_block];
    for (int i=0; i<m_total_block; i++) {
      m_pblock[i].ReSet(i, cnt_max_task);

      m_size += (16+13*m_pblock[i].m_total_task);
    }

    m_size++;
  }

  virtual ~Message() {
    delete []m_pdevices;
    delete []m_pblock;
  }

  int pack(char* buf) {
    char* p = buf;

    ((uint8_t*)p++)[0] = m_tag;
    ((uint16_t*)p)[0] = htons(m_size);
    p+=2;
    ((uint8_t*)p++)[0] = m_type;

    memcpy(p, m_name, 10);
    p+=10;
    ((uint8_t*)p++)[0] = m_index;

    ((uint8_t*)p++)[0] = m_total_cpu;
    ((uint8_t*)p++)[0] = m_total_dsp;
    ((uint8_t*)p++)[0] = m_total_gpu;
    ((uint8_t*)p++)[0] = m_total_fpga;
    ((uint8_t*)p++)[0] = m_total_reserved;

    ((uint8_t*)p++)[0] = m_exchange_idx;

    ((uint16_t*)p)[0] = htons(m_total_mem);
    p+=2;
    ((uint16_t*)p)[0] = htons(m_total_disk);
    p+=2;

    ((uint8_t*)p++)[0] = m_mem_rate;
    ((uint8_t*)p++)[0] = m_cpu_rate;
    ((uint8_t*)p++)[0] = m_dsp_rate;
    ((uint8_t*)p++)[0] = m_gpu_rate;
    ((uint8_t*)p++)[0] = m_disk_rate;

    //cout<<"p head: "<<(int)(p-buf)<<endl;
    int total_dev = m_total_cpu + m_total_dsp + m_total_gpu + m_total_fpga;
    for (int i=0; i<total_dev; i++) {
      p+=m_pdevices[i].pack(p);
      //cout<<"p dev "<<i<<":"<<(int)(p-buf)<<endl;
    }

    ((uint8_t*)p++)[0] = m_total_block;
    for (int i=0; i<m_total_block; i++) {
      p+=m_pblock[i].pack(p);
      //cout<<"p block "<<i<<":"<<(int)(p-buf)<<endl;
    }

    ((uint8_t*)p++)[0] = 0; //crc
    //cout<<"p crc: "<<(int)(p-buf)<<endl;
    return (int)(p-buf);
  }
};

int main(int argc, char* argv[]) {
  if (argc!=2) {
    cerr<<"Usage: "<<argv[0]<<" /dev/stderr"<<endl;
    exit(0);
  }
  srand((unsigned)time(NULL));
  int idx = random()%10,
      idx_exch = random()%5,
      cnt_cpu = random()%10,
      cnt_dsp = random()%10,
      cnt_gpu = random()%10,
      cnt_fpga = random()%10,
      cnt_block = random()%6,
      cnt_max_task = random()%6;
  cnt_max_task==0?cnt_max_task=1:0;

  /*
  cout<<"idx: "<<idx<<"\n"
      <<"idx_exch: "<<idx_exch<<"\n"
      <<"cnt_cpu: "<<cnt_cpu<<"\n"
      <<"cnt_dsp: "<<cnt_dsp<<"\n"
      <<"cnt_gpu: "<<cnt_gpu<<"\n"
      <<"cnt_fpga: "<<cnt_fpga<<"\n"
      <<"cnt_block: "<<cnt_block<<"\n"
      <<"cnt_max_task: "<<cnt_max_task<<endl;
      */

  Message msg(
    idx,
    idx_exch,
    cnt_cpu,
    cnt_dsp,
    cnt_gpu,
    cnt_fpga,
    cnt_block,
    cnt_max_task
  );

  //cout<<"msg.m_size:"<<(int)msg.m_size<<endl;
  char* buf = new char[msg.m_size+16];
  memset(buf, 0, msg.m_size+16);
  int sz_out = msg.pack(buf);
  //cout<<"sz_out:"<<(int)sz_out<<endl;
  ofstream binaryio(argv[1], ios::binary);
  if (!binaryio) {
    cerr<<"open "<<argv[1]<<" failed."<<endl;
    exit(0);
  }
  binaryio.write(buf,msg.m_size);
  delete [] buf;
}
