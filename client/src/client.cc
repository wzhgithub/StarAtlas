#include "telemsg.h"
#include <sys/socket.h>
#include <sys/types.h>
#include <unistd.h>
#include <arpa/inet.h>

int main(int argc, char* argv[]) {
  if (argc==1) {
    //cerr<<"Usage: "<<argv[0]<<" /dev/stderr"<<endl;
    fprintf(stderr, "Usage: %s <output:4debug> <vmc_index> <_exchange_idx> [ip:port]\n", argv[0]);
    exit(0);
  }

  // default
  int idx = 0,
      idx_exch = 5000,
      cnt_remote = 9,
      cnt_exchange = 7;
  if (argc>=3) idx = atoi(argv[2]);
  if (argc>=4) idx_exch = atoi(argv[3]);
  if (idx_exch<5000) idx_exch+=5000;

  const char* ip = "127.0.0.1";
  unsigned short port = 9191;
# define LEN_BUF 128
  char szBuf[LEN_BUF] = {0};
  if (argc==5) {
    size_t len = strlen(argv[4]); 
    if (len>=LEN_BUF) {
      cerr<<"Invalid ip address: "<<argv[4]<<endl;
      exit(-1);
    }
    strncpy(szBuf, argv[4], LEN_BUF);
    char* p = szBuf;
    for (; p[0] && p[0]!=':'; p++) {}
    if (!p[0]) {
      cerr<<"Invalid ip address: "<<argv[4]<<endl;
      exit(-1);
    }
    *p++='\0';
    ip = szBuf;
    port = atoi(p);
  }
  cout<<"ip: "<<ip<<"; port:"<<port<<"."<<endl;

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

  // send
  int sockfd;
  socklen_t len = sizeof(struct sockaddr_in);

  struct sockaddr_in serveraddr;
  memset(&serveraddr, 0, len);

  serveraddr.sin_family = AF_INET;
  serveraddr.sin_port = htons(port);
  serveraddr.sin_addr.s_addr = inet_addr(ip);

  sockfd = socket(AF_INET, SOCK_DGRAM, 0);
  if (sockfd < 0)
  {
    cerr << "fail to create socker" << endl;
    delete [] buf;
    exit(-1);
  }
  ssize_t ret = sendto(sockfd, buf, sz_out, 0, (struct sockaddr*)&serveraddr, len);
  if (ret == -1) {
    cerr << "sendto failed, ret code: "<<ret<<endl;
    close(sockfd);
    delete [] buf;
    exit(-1);
  }
  cout << "send:"<<ret<<endl;
  close(sockfd);

  delete [] buf;
}
