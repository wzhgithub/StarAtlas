package model

import "encoding/binary"

type DeviceData struct {
	Name string // 10bytes
	ID   uint8
	Type uint8 // ARM:0,RISC_V:1,SPARC:2,PPC:3,MIPS:4
	// fpga don`t has below fields
	Num                 uint8  //
	IntComputingPower   uint16 // gpu this feild is nil
	FloatComputingPower uint16
	TotalMemory         uint16
	MemoryUsage         uint8
	Usage               uint8
}

type Task struct {
	Name        uint16
	TaskType    uint8
	TaskStatus  uint8
	ExecuteTime uint8
}

type VMCData struct {
	frameHeader    uint8
	length         uint16
	protoType      uint8
	VMCName        string // 10bytes
	VMCID          uint8
	CPUNumber      uint8
	DSPNumber      uint8
	GPUNumber      uint8
	FPAGNumber     uint8
	SwitchID       uint8
	TotalMemory    uint16
	TotalDisk      uint16
	MemoryUsage    uint8
	TotalCPUUsage  uint8
	TotalDSPUsage  uint8
	TotalGPUUsage  uint8
	TotalDiskUsage uint8
	// cpu
	CPUSet []*DeviceData // 21bytes
	// dsp
	DSPSet []*DeviceData // 21bytes
	// gpu
	GPUSet []*DeviceData // 19bytes
	// fpga
	FPGASet      []*DeviceData // 12bytes
	APPNum       uint
	APPName      string // 10bytes
	TaskNum      uint   // max 6
	TaskPeriod   uint16 // 250ms
	TaskDispatch uint16 // 100ms
	TaskSet      []*Task
}

func parseCPUDevice(bytes []byte) []*DeviceData {
	l := len(bytes)
	s := binary.BigEndian.Uint16(bytes[0:2])
	e := binary.BigEndian.Uint16(bytes[l-2 : l])
	if (l-4)%21 == 0 && s == 0xeba0 && e == 0xebaa {
		arr := make([]*DeviceData, (l-4)/21)
		for i := 2; i < l-1; i = i + 21 {

			DeviceData := &DeviceData{
				Name:                string(bytes[i : i+10]),
				ID:                  bytes[i+10],
				Type:                bytes[i+11],
				Num:                 bytes[i+12],
				IntComputingPower:   binary.BigEndian.Uint16(bytes[i+13 : i+15]),
				FloatComputingPower: binary.BigEndian.Uint16(bytes[i+15 : i+17]),
				TotalMemory:         binary.BigEndian.Uint16(bytes[i+17 : i+19]),
				MemoryUsage:         bytes[i+19],
				Usage:               bytes[i+20],
			}
			arr = append(arr, DeviceData)
		}
		return arr
	}
	return nil
}

func parseGPUDevice(bytes []byte) []*DeviceData {

	return nil
}

func parseFPGADevice(bytes []byte) []*DeviceData {

	return nil
}

func parseDSPDevice(bytes []byte) []*DeviceData {

	return nil
}

// todo
func parseTask(bytes []byte) (*Task, error) {
	return nil, nil
}

// todo
func parse(bytes []byte) (*VMCData, error) {
	v := &VMCData{
		frameHeader:    bytes[0],
		length:         binary.BigEndian.Uint16(bytes[1:3]),
		protoType:      bytes[3],
		VMCName:        string(bytes[4:14]),
		VMCID:          bytes[15],
		CPUNumber:      bytes[16],
		DSPNumber:      bytes[17],
		GPUNumber:      bytes[18],
		FPAGNumber:     bytes[19],
		SwitchID:       bytes[20],
		TotalMemory:    binary.BigEndian.Uint16(bytes[21:23]),
		TotalDisk:      binary.BigEndian.Uint16(bytes[23:25]),
		MemoryUsage:    bytes[25],
		TotalCPUUsage:  0,
		TotalDSPUsage:  0,
		TotalGPUUsage:  0,
		TotalDiskUsage: 0,
		CPUSet:         parseCPUDevice(bytes[26 : 26+1+21*bytes[16]+2+1]),
		DSPSet:         []*DeviceData{},
		GPUSet:         []*DeviceData{},
		FPGASet:        []*DeviceData{},
		APPNum:         0,
		APPName:        "",
		TaskNum:        0,
		TaskPeriod:     0,
		TaskDispatch:   0,
		TaskSet:        []*Task{},
	}
	return v, nil
}

// read bytes from udp
func NewVMCData(str string) (*VMCData, error) {
	return parse([]byte(str))
}
