#include "telemsg.h"
#include <sys/socket.h>
#include <sys/types.h>
#include <unistd.h>
#include <arpa/inet.h>
#include <assert.h>
#include <memory>
#include <regex>
#include <string>
#include <vector>

#include "curl_impl.h"
#include "rapidjson/istreamwrapper.h"
#include "rapidjson/schema.h"
#include "rapidjson/stringbuffer.h"

using namespace std;

typedef enum {
  eTransVMC = 0,
  eTransDev,
  eTransHolder
} ETransType;

int main(int argc, char* argv[]) {
  if (argc==1) {
    fprintf(stderr, "Usage: %s <url> <type:[0|1]> <pack>\n", argv[0]);
    exit(0);
  }

  const char* ip = "127.0.0.1";
  unsigned short port = 9191;
  const char* _url = "http://www.baidu.com/";
  int _dtype = 0;
  const char* _data = "0|100|0|200|a|1|1";
  if (argc>1) { _url = argv[1]; }
  if (argc>2) { _dtype = atoi(argv[2]); }
  if (argc>3) { _data = argv[3]; }

  std::shared_ptr<CurlPSImpl> client(new CurlPSImpl);
  client->Initialize(_url);
  std::string res;
  client->Execute(res);
  std::cout<<res<<std::endl;
  // replace ip & port
  rapidjson::Document _doc;
  _doc.Parse(res.c_str(), res.length());
  if (_doc.HasParseError() ||
    _doc.HasMember("123")) {
    //exit(-1);
  }
  
  // 1. send to central control
  uint8_t _from = 0, _from_dev = 0, _to = 0, _to_dev = 0, _type = 0, _idx = 0;
  ControlMessage msg;
  cout<<"send control message to central control"<<endl;

  string _sdata = string(_data);
  string _del = string("|");
  vector<string> _items;
  try {
    regex re{_del};
    _items = vector<string>{
      sregex_token_iterator(_sdata.begin(), _sdata.end(), re, -1),
      sregex_token_iterator()
    };      
  }
  catch(const std::exception& e) {
    cout<<"error:"<<e.what()<<std::endl;
    exit(-1);
  }

  if (_dtype == eTransVMC) {
    // "from|to|idx"
    assert (_items.size()==3);
    _from = atoi(_items[0].c_str());
    _to = atoi(_items[1].c_str());
    _idx  = atoi(_items[2].c_str());
    msg.setMessage(_from, _to, _idx);
  } else if (_dtype == eTransDev) {
    // "0|100|0|200|a|1|1";
    assert (_items.size()==7);
    _from = atoi(_items[0].c_str());
    _from_dev = atoi(_items[1].c_str());
    _to = atoi(_items[2].c_str());
    _to_dev = atoi(_items[3].c_str());
    const char* _name = _items[4].c_str();
    _type = atoi(_items[5].c_str());
    _idx = atoi(_items[6].c_str());
    msg.setMessage(_from, _from_dev, _to, _to_dev, _name, _type, _idx);
  } else {
    cout<<"Invalide data"<<endl;
    exit(-1);
  }

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

  size_t _sz = msg.getSize()+4;
  char* buf = new char[_sz];
  memset(buf, 0, _sz);
  int sz_out = msg.pack(buf);
  ssize_t ret = sendto(sockfd, buf, sz_out, 0, (struct sockaddr*)&serveraddr, len);
  if (ret == -1) {
    cerr << "sendto failed, ret code: "<<ret<<endl;
    close(sockfd);
    exit(-1);
  }
    
  close(sockfd);
  cout<<"send message to"<<ip<<":"<<port<<"successed."<<endl;
  exit(0);
}
