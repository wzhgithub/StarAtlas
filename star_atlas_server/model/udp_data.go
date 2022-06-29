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
	Name        string
	TaskType    uint8
	TaskStatus  uint8
	ExecuteTime uint8
}

type App struct {
	APPName      string // 10bytes
	TaskNum      uint8  // max 6
	TaskPeriod   uint16 // 250ms
	TaskDispatch uint16 // 100ms
	TaskSet      []*Task
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
	FPGASet []*DeviceData // 12bytes
	APPNum  uint8
	APPInfo []*App
}

func parseCPUDevice(bytes []byte, start, end int) []*DeviceData {
	if end <= start {
		return nil
	}
	bytes = bytes[start:end]
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

func parseGPUDevice(bytes []byte, start, end int) []*DeviceData {
	if end <= start {
		return nil
	}
	bytes = bytes[start:end]
	l := len(bytes)
	s := binary.BigEndian.Uint16(bytes[0:2])
	e := binary.BigEndian.Uint16(bytes[l-2 : l])
	if (l-4)%19 == 0 && s == 0xebc0 && e == 0xebcc {
		arr := make([]*DeviceData, (l-4)/19)
		for i := 2; i < l-1; i = i + 19 {

			DeviceData := &DeviceData{
				Name:                string(bytes[i : i+10]),
				ID:                  bytes[i+10],
				Type:                bytes[i+11],
				Num:                 bytes[i+12],
				IntComputingPower:   0,
				FloatComputingPower: binary.BigEndian.Uint16(bytes[i+13 : i+15]),
				TotalMemory:         binary.BigEndian.Uint16(bytes[i+15 : i+17]),
				MemoryUsage:         bytes[i+17],
				Usage:               bytes[i+18],
			}
			arr = append(arr, DeviceData)
		}
		return arr
	}
	return nil
}

func parseFPGADevice(bytes []byte, start, end int) []*DeviceData {
	if end <= start {
		return nil
	}
	bytes = bytes[start:end]
	l := len(bytes)
	s := binary.BigEndian.Uint16(bytes[0:2])
	e := binary.BigEndian.Uint16(bytes[l-2 : l])
	if (l-4)%12 == 0 && s == 0xebd0 && e == 0xebdd {
		arr := make([]*DeviceData, (l-4)/12)
		for i := 2; i < l-1; i = i + 12 {
			DeviceData := &DeviceData{
				Name: string(bytes[i : i+10]),
				ID:   bytes[i+10],
				Type: bytes[i+11],
			}
			arr = append(arr, DeviceData)
		}
		return arr
	}

	return nil
}

func parseDSPDevice(bytes []byte, start, end int) []*DeviceData {
	if end <= start {
		return nil
	}
	bytes = bytes[start:end]
	l := len(bytes)
	s := binary.BigEndian.Uint16(bytes[0:2])
	e := binary.BigEndian.Uint16(bytes[l-2 : l])
	if (l-4)%21 == 0 && s == 0xebb0 && e == 0xebbb {
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

func parseTask(bytes []byte, start, end int) []*Task {
	l := len(bytes)
	arr := make([]*Task, 6)
	t := &Task{
		Name:        string(bytes[0:2]),
		TaskType:    bytes[2],
		TaskStatus:  bytes[3],
		ExecuteTime: bytes[4],
	}
	arr = append(arr, t)

	for i := 5; i < l; i = i + 13 {
		t := &Task{
			Name:        string(bytes[i : 10+i]),
			TaskType:    bytes[10+i],
			TaskStatus:  bytes[11+i],
			ExecuteTime: bytes[12+i],
		}
		arr = append(arr, t)
	}

	return arr
}

func parseApp(bytes []byte, start, end int) []*App {
	if end <= start {
		return nil
	}
	bytes = bytes[start:end]
	length := len(bytes)
	if length%86 == 0 {
		arr := make([]*App, length/86)
		for i := 0; i < length; i = i + 86 {
			a := &App{
				APPName:      string(bytes[i : i+10]),
				TaskNum:      bytes[i+10],
				TaskPeriod:   binary.BigEndian.Uint16(bytes[i+11 : i+13]),
				TaskDispatch: binary.BigEndian.Uint16(bytes[i+13 : i+15]),
				TaskSet:      parseTask(bytes, i+15, i+86),
			}
			arr = append(arr, a)
		}
		return arr
	}
	return nil
}

func calcStartEnd(start int, num uint8, l int) (int, int) {
	if num == 0 {
		return start, start
	}

	return start, 2 + l*int(num) + 2
}

// todo
func parse(bytes []byte) (*VMCData, error) {

	l := len(bytes)
	cpuStart, cpuEnd := calcStartEnd(26, uint8(bytes[16]), 21)
	dspStart, dspEnd := calcStartEnd(cpuEnd, uint8(bytes[17]), 21)
	gpusStart, gpusEnd := calcStartEnd(dspEnd, uint8(bytes[18]), 19)
	fpgaStart, fpgaEnd := calcStartEnd(gpusEnd, uint8(bytes[19]), 12)
	appIdx := cpuStart
	if cpuStart < fpgaEnd {
		appIdx = fpgaEnd
	}

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
		CPUSet:         parseCPUDevice(bytes, cpuStart, cpuEnd),
		DSPSet:         parseDSPDevice(bytes, dspStart, dspEnd),
		GPUSet:         parseGPUDevice(bytes, gpusStart, gpusEnd),
		FPGASet:        parseFPGADevice(bytes, fpgaStart, fpgaEnd),
		APPNum:         bytes[appIdx],
		APPInfo:        parseApp(bytes, appIdx+1, l),
	}
	return v, nil
}

// read bytes from udp
func NewVMCData(str string) (*VMCData, error) {
	return parse([]byte(str))
}
