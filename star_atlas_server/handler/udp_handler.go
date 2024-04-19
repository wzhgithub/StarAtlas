package handler

import (
	"encoding/binary"
	"fmt"
	"math"
	"net"
	"star_atlas_server/model"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

const (
	CChanLen  = 1000
	CDataSize = 10240
)

type ComputerPowerInfo struct {
	IntComputerPower   float32 `json:"intComputerPower"`
	FloatComputerPower float32 `json:"floatComputerPower"`
	FaultTolerance     float32 `json:"faultTolerance"`
}

var limitChan = make(chan string, CChanLen)
var doneChan = make(chan bool, CChanLen)
var vmcDataChan = make(chan *model.VMCData, CDataSize)
var srcAddress *net.UDPAddr = nil
var connect *net.UDPConn = nil

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
	connect = conn
	for {
		doneChan <- true
		udpProcess(conn)
	}
}

func float32ToBytes(float float32) []byte {
	bits := math.Float32bits(float)
	bytes := make([]byte, 4)
	binary.BigEndian.PutUint32(bytes, bits)
	return bytes
}

func (cp *ComputerPowerInfo) toBytes() []byte {
	head := []byte{0xeb}
	b1 := float32ToBytes(cp.IntComputerPower)
	b2 := float32ToBytes(cp.FloatComputerPower)
	b3 := float32ToBytes(cp.FaultTolerance)
	data := append(append(b1, b2...), b3...)
	data = append(head, data...)

	var sum uint8 = 0
	for _, v := range data {
		sum += uint8(v)
	}
	data = append(data, byte(sum))
	return data
}

func SendComputerPower(c *gin.Context) {
	cp := &ComputerPowerInfo{}
	if err := c.ShouldBindJSON(cp); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	glog.Infof("recv computer power data: %+v\n", cp)
	data := cp.toBytes()
	err := sendUDPData(data)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "success", "code": 0})
}

func sendUDPData(data []byte) error {
	if srcAddress == nil || connect == nil {
		glog.Errorf("send udp data failed, srcAddress or connect is nil\n")
		return fmt.Errorf("send udp data failed, srcAddress or connect is nil")
	}
	if len(data) > CDataSize {
		glog.Errorf("send udp data failed, data len is too long\n")
		return fmt.Errorf("send udp data failed, data len is too long")
	}
	glog.Infof("send byte data:%+v\n", data)
	_, err := connect.WriteToUDP(data, srcAddress)
	return err
}

func udpProcess(conn *net.UDPConn) {
	// 最大读取数据大小
	data := make([]byte, CDataSize)
	n, address, err := conn.ReadFromUDP(data)
	if err != nil {
		glog.Errorf("failed read udp msg, error:%s\n", err.Error())
	}
	glog.Infof("received address:%+v\n", address)
	if address != nil {
		srcAddress = address
	}
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
			go func() {
				defer func() {
					if err := recover(); err != nil {
						glog.Errorf("go controller handle error: %v\n", err)
					}
				}()
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
			glog.Infof("controller async handle by p:%d\n", p)
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
