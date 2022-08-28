package handler

import (
	"net"
	"star_atlas_server/model"

	"github.com/golang/glog"
)

const (
	CChanLen  = 1000
	CDataSize = 10240
)

var limitChan = make(chan string, CChanLen)
var doneChan = make(chan bool, CChanLen)
var vmcDataChan = make(chan *model.VMCData, CDataSize)

func UdpDataRev(port int) {
	glog.Infof("start listening on port:%d\n", port)
	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: port,
	})

	defer func() {
		conn.Close()
		if err := recover(); err != nil {
			glog.Errorf("go UdpDataRev error: %v", err)
		}
	}()

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
	glog.Infof("udp channel len:%d\n", len(limitChan))
	limitChan <- str
}

func UpdateVMC() {
	defer func() {
		if err := recover(); err != nil {
			glog.Errorf("go UpdateVMC error: %v", err)
		}
	}()
	for {
		vmcData, ok := <-vmcDataChan
		if !ok {
			glog.Errorf("received vmcData failed")
			continue
		}
		topoTable := &model.TopoTable{}
		err := topoTable.CreateOp(vmcData)
		if err != nil {
			glog.Errorf("failed create topotable into db, error: %s\n", err.Error())
			continue
		}
		err = vmcData.CreateData()
		if err != nil {
			glog.Errorf("failed create vmcdata into db, error: %s\n", err.Error())
			continue
		}
	}
}

func ParseData() {

	defer func() {
		if err := recover(); err != nil {
			glog.Errorf("go parse error: %v", err)
			go ParseData()
		}
	}()

	for {
		data, ok := <-limitChan
		if !ok {
			glog.Errorf("recv err\n")
			continue
		}
		ld := len(data)
		glog.Infof("recv data len: %v\n", ld)
		if ld < 4 {
			glog.Warningf("received data length %d err\n", ld)
			continue
		}
		raw := []byte(data)
		p := raw[3]
		if p == 0xaa {
			go func(){
				controller, err := model.NewVMCController(raw)
				if err != nil {
					glog.Errorf("failed create vmc controller, error: %s\n", err.Error())
					return
				}
				err = controller.SaveSelf()
				if err != nil {
					glog.Errorf("failed save controller data, error: %s\n", err.Error())
				}
				err = controller.FindAndSetFailureEntity()
				if err != nil {
					glog.Errorf("failed FindAndSetFailureEntity, error: %s\n", err.Error())
				}
			}()
			
			continue
		}
		vmcData, err := model.NewVMCData(data)
		if err != nil {
			glog.Errorf("recv err: %v", err)
			continue
		}
		vmcDataChan <- vmcData
		<-doneChan
	}
}
