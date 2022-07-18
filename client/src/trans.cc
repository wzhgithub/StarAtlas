#include "telemsg.h"
#include <sys/socket.h>
#include <sys/types.h>
#include <unistd.h>
#include <arpa/inet.h>

int main(int argc, char* argv[]) {
  if (argc==1) {
    fprintf(stderr, "Usage: %s <source> <dest> [ip:port]\n", argv[0]);
    exit(0);
  }

  uint8_t idx_source = 0, idx_dest = 1;
  if (argc>1) { idx_source = atoi(argv[1]); }
  if (argc>2) { idx_dest = atoi(argv[2]); }

  const char* ip = "127.0.0.1";
  unsigned short port = 9191;
# define LEN_BUF 128
  char szBuf[LEN_BUF] = {0};
  if (argc==4) {
    size_t len = strlen(argv[3]); 
    if (len>=LEN_BUF) {
      cerr<<"Invalid ip address: "<<argv[3]<<endl;
      exit(-1);
    }
    strncpy(szBuf, argv[3], LEN_BUF);
    char* p = szBuf;
    for (; p[0] && p[0]!=':'; p++) {}
    if (!p[0]) {
      cerr<<"Invalid ip address: "<<argv[3]<<endl;
      exit(-1);
    }
    *p++='\0';
    ip = szBuf;
    port = atoi(p);
  }

  // 1. send to central control
  ControlMessage msg;
  cout<<"send control message to central control"<<endl;

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
    exit(-1);
  }

  uint8_t anVmc[] = {
    idx_source,
    idx_dest,
  };
  for (int h=0; h<sizeof(anVmc); h++) {
    TeleMessage msg2(
      anVmc[h],
      0, // idx_exch,
      0, // cnt_cpu,
      0, // cnt_dsp,
      0, // cnt_gpu,
      0, // cnt_fpga,
      0, // cnt_exchange,
      0, // cnt_remote,
      0, // cnt_block,
      0 //cnt_max_task
    );
    
    char* buf = new char[msg2.m_size+16];
    memset(buf, 0, msg2.m_size+16);
    int sz_out = msg2.pack(buf);
    ssize_t ret = sendto(sockfd, buf, sz_out, 0, (struct sockaddr*)&serveraddr, len);
    if (ret == -1) {
      cerr << "sendto failed, ret code: "<<ret<<endl;
      close(sockfd);
      exit(-1);
    }
    
    sleep(1);
  }
  close(sockfd);
  cout<<"send message to"<<ip<<":"<<port<<"successed."<<endl;
  exit(0);
}
