package handler

import (
	"net"
	"star_atlas_server/db"
	"star_atlas_server/model"

	"github.com/golang/glog"
)

const (
	CChanLen  = 1000
	CDataSize = 10240
)

var limitChan = make(chan string, CChanLen)
var doneChan = make(chan bool, CChanLen)
var vmcDataChan = make(chan *model.VMCData, CChanLen)

func UdpDataRev(port int) {
	glog.Infof("start listening on port:%d\n", port)
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
		vmcdata, _ := model.NewVMCData(data) /*  */
		vmcDataChan <- vmcdata
		//todo

		<-doneChan
	}
}

func CreateDBData() {
	for {
		vmcdata, ok := <-vmcDataChan
		if !ok {
			glog.Errorf("get vmcdata err\n")
			continue
		}
		db.CreateVMCData(vmcdata)
		// TODO: wirte all of data into db
	}

}
