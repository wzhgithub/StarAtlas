package handler

import (
	"github.com/golang/glog"
	"net"
	"star_atlas_server/db"
	"star_atlas_server/model"
)

const (
	CChanLen  = 1000
	CDataSize = 10240
)

var limitChan = make(chan string, CChanLen)
var doneChan = make(chan bool, CChanLen)

func UdpDataRev(port int) {
	glog.Info("start listening on port:", port)
	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: port,
	})

	defer conn.Close()
	if err != nil {
		glog.Fatalf("read from connect failed, err:%s\n", err.Error())
	}

	for {
		doneChan <- true
		udpProcess(conn)
	}
}

func udpProcess(conn *net.UDPConn) {
	// 最大读取数据大小
	data := make([]byte, CDataSize)
	n, address, err := conn.ReadFromUDP(data)
	if err != nil {
		glog.Errorf("failed read udp msg, error:%s\n", err.Error())
	}
	glog.Infof("received adddress:%s\n", *address)
	str := string(data[:n])
	glog.Infof("received adddressfrom client data:%s\n", str)
	limitChan <- str
}

func ParseData() {
	for {
		data, ok := <-limitChan
		if !ok {
			glog.Errorf("recv err\n")
			continue
		}
		_, _ = model.NewVMCData(data)
		db.Test()
		<-doneChan
	}
}
