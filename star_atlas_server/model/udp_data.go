package model

import (
	"encoding/binary"
	"fmt"
	"star_atlas_server/config"

	"github.com/golang/glog"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	cCPU_SIZE  = 21
	cDSP_SIZE  = 21
	cFPGA_SIZE = 12
	cGPU_SIZE  = 19

	cCPU_START  = 0xeba0
	cCPU_END    = 0xebaa
	cDSP_START  = 0xebb0
	cDSP_END    = 0xebbb
	cFPGA_START = 0xebd0
	cFPGA_END   = 0xebdd
	cGPU_START  = 0xebc0
	cGPU_END    = 0xebcc

	cTAST_SIZE = 13
	cAPP_SIZE  = 94
)

type DeviceData struct {
	Name string `json:"name" bson:"name"` // 10bytes
	ID   uint8  `json:"id" bson:"id"`
	Type uint8  `json:"type" bson:"type"` // ARM:0,RISC_V:1,SPARC:2,PPC:3,MIPS:4
	// fpga don`t has below fields
	Num                 uint8  `json:"num" bson:"num"`
	IntComputingPower   uint16 `json:"int_computing_power" bson:"int_computing_power"` // gpu this feild is nil
	FloatComputingPower uint16 `json:"float_computing_power" bson:"float_computing_power"`
	TotalMemory         uint16 `json:"total_memory" bson:"total_memory"`
	MemoryUsage         uint8  `json:"memory_usage" bson:"memory_usage"`
	Usage               uint8  `json:"usage" bson:"usage"`
}

type Task struct {
	Name        string `json:"name" bson:"name"`
	TaskType    uint8  `json:"task_type" bson:"task_type"`
	TaskStatus  uint8  `json:"task_status" bson:"task_status"`
	ExecuteTime uint8  `json:"execute_time" bson:"execute_time"`
}

type App struct {
	APPName      string  `json:"app_name" bson:"app_name"`           // 10bytes
	TaskNum      uint8   `json:"task_num" bson:"task_num"`           // max 6
	TaskPeriod   uint16  `json:"task_period" bson:"task_period"`     // 250ms
	TaskDispatch uint16  `json:"task_dispatch" bson:"task_dispatch"` // 100ms
	TaskSet      []*Task `json:"task_set" bson:"task_set"`
	CurrentTask  uint8   `json:"current_task" bson:"current_task"`
}

type VMCData struct {
	mgm.DefaultModel `bson:",inline"`
	frameHeader      uint8
	length           uint16
	protoType        uint8
	VMCName          string `json:"vmc_name" bson:"vmc_name"` // 10bytes
	VMCID            uint8  `json:"vmc_id" bson:"vmc_id"`
	CPUNumber        uint8  `json:"cpu_number" bson:"cpu_number"`
	DSPNumber        uint8  `json:"dsp_number" bson:"dsp_number"`
	GPUNumber        uint8  `json:"gpu_number" bson:"gpu_number"`
	FPGANumber       uint8  `json:"fpga_number" bson:"fpga_number"`
	SwitchID         uint8  `json:"switch_id" bson:"switch_id"`
	TotalMemory      uint16 `json:"total_memory" bson:"total_memory"`
	TotalDisk        uint16 `json:"total_disk" bson:"total_disk"`
	MemoryUsage      uint8  `json:"memory_usage" bson:"memory_usage"`
	TotalCPUUsage    uint8  `json:"total_cpu_usage" bson:"total_cpu_usage"`
	TotalDSPUsage    uint8  `json:"total_dsp_usage" bson:"total_dsp_usage"`
	TotalGPUUsage    uint8  `json:"total_gpu_usage" bson:"total_gpu_usage"`
	TotalDiskUsage   uint8  `json:"total_disk_usage" bson:"total_disk_usage"`
	// cpu
	CPUSet []*DeviceData `json:"cpu_set" bson:"cpu_set"` // 21bytes
	// dsp
	DSPSet []*DeviceData `json:"dsp_set" bson:"dsp_set"` // 21bytes
	// gpu
	GPUSet []*DeviceData `json:"gpu_set" bson:"gpu_set"` // 19bytes
	// fpga
	FPGASet []*DeviceData `json:"fpga_set" bson:"fpga_set"` // 21bytes
	APPNum  uint8         `json:"app_num" bson:"app_num"`
	APPInfo []*App        `json:"app_info" bson:"app_info"`
	Sum     uint8         `json:"sum" bson:"sum"`
}

