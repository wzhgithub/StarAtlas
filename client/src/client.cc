#include "telemsg.h"

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

  TeleMessage msg(
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
