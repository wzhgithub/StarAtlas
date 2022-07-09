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
	glog.Infof("received address:%+v\n", address)
	str := string(data[:n])
	limitChan <- str
}

func ParseData() {
	for {
		data, ok := <-limitChan
		if !ok {
			glog.Errorf("recv err\n")
			continue
		}
		vmcData, _ := model.NewVMCData(data)
		var topoTable = model.TopoTable{}
		topoTable.CreateOp(vmcData)
		err := vmcData.CreateData()
		if err != nil {
			glog.Error("failed create vmcdata into db, error: %s\n", err.Error())
		}

		vmcdata_read := &model.VMCData{}
		err = vmcdata_read.CollectVMCData()
		if err != nil {
			glog.Error("failed read vmcdata from db, error: %s\n", err.Error())
		}
		db.Test()
		<-doneChan
	}
}
