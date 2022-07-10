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
  "CPU_%02d",
  "DSP_%02d",
  "GPU_%02d",
  "FPGA_%02d",
  "REMO_%02d",
  "EXCH_%02d",
  ""
};

uint8_t gDevType[] = {
  4,
  2,
  0,
  0,
  2,
  1,
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
  uint8_t m_connect_to; // exchange && remote
  
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
    memset(m_device_name, 0, sizeof(m_device_name));
    m_device_index = 0;
    m_device_type= 0;
    m_connect_to = 0;
  }

  int pack(char* buf)  {
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
};

class Task {
public:
  //char m_name[10];
  char m_name[2]; // ???

  uint16_t m_index;
  uint8_t m_type;
  uint8_t m_status;
  uint32_t m_exe_time;
  uint8_t m_ret_code;
  uint8_t m_start_time;

public:
  Task() {
    m_name[0] = m_name[1] = '\0';
    m_index = 0;
    m_status = 0;
    m_exe_time = 0;
    m_ret_code = 0;
    m_start_time = 0;
  }

public:
  int pack(char* buf) {
    char* p = buf;
    memcpy(p, m_name, 2); p+=2;
    ((uint16_t*)p)[0] = htons(m_index); p+=2;
    ((uint8_t*)p++)[0] = m_type;
    ((uint8_t*)p++)[0] = m_status;
    ((uint32_t*)p)[0] = htonl(m_exe_time); p+=4;
    ((uint8_t*)p++)[0] = m_ret_code;
    ((uint8_t*)p++)[0] = m_start_time;
    return (int)(p-buf);
  }
};

class Block {
public:
  char m_name[10];
  uint8_t m_total_task;
  uint16_t m_duration;
  uint16_t m_time;

  uint8_t m_index;
  uint8_t m_cnt_reset;
  uint8_t m_vmc_idx;

  Task* m_ptask; // 

public:
  Block() {
    m_name[0] = '\0';
    memset(m_name, 0, sizeof(m_name));
    m_total_task = 0;
    m_duration = 0;
    m_time = 0;
    m_index = 0;
    m_cnt_reset= 0;
    m_vmc_idx= 0;
    m_ptask = 0;
  }
  virtual ~Block() {
    delete []m_ptask;
  }

public:
  void ReSet(int idx, int max_task, int idx_vmc) {
    m_index = (uint8_t)idx;
    snprintf(m_name, 10, "part_%02d", idx);

    const uint8_t N_TASK_COUNT = 6;
    m_total_task = random()%max_task+1; // 
    if (m_total_task==0) {
      m_total_task = N_TASK_COUNT;
    }
    m_duration = 250;
    m_time = 100;
    m_cnt_reset = 0;
    m_vmc_idx  = idx_vmc;
    m_ptask = new Task[m_total_task];
    for (int i=1; i<=m_total_task; i++) {
      Task& tsk = m_ptask[i-1];

      snprintf(tsk.m_name, 2, "%d", i);
      tsk.m_index = uint16_t(i);
      tsk.m_type = random()%2;

      const uint8_t task_status[] = {
	0, 1, 2, 3, 0xff
      }; 
      tsk.m_status = task_status[ random()%sizeof(task_status) ];
      tsk.m_exe_time = 10240;
      tsk.m_ret_code = random()%3;
      const uint8_t task_stime[] = {
	0, 10, 40, 60, 80, 90, 
      }; 
      tsk.m_start_time = i-1<sizeof(task_stime)?task_stime[i-1]:100;
    } 
  }

  int pack(char* buf) {
    char* p = buf;
    memcpy(p, m_name, 10); p+=10;
    ((uint8_t*)p++)[0] = m_total_task;
    ((uint16_t*)p)[0] = htons(m_duration); p+=2;
    ((uint16_t*)p)[0] = htons(m_time); p+=2;

    ((uint8_t*)p++)[0] = m_index;
    ((uint8_t*)p++)[0] = m_cnt_reset;
    ((uint8_t*)p++)[0] = m_vmc_idx;
    for (int i=0; i<m_total_task; i++) {
      p += m_ptask[i].pack(p);
    }
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

  // 0707
  uint8_t m_cnt_exchange;
  uint8_t m_cnt_remote;

  // device related
  Device* m_pdevices;
  
  // block
  uint8_t m_total_block;
  Block* m_pblock;

public:
  void SetXPU(Device& rd, uint8_t typ, int idx, int connect_to=0) {
    rd.m_dev_type = typ;
    rd.m_tag_head  = g_device_tag[typ*2];
    rd.m_tag_tail  = g_device_tag[typ*2+1];

    snprintf(rd.m_device_name, 10, gaszDevNameFmt[typ], idx);
    rd.m_device_index = idx;
    rd.m_device_type = random()%(gDevType[typ]+1);
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

  Message(
    int idx, // vmc index
    int idx_exchange,
    int cnt_cpu,
    int cnt_dsp,
    int cnt_gpu,
    int cnt_fpga,
    int cnt_exchange,
    int cnt_remote,
    int cnt_block,
    int cnt_max_task
  ) {
    m_tag = 0xeb;
    m_size = 31; // remove reserved & 20220707, add exchange, remote
    m_type = 0x55;
    m_index = idx;
    memset(m_name, 0, sizeof(m_name));
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

  virtual ~Message() {
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

int main(int argc, char* argv[]) {
  if (argc==1) {
    //cerr<<"Usage: "<<argv[0]<<" /dev/stderr"<<endl;
    fprintf(stderr, "Usage: %s <vmc_index> <_exchange_idx>\n", argv[0]);
    exit(0);
  }

  // default
  int idx = 0,
      idx_exch = 4,
      cnt_remote = 4,
      cnt_exchange = 2;
  if (argc>=3) idx = atoi(argv[2]);
  if (argc>=4) idx_exch = atoi(argv[3]);

  srand((unsigned)time(NULL));
  //int idx = random()%10,
  //    idx_exch = random()%5,
  int  cnt_cpu = random()%10+1,
      cnt_dsp = random()%10+1,
      cnt_gpu = random()%10+1,
      cnt_fpga = random()%10+1,
      cnt_block = random()%6+1,
      cnt_max_task = random()%7;

      //cnt_remote = random()%10,
      //cnt_exchange= random()%10;
  cnt_max_task==0?cnt_max_task=1:0;

  /*cout<<"idx: "<<idx<<"\n"
      <<"idx_exch: "<<idx_exch<<"\n"
      <<"cnt_cpu: "<<cnt_cpu<<"\n"
      <<"cnt_dsp: "<<cnt_dsp<<"\n"
      <<"cnt_gpu: "<<cnt_gpu<<"\n"
      <<"cnt_fpga: "<<cnt_fpga<<"\n"
      <<"cnt_block: "<<cnt_block<<"\n"
      <<"cnt_remote: "<<cnt_remote<<"\n"
      <<"cnt_exchange: "<<cnt_exchange<<"\n"
      <<"cnt_max_task: "<<cnt_max_task<<endl;*/

  Message msg(
    idx,
    idx_exch,
    cnt_cpu,
    cnt_dsp,
    cnt_gpu,
    cnt_fpga,
    cnt_exchange,
    cnt_remote,
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