func parseCPUDevice(bytes []byte, start, end int) []*DeviceData {
	if end <= start {
		return nil
	}
	bytes = bytes[start:end]
	l := len(bytes)
	s := binary.BigEndian.Uint16(bytes[0:2])
	e := binary.BigEndian.Uint16(bytes[l-2 : l])
	glog.Infof("l:%d s:%d e:%d\n", l, s, e)
	if (l-4)%cCPU_SIZE == 0 && s == cCPU_START && e == cCPU_END {
		num := (l - 4) / cCPU_SIZE
		arr := make([]*DeviceData, num)
		ss := 2
		for i := 0; i < num; i++ {
			si := i*cCPU_SIZE + ss
			glog.Infof("cpu index %d num:%d\n", si, num)
			DeviceData := &DeviceData{
				Name:                string(bytes[si : si+10]),
				ID:                  bytes[si+10],
				Type:                bytes[si+11],
				Num:                 bytes[si+12],
				IntComputingPower:   binary.BigEndian.Uint16(bytes[si+13 : si+15]),
				FloatComputingPower: binary.BigEndian.Uint16(bytes[si+15 : si+17]),
				TotalMemory:         binary.BigEndian.Uint16(bytes[si+17 : si+19]),
				MemoryUsage:         bytes[si+19],
				Usage:               bytes[si+20],
			}
			glog.Infof("cpu device %+v\n", DeviceData)
			arr[i] = DeviceData
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
	if (l-4)%cGPU_SIZE == 0 && s == cGPU_START && e == cGPU_END {
		num := (l - 4) / cGPU_SIZE
		arr := make([]*DeviceData, num)
		ss := 2
		for j := 0; j < num; j++ {
			i := j*cGPU_SIZE + ss
			glog.Infof("gpu i:%d num:%d\n", i, num)
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
			glog.Infof("gpu: %+v\n", DeviceData)
			arr[j] = DeviceData
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
	ss := 2
	if (l-4)%cFPGA_SIZE == 0 && s == cFPGA_START && e == cFPGA_END {
		num := (l - 4) / cFPGA_SIZE
		arr := make([]*DeviceData, num)
		for j := 0; j < num; j++ {
			i := j*cFPGA_SIZE + ss
			glog.Infof("fpga i:%d num:%d\n", i, num)
			DeviceData := &DeviceData{
				Name: string(bytes[i : i+10]),
				ID:   bytes[i+10],
				Type: bytes[i+11],
			}
			glog.Infof("fpga:%+v", DeviceData)
			arr[j] = DeviceData
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
	glog.Infof("l:%d s:%d e:%d\n", l, s, e)
	if (l-4)%cDSP_SIZE == 0 && s == cDSP_START && e == cDSP_END {
		num := (l - 4) / cDSP_SIZE
		arr := make([]*DeviceData, num)
		ss := 2
		for i := 0; i < num; i++ {
			si := i*cDSP_SIZE + ss
			glog.Infof("dsp index %d num:%d\n", si, num)
			DeviceData := &DeviceData{
				Name:                string(bytes[si : si+10]),
				ID:                  bytes[si+10],
				Type:                bytes[si+11],
				Num:                 bytes[si+12],
				IntComputingPower:   binary.BigEndian.Uint16(bytes[si+13 : si+15]),
				FloatComputingPower: binary.BigEndian.Uint16(bytes[si+15 : si+17]),
				TotalMemory:         binary.BigEndian.Uint16(bytes[si+17 : si+19]),
				MemoryUsage:         bytes[si+19],
				Usage:               bytes[si+20],
			}
			glog.Infof("dsp device %+v\n", DeviceData)
			arr[i] = DeviceData
		}
		return arr
	}
	return nil
}

func parseTask(bytes []byte, start, end int) []*Task {
	bytes = bytes[start:end]
	length := len(bytes)
	glog.Infof("task length: %d", length)
	arr := make([]*Task, 6)
	for j := 0; j < 6; j++ {
		i := j * cTAST_SIZE
		glog.Infof("task start index: %d", i)
		t := &Task{
			Name:        string(bytes[i : 10+i]),
			TaskType:    bytes[10+i],
			TaskStatus:  bytes[11+i],
			ExecuteTime: bytes[12+i],
		}
		arr[j] = t
	}

	return arr
}

func parseApp(bytes []byte, start, end int) []*App {
	if end <= start {
		return nil
	}
	bytes = bytes[start:end]
	length := len(bytes)
	if length%cAPP_SIZE != 0 {
		glog.Warningf("app len is illegal length %d", length)
		return nil
	}
	cnt := length / cAPP_SIZE
	glog.Infof("app length %d, start %d, end %d cnt %d\n", length, start, end, cnt)
	arr := make([]*App, cnt)
	for j := 0; j < cnt; j++ {
		i := j * cAPP_SIZE
		a := &App{
			APPName:      string(bytes[i : i+10]),
			TaskNum:      bytes[i+10],
			TaskPeriod:   binary.BigEndian.Uint16(bytes[i+11 : i+13]),
			TaskDispatch: binary.BigEndian.Uint16(bytes[i+13 : i+15]),
			TaskSet:      parseTask(bytes, i+15, i+93),
			CurrentTask:  bytes[i+93],
		}
		arr[j] = a
	}

	return arr
}

func calcStartEnd(start int, num uint8, l int) (int, int) {
	if num == 0 {
		glog.Infof("num is zero return same as start\n")
		return start, start
	}

	glog.Infoln(fmt.Sprintf("s:%d, n:%d, l:%d", start, num, l))
	return start, start + 1 + l*int(num) + 2 + 1
}

// todo
func parse(bytes []byte) (*VMCData, error) {

	l := len(bytes)
	deviceIdx := 29
	cpuStart, cpuEnd := calcStartEnd(deviceIdx, uint8(bytes[15]), 21)
	glog.Infof("cpu start:%d cpu end:%d\n", cpuStart, cpuEnd)
	dspStart, dspEnd := calcStartEnd(cpuEnd, uint8(bytes[16]), 21)
	glog.Infof("dsp start:%d dsp end:%d\n", dspStart, dspEnd)
	gpusStart, gpusEnd := calcStartEnd(dspEnd, uint8(bytes[17]), 19)
	glog.Infof("gpu start:%d gpus end:%d\n", gpusStart, gpusEnd)
	fpgaStart, fpgaEnd := calcStartEnd(gpusEnd, uint8(bytes[18]), 12)
	glog.Infof("fpga start:%d fpga end:%d\n", fpgaStart, fpgaEnd)
	appIdx := cpuStart
	if cpuStart < fpgaEnd {
		appIdx = fpgaEnd
	}
	glog.Infof("app idx: %d\n", appIdx)

	v := &VMCData{
		frameHeader:    bytes[0],
		length:         binary.BigEndian.Uint16(bytes[1:3]),
		protoType:      bytes[3],
		VMCName:        string(bytes[4:14]),
		VMCID:          bytes[14],
		CPUNumber:      bytes[15],
		DSPNumber:      bytes[16],
		GPUNumber:      bytes[17],
		FPGANumber:     bytes[18],
		SwitchID:       bytes[19],
		TotalMemory:    binary.BigEndian.Uint16(bytes[20:22]),
		TotalDisk:      binary.BigEndian.Uint16(bytes[22:24]),
		MemoryUsage:    bytes[24],
		TotalCPUUsage:  bytes[25],
		TotalDSPUsage:  bytes[26],
		TotalGPUUsage:  bytes[27],
		TotalDiskUsage: bytes[28],
		CPUSet:         parseCPUDevice(bytes, cpuStart, cpuEnd),
		DSPSet:         parseDSPDevice(bytes, dspStart, dspEnd),
		GPUSet:         parseGPUDevice(bytes, gpusStart, gpusEnd),
		FPGASet:        parseFPGADevice(bytes, fpgaStart, fpgaEnd),
		APPNum:         bytes[appIdx],
		APPInfo:        parseApp(bytes, appIdx+1, l-1),
		Sum:            bytes[l-1],
	}
	glog.Infof("debug vmc:%+v\n", v)
	return v, nil
}

// read bytes from udp
func NewVMCData(str string) (*VMCData, error) {
	return parse([]byte(str))
}

type DBVMCDataInterface interface {
	CreateData(vmc_data *VMCData) error
	CollectData(vmc_data *VMCData) error
}

func (vmc_data *VMCData) CreateData() error {
	return mgm.CollectionByName(config.CommonConfig.DBVMCDataTableName).Create(vmc_data)
}

func (vmc_data *VMCData) CollectVMCData() error {
	return mgm.CollectionByName(config.CommonConfig.DBVMCDataTableName).First(bson.M{}, vmc_data, &options.FindOneOptions{Sort: bson.M{"_id": -1}})
}
