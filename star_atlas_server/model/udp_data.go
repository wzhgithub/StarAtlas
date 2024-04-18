package model

import (
	bbytes "bytes"
	"context"
	"encoding/binary"
	"fmt"
	"math"
	"star_atlas_server/config"
	"time"

	"github.com/golang/glog"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	cCPU_SIZE    = 25
	cDSP_SIZE    = 25
	cFPGA_SIZE   = 12
	cGPU_SIZE    = 21
	cREMOTE_SIZE = 13
	cSWITCH_SIZE = 13

	cREMOTE_START = 0xebe0
	cSWITCH_START = 0xebf0
	cCPU_START    = 0xeba0
	// cCPU_END      = 0xebaa
	cDSP_START = 0xebb0
	// cDSP_END      = 0xebbb
	cFPGA_START = 0xebd0
	// cFPGA_END   = 0xebdd
	cGPU_START = 0xebc0
	// cGPU_END    = 0xebcc

	cTAST_SIZE = 12

	cMAX_DOCUMENT_NUM = 200

	cSTR_END = 32
)

type FailureOverInfo struct {
	VMCID    string `json:"vmc_id" bson:"vmc_id"`
	AppID    string `json:"app_id" bson:"app_id"`
	TaskID   string `json:"task_id" bson:"task_id"`
	DeviceId string `json:"device_id" bson:"device_id"`
}

type FailureOverRequest struct {
	mgm.DefaultModel `bson:",inline"`
	From             FailureOverInfo `json:"from" bson:"from"`
	To               FailureOverInfo `json:"to" bson:"to"`
	TransStatus      uint            `json:"trans_status" bson:"trans_status"`
	UniqueKey        string          `json:"unique_key" bson:"unique_key"`
}
type FailureTask struct {
	FailureName uint16 `json:"failure_name" bson:"failure_name"`
	FailureDesc uint8  `json:"failure_type" bson:"failure_type"`
	AppId       uint8  `json:"app_id" bson:"app_id"`
	IsFault     uint8  `json:"is_fault" bson:"is_fault"`
}

type VMCController struct {
	mgm.DefaultModel `bson:",inline"`
	FrameHeader      uint8       `json:"frame_header" bson:"frame_header"`
	Length           uint16      `json:"length" bson:"length"`
	DataType         uint8       `json:"data_type" bson:"data_type"`
	FromVmc          uint8       `json:"from_vmc" bson:"from_vmc"`
	FromDevice       uint8       `json:"from_device" bson:"from_device"`
	ToVmc            uint8       `json:"to_vmc" bson:"to_vmc"`
	ToDevice         uint8       `json:"to_device" bson:"to_device"`
	Task             FailureTask `json:"failure_task" bson:"failure_task"`
	Sum              uint8       `json:"sum" bson:"sum"`
}

type DeviceData struct {
	Name string `json:"name" bson:"name"` // 10bytes
	ID   uint8  `json:"id" bson:"id"`
	Type uint8  `json:"type" bson:"type"` // ARM:0,RISC_V:1,SPARC:2,PPC:3,MIPS:4
	// fpga don`t has below fields
	Num                 uint8   `json:"num" bson:"num"`
	IntComputingPower   float32 `json:"int_computing_power" bson:"int_computing_power"` // gpu this feild is nil
	FloatComputingPower float32 `json:"float_computing_power" bson:"float_computing_power"`
	TotalMemory         uint16  `json:"total_memory" bson:"total_memory"`
	MemoryUsage         uint8   `json:"memory_usage" bson:"memory_usage"`
	Usage               uint8   `json:"usage" bson:"usage"`
}

type Task struct {
	Name        string `json:"name" bson:"name"` // 2bytes
	ID          uint16 `json:"id" bson:"id"`     // 2bytes
	TaskType    uint8  `json:"task_type" bson:"task_type"`
	TaskStatus  uint8  `json:"task_status" bson:"task_status"`
	ExecuteTime uint32 `json:"execute_time" bson:"execute_time"` // 4 bytes
	StatusCode  uint8  `json:"status_code" bson:"status_code"`
	StartTime   uint8  `json:"start_time" bson:"start_time"`
	IsTransfer  bool   `json:"is_transfer" bson:"is_transfer"`
}

type App struct {
	APPName      string  `json:"app_name" bson:"app_name"` // 10bytes
	TaskNum      uint8   `json:"task_num" bson:"task_num"`
	TaskPeriod   uint16  `json:"task_period" bson:"task_period"`     // 250ms
	TaskDispatch uint16  `json:"task_dispatch" bson:"task_dispatch"` // 100ms
	ID           uint8   `json:"id" bson:"id"`                       //
	ResetNumber  uint8   `json:"reset_number" bson:"reset_number"`
	BelongsTo    uint8   `json:"belongs_to" bson:"belongs_to"`
	DeviceId     uint8   `json:"device_id" bson:"device_id"`
	TaskSet      []*Task `json:"task_set" bson:"task_set"`
	AppStatus    uint8   `json:"app_status" bson:"app_status"`
	IsTransfer   bool    `json:"is_transfer" bson:"is_transfer"`
}

type RemoteUnit struct {
	RemoteUnitName  string `json:"remote_unit_name" bson:"remote_unit_name"`
	RemoteUnitOrder uint8  `json:"remote_unit_order" bson:"remote_unit_order"`
	RemoteUnitType  uint8  `json:"remote_unit_type" bson:"remote_unit_type"` // 敏感器：0；执行机构：1；载荷：2
	LinkTo          uint8  `json:"link_to" bson:"link_to"`
}

type SwitchDevice struct {
	SwitchName  string `json:"switch_name" bson:"switch_name"` // 10bytes
	SwitchOrder uint8  `json:"switch_order" bson:"switch_order"`
	SwitchType  uint8  `json:"switch_type" bson:"switch_type"` // 中心交换机：0；接入交换机：1
	LinkTo      uint8  `json:"link_to" bson:"link_to"`
}

type VMCStatus struct {
	UpdatedAt            time.Time `json:"time"`                 // 时间
	CPUComputingPower    float32   `json:"cpuComputingPower"`    // cpu算力
	GPUComputingPower    float32   `json:"gpuComputingPower"`    // gpu算力
	DSPIntComputingPower float32   `json:"dspIntComputingPower"` // dsp算力
	MomoryUsage          uint8     `json:"memoryUsage"`          // 内存利用率
	DiskUsage            uint8     `json:"diskUsage"`            // 外存利用率
	TotalUsage           uint8     `json:"totalUsage"`           // 总利用率
	GpuUsage             uint8     `json:"gpuUsage"`             // gpu利用率
	DspUsage             uint8     `json:"dspUsage"`             // dsp利用率
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
	// todo uint8 to float32 uint32
	// see https://blog.csdn.net/weiyuefei/article/details/77977734
	TotalCPUUsage    uint8 `json:"total_cpu_usage" bson:"total_cpu_usage"`
	TotalDSPUsage    uint8 `json:"total_dsp_usage" bson:"total_dsp_usage"`
	TotalGPUUsage    uint8 `json:"total_gpu_usage" bson:"total_gpu_usage"`
	TotalDiskUsage   uint8 `json:"total_disk_usage" bson:"total_disk_usage"`
	SwitchNumber     uint8 `json:"switch_number" bson:"switch_number"`
	RemoteUnitNumber uint8 `json:"remote_unit_number" bson:"remote_unit_number"`

	TotalRemoteUnitBytes uint8         `json:"total_remote_unit_bytes" bson:"total_remote_unit_bytes"`
	RemoteUnitSet        []*RemoteUnit `json:"remote_unit_set" bson:"remote_unit_set"`

	TotalSwitchDeviceBytes uint8           `json:"total_switch_device_bytes" bson:"total_switch_device_bytes"`
	SwitchDeviceSet        []*SwitchDevice `json:"switch_device_set" bson:"switch_device_set"`

	// cpu
	TotalCPUBytes uint8         `json:"total_cpu_bytes" bson:"total_cpu_bytes"`
	CPUSet        []*DeviceData `json:"cpu_set" bson:"cpu_set"` // 21bytes
	// dsp
	TotalDSPBytes uint8         `json:"total_dsp_bytes" bson:"total_dsp_bytes"`
	DSPSet        []*DeviceData `json:"dsp_set" bson:"dsp_set"` // 21bytes
	// gpu
	TotalGPUBytes uint8         `json:"total_gpu_bytes" bson:"total_gpu_bytes"`
	GPUSet        []*DeviceData `json:"gpu_set" bson:"gpu_set"` // 19bytes
	// fpga
	TotalFPGABytes uint8         `json:"total_fpga_bytes" bson:"total_fpga_bytes"`
	FPGASet        []*DeviceData `json:"fpga_set" bson:"fpga_set"` // 12bytes
	//system time
	SystemRunTime uint32 `json:"system_run_time" bson:"system_run_time"` //
	TimeUnit      uint8  `json:"time_unit" bson:"time_unit"`
	// app
	APPNum     uint8  `json:"app_num" bson:"app_num"`
	APPInfo    []*App `json:"app_info" bson:"app_info"`
	Sum        uint8  `json:"sum" bson:"sum"`
	Status     uint8  `json:"status" bson:"status"`
	IsTransfer bool   `json:"is_transfer" bson:"is_transfer"`
}

const cFailureOverTable = "failure_over_log"

func (controller *VMCController) SaveSelf() error {
	return mgm.CollectionByName("vmc_controller_data").Create(controller)
}

func (controller *VMCController) FindAndSetFailureEntity() error {
	uiq := controller.getControllerUnique()
	uniqFilter := bson.M{"unique_key": uiq}
	statusFilter := bson.M{"trans_status": 500}
	filter := bson.M{"$and": []bson.M{uniqFilter, statusFilter}}
	updateM := bson.M{"$set": bson.M{"to": FailureOverInfo{
		VMCID:    fmt.Sprintf("%d", controller.ToVmc),
		AppID:    fmt.Sprintf("%d", controller.Task.AppId),
		DeviceId: fmt.Sprintf("%d", controller.ToDevice),
	}, "trans_status": 200}}
	res, err := mgm.CollectionByName(cFailureOverTable).UpdateOne(mgm.Ctx(), filter, updateM)
	if err != nil {
		return err
	}
	glog.Infof("update result:%+v\n", res)
	return nil
}

func (controller *VMCController) getControllerUnique() string {
	trans := uint8(1)
	if controller.FromVmc != controller.ToVmc {
		trans = uint8(0)
	}

	return fmt.Sprintf("%d_%d_%d", trans, controller.FromVmc, controller.FromDevice)
}

func nameHandler(b_name []byte) string {
	name_len := bbytes.IndexByte(b_name[:], cSTR_END)
	if name_len < 0 {
		return string(b_name)
	}
	return string(b_name[:name_len])
}

func parseCPUDevice(bytes []byte, start, end int) ([]*DeviceData, uint8) {
	if end <= start {
		return nil, 0
	}
	bytes = bytes[start:end]
	l := len(bytes)
	s := binary.BigEndian.Uint16(bytes[0:2])
	total := bytes[2]
	glog.Infof("l:%d s:%d t:%d\n", l, s, total)
	if (l-3)%cCPU_SIZE == 0 && s == cCPU_START {
		num := (l - 3) / cCPU_SIZE
		arr := make([]*DeviceData, num)
		ss := 3
		for i := 0; i < num; i++ {
			si := i*cCPU_SIZE + ss
			glog.Infof("cpu index %d num:%d\n", si, num)
			DeviceData := &DeviceData{
				Name:                nameHandler(bytes[si : si+10]),
				ID:                  bytes[si+10],
				Type:                bytes[si+11],
				Num:                 bytes[si+12],
				IntComputingPower:   bytesFloat32(bytes[si+13 : si+17]),
				FloatComputingPower: bytesFloat32(bytes[si+17 : si+21]),
				TotalMemory:         binary.BigEndian.Uint16(bytes[si+21 : si+23]),
				MemoryUsage:         bytes[si+23],
				Usage:               bytes[si+24],
			}
			glog.Infof("cpu device %+v\n", DeviceData)
			arr[i] = DeviceData
		}

		return arr, total
	}
	return nil, 0
}

func bytesFloat32(bytes []byte) float32 {
	bits := binary.BigEndian.Uint32(bytes)
	float := math.Float32frombits(bits)
	return float
}

func parseGPUDevice(bytes []byte, start, end int) ([]*DeviceData, uint8) {
	if end <= start {
		return nil, 0
	}
	bytes = bytes[start:end]
	l := len(bytes)
	s := binary.BigEndian.Uint16(bytes[0:2])
	total := bytes[2]
	glog.Infof("l:%d s:%d t:%d\n", l, s, total)
	if s == cGPU_START && (l-3)%cGPU_SIZE == 0 {
		num := (l - 3) / cGPU_SIZE
		arr := make([]*DeviceData, num)
		ss := 3
		for j := 0; j < num; j++ {
			i := j*cGPU_SIZE + ss
			glog.Infof("gpu i:%d num:%d\n", i, num)
			DeviceData := &DeviceData{
				Name:                nameHandler(bytes[i : i+10]),
				ID:                  bytes[i+10],
				Type:                bytes[i+11],
				Num:                 bytes[i+12],
				IntComputingPower:   0,
				FloatComputingPower: bytesFloat32(bytes[i+13 : i+17]), //binary.BigEndian.Uint16(bytes[i+13 : i+15]),
				TotalMemory:         binary.BigEndian.Uint16(bytes[i+17 : i+19]),
				MemoryUsage:         bytes[i+19],
				Usage:               bytes[i+20],
			}
			glog.Infof("gpu: %+v\n", DeviceData)
			arr[j] = DeviceData
		}
		return arr, total
	}
	return nil, 0
}

func parseFPGADevice(bytes []byte, start, end int) ([]*DeviceData, uint8) {
	if end <= start {
		return nil, 0
	}
	bytes = bytes[start:end]
	l := len(bytes)
	s := binary.BigEndian.Uint16(bytes[0:2])
	total := bytes[2]
	glog.Infof("l:%d s:%d t:%d\n", l, s, total)
	ss := 3
	if (l-3)%cFPGA_SIZE == 0 && s == cFPGA_START {
		num := (l - 3) / cFPGA_SIZE
		arr := make([]*DeviceData, num)
		for j := 0; j < num; j++ {
			i := j*cFPGA_SIZE + ss
			glog.Infof("fpga i:%d num:%d\n", i, num)
			DeviceData := &DeviceData{
				Name: nameHandler(bytes[i : i+10]),
				ID:   bytes[i+10],
				Type: bytes[i+11],
			}
			glog.Infof("fpga:%+v", DeviceData)
			arr[j] = DeviceData
		}
		return arr, total
	}

	return nil, 0
}

func parseDSPDevice(bytes []byte, start, end int) ([]*DeviceData, uint8) {
	if end <= start {
		return nil, 0
	}
	bytes = bytes[start:end]
	l := len(bytes)
	s := binary.BigEndian.Uint16(bytes[0:2])
	total := bytes[2]
	glog.Infof("l:%d s:%d total:%d\n", l, s, total)
	if (l-3)%cDSP_SIZE == 0 && s == cDSP_START {
		num := (l - 3) / cDSP_SIZE
		arr := make([]*DeviceData, num)
		ss := 3
		for i := 0; i < num; i++ {
			si := i*cDSP_SIZE + ss
			glog.Infof("dsp index %d num:%d\n", si, num)
			DeviceData := &DeviceData{
				Name:                nameHandler(bytes[si : si+10]),
				ID:                  bytes[si+10],
				Type:                bytes[si+11],
				Num:                 bytes[si+12],
				IntComputingPower:   bytesFloat32(bytes[si+13 : si+17]),
				FloatComputingPower: bytesFloat32(bytes[si+17 : si+21]),
				TotalMemory:         binary.BigEndian.Uint16(bytes[si+21 : si+23]),
				MemoryUsage:         bytes[si+23],
				Usage:               bytes[si+24],
			}
			glog.Infof("dsp device %+v\n", DeviceData)
			arr[i] = DeviceData
		}
		return arr, total
	}
	return nil, 0
}

func parseTask(bytes []byte, start, end int) ([]*Task, uint8) {
	if end <= start {
		return nil, 1
	}
	bytes = bytes[start:end]
	length := len(bytes)
	glog.Infof("task length: %d", length)

	var statusCode uint8
	if length%cTAST_SIZE == 0 {
		taskNum := length / cTAST_SIZE
		arr := make([]*Task, taskNum)
		for j := 0; j < taskNum; j++ {
			i := j * cTAST_SIZE
			glog.Infof("task start index: %d", i)
			t := &Task{
				Name:        nameHandler(bytes[i : i+2]),
				ID:          binary.BigEndian.Uint16(bytes[i+2 : i+4]),
				TaskType:    bytes[i+4],
				TaskStatus:  bytes[i+5],
				ExecuteTime: binary.BigEndian.Uint32(bytes[i+6 : i+10]),
				StatusCode:  bytes[i+10],
				StartTime:   bytes[i+11],
			}
			arr[j] = t
			statusCode |= t.StatusCode
			glog.Infof("task details: %+v\n", t)
		}
		return arr, statusCode
	}

	return nil, 1
}

func parseApp(bytes []byte, start, end int) ([]*App, uint8) {
	if end <= start {
		return nil, 1
	}
	bytes = bytes[start:end]
	length := len(bytes)
	arr := make([]*App, 0)
	appStart := 0
	var vmcStatus uint8
	for {
		glog.Infof("app start: %d len:%d", appStart, length)
		if appStart >= length {
			break
		}
		name := nameHandler(bytes[appStart : appStart+10])
		taskNum := bytes[appStart+10]
		runPeriod := binary.BigEndian.Uint16(bytes[appStart+11 : appStart+13])
		dispatchTime := binary.BigEndian.Uint16(bytes[appStart+13 : appStart+15])
		id := bytes[appStart+15]
		resetNumber := bytes[appStart+16]
		vmcId := bytes[appStart+17]
		did := bytes[appStart+18]
		taskLen := uint8(taskNum) * cTAST_SIZE
		taskStart := appStart + 19
		taskEnd := taskStart + int(taskLen)
		glog.Infof("task start: %d, task end: %d", taskStart, taskEnd)
		taskSet, appStatus := parseTask(bytes, taskStart, taskEnd)
		vmcStatus |= appStatus
		a := &App{
			APPName:      name,
			TaskNum:      taskNum,
			TaskPeriod:   runPeriod,
			TaskDispatch: dispatchTime,
			ID:           id,
			ResetNumber:  resetNumber,
			BelongsTo:    vmcId,
			TaskSet:      taskSet,
			AppStatus:    appStatus,
			DeviceId:     did,
		}
		appStart = taskEnd
		glog.Infof("app details: %+v\n", a)
		arr = append(arr, a)
	}

	return arr, vmcStatus
}

func parseRemoteUnit(bytes []byte, start, end int) ([]*RemoteUnit, uint8) {
	if end <= start {
		return nil, 0
	}
	bytes = bytes[start:end]
	s := binary.BigEndian.Uint16(bytes[0:2])
	l := uint8(bytes[2])

	if s == cREMOTE_START && int(l) == len(bytes)-3 && l%cREMOTE_SIZE == 0 {
		num := l / cREMOTE_SIZE
		idx := 3
		arr := make([]*RemoteUnit, num)
		for i := 0; i < int(num); i++ {
			t := i*cREMOTE_SIZE + idx
			r := &RemoteUnit{
				RemoteUnitName:  nameHandler(bytes[t : t+10]),
				RemoteUnitOrder: bytes[t+10],
				RemoteUnitType:  bytes[t+11],
				LinkTo:          bytes[t+12],
			}
			arr[i] = r
			glog.Infof("remote unit %+v\n", bytes[t:t+13])
			glog.Infof("remote unit %+v\n", r)
		}

		return arr, l
	}

	return nil, 0
}

func parseSwitch(bytes []byte, start, end int) ([]*SwitchDevice, uint8) {
	if end <= start {
		return nil, 0
	}
	bytes = bytes[start:end]
	s := binary.BigEndian.Uint16(bytes[0:2])
	l := uint8(bytes[2])

	if s == cSWITCH_START && int(l) == len(bytes)-3 && l%cSWITCH_SIZE == 0 {
		num := l / cSWITCH_SIZE
		idx := 3
		arr := make([]*SwitchDevice, num)
		for i := 0; i < int(num); i++ {
			t := i*cSWITCH_SIZE + idx
			r := &SwitchDevice{
				SwitchName:  nameHandler(bytes[t : t+10]),
				SwitchOrder: bytes[t+10],
				SwitchType:  bytes[t+11],
				LinkTo:      bytes[t+12],
			}
			arr[i] = r
			glog.Infof("switch unit %+v\n", bytes[t:t+13])
			glog.Infof("switch unit %+v\n", r)
		}

		return arr, l
	}

	return nil, 0
}

func calcStartEnd(start int, num uint8, l int) (int, int) {
	if num == 0 {
		glog.Infof("num is zero return same as start\n")
		return start, start
	}

	glog.Infoln(fmt.Sprintf("s:%d, n:%d, l:%d", start, num, l))
	return start, start + 1 + 1 + l*int(num) + 1
}

// todo
func parse(bytes []byte) (*VMCData, error) {
	l := len(bytes)
	deviceIdx := 31
	remoteStart, remoteEnd := calcStartEnd(deviceIdx, uint8(bytes[30]), 13)
	glog.Infof("remote unit start:%d cpu end:%d\n", remoteStart, remoteEnd)
	switchStart, switchEnd := calcStartEnd(remoteEnd, uint8(bytes[29]), 13)
	glog.Infof("switch start:%d cpu end:%d\n", switchStart, switchEnd)
	cpuStart, cpuEnd := calcStartEnd(switchEnd, uint8(bytes[15]), cCPU_SIZE)
	glog.Infof("cpu start:%d cpu end:%d\n", cpuStart, cpuEnd)
	dspStart, dspEnd := calcStartEnd(cpuEnd, uint8(bytes[16]), cDSP_SIZE)
	glog.Infof("dsp start:%d dsp end:%d\n", dspStart, dspEnd)
	gpusStart, gpusEnd := calcStartEnd(dspEnd, uint8(bytes[17]), cGPU_SIZE) // no int computer power
	glog.Infof("gpu start:%d gpus end:%d\n", gpusStart, gpusEnd)
	fpgaStart, fpgaEnd := calcStartEnd(gpusEnd, uint8(bytes[18]), cFPGA_SIZE)
	glog.Infof("fpga start:%d fpga end:%d\n", fpgaStart, fpgaEnd)

	sysRunTimeStart := remoteEnd
	if remoteStart < fpgaEnd {
		sysRunTimeStart = fpgaEnd
	}

	remoteSet, totalRemoteBytes := parseRemoteUnit(bytes, remoteStart, remoteEnd)
	switchSet, totalSwitchDeviceBytes := parseSwitch(bytes, switchStart, switchEnd)
	cpuSet, totalCpuBytes := parseCPUDevice(bytes, cpuStart, cpuEnd)
	dspSet, totalDspDeviceBytes := parseDSPDevice(bytes, dspStart, dspEnd)
	gpuSet, totalGpuBytes := parseGPUDevice(bytes, gpusStart, gpusEnd)
	fpagSet, totalFpagBytes := parseFPGADevice(bytes, fpgaStart, fpgaEnd)
	// sys time
	sysRunTimeEnd := sysRunTimeStart + 4
	timeUnitStart := sysRunTimeEnd
	sysTime := binary.BigEndian.Uint32(bytes[sysRunTimeStart:sysRunTimeEnd])
	tu := bytes[timeUnitStart]
	glog.Infof("sysRunTimeStart idx: %d, timeUnitStart idx: %d systime:%d tu:%d\n", sysRunTimeStart, timeUnitStart, sysTime, tu)
	// app skip sum
	appIdx := timeUnitStart + 2
	glog.Infof("app idx: %d, app num: %d\n", appIdx, bytes[appIdx])

	appSet, vmcStatus := parseApp(bytes, appIdx+1, l-1)

	v := &VMCData{
		frameHeader:      bytes[0],
		length:           binary.BigEndian.Uint16(bytes[1:3]),
		protoType:        bytes[3],
		VMCName:          nameHandler(bytes[4:14]),
		VMCID:            bytes[14],
		CPUNumber:        bytes[15],
		DSPNumber:        bytes[16],
		GPUNumber:        bytes[17],
		FPGANumber:       bytes[18],
		SwitchID:         bytes[19],
		TotalMemory:      binary.BigEndian.Uint16(bytes[20:22]),
		TotalDisk:        binary.BigEndian.Uint16(bytes[22:24]),
		MemoryUsage:      bytes[24],
		TotalCPUUsage:    bytes[25],
		TotalDSPUsage:    bytes[26],
		TotalGPUUsage:    bytes[27],
		TotalDiskUsage:   bytes[28],
		SwitchNumber:     bytes[29],
		RemoteUnitNumber: bytes[30],

		TotalRemoteUnitBytes: totalRemoteBytes,
		RemoteUnitSet:        remoteSet,

		TotalSwitchDeviceBytes: totalSwitchDeviceBytes,
		SwitchDeviceSet:        switchSet,

		TotalCPUBytes: totalCpuBytes,
		CPUSet:        cpuSet,

		TotalDSPBytes: totalDspDeviceBytes,
		DSPSet:        dspSet,

		TotalGPUBytes: totalGpuBytes,
		GPUSet:        gpuSet,

		TotalFPGABytes: totalFpagBytes,
		FPGASet:        fpagSet,

		SystemRunTime: sysTime,
		TimeUnit:      tu,

		APPNum:  bytes[appIdx],
		APPInfo: appSet,
		Sum:     bytes[l-1],
		Status:  vmcStatus,
	}
	glog.Infof("debug vmc:%+v\n", v)
	return v, nil
}

type VMCDataJson struct {
	VMCName        string `json:"vmc_name" bson:"vmc_name"` // 10bytes
	VMCID          uint8  `json:"vmc_id" bson:"vmc_id"`
	CPUNumber      uint8  `json:"cpu_number" bson:"cpu_number"`
	DSPNumber      uint8  `json:"dsp_number" bson:"dsp_number"`
	GPUNumber      uint8  `json:"gpu_number" bson:"gpu_number"`
	FPGANumber     uint8  `json:"fpga_number" bson:"fpga_number"`
	SwitchID       uint8  `json:"switch_id" bson:"switch_id"`
	TotalMemory    uint16 `json:"total_memory" bson:"total_memory"`
	TotalDisk      uint16 `json:"total_disk" bson:"total_disk"`
	MemoryUsage    uint8  `json:"memory_usage" bson:"memory_usage"`
	TotalCPUUsage  uint8  `json:"total_cpu_usage" bson:"total_cpu_usage"`
	TotalDSPUsage  uint8  `json:"total_dsp_usage" bson:"total_dsp_usage"`
	TotalGPUUsage  uint8  `json:"total_gpu_usage" bson:"total_gpu_usage"`
	TotalDiskUsage uint8  `json:"total_disk_usage" bson:"total_disk_usage"`
}

func (src *VMCData) TransferVMCDataToJson() *VMCDataJson {
	if src == nil {
		return nil
	}
	dst := &VMCDataJson{}
	dst.VMCName = src.VMCName
	dst.VMCID = src.VMCID
	dst.CPUNumber = src.CPUNumber
	dst.DSPNumber = src.DSPNumber
	dst.GPUNumber = src.GPUNumber
	dst.FPGANumber = src.FPGANumber
	dst.SwitchID = src.SwitchID
	dst.TotalMemory = src.TotalMemory
	dst.TotalDisk = src.TotalDisk
	dst.MemoryUsage = src.MemoryUsage
	dst.TotalCPUUsage = src.TotalCPUUsage
	dst.TotalDSPUsage = src.TotalDSPUsage
	dst.TotalGPUUsage = src.TotalGPUUsage
	dst.TotalDiskUsage = src.TotalDiskUsage

	return dst
}

func NewVMCController(raw []byte) (*VMCController, error) {
	if len(raw) != 14 {
		return nil, fmt.Errorf("invalidate data length: %d", len(raw))
	}
	controller := &VMCController{
		FrameHeader: raw[0],
		Length:      binary.BigEndian.Uint16(raw[1:3]),
		DataType:    raw[3],
		FromVmc:     raw[4],
		FromDevice:  raw[5],
		ToVmc:       raw[6],
		ToDevice:    raw[7],
		Task: FailureTask{
			FailureName: binary.BigEndian.Uint16(raw[8:10]),
			FailureDesc: raw[10],
			AppId:       raw[11],
			IsFault:     raw[12],
		},
		Sum: raw[13],
	}

	return controller, nil
}

// read bytes from udp
func NewVMCData(str string) (*VMCData, error) {
	return parse([]byte(str))
}

func (vmc_data *VMCData) CreateData() error {
	if vmc_data == nil {
		return fmt.Errorf("vcm data is nil")
	}

	coll := mgm.CollectionByName(config.CommonConfig.DBVMCDataTableName)
	count, err := coll.CountDocuments(context.TODO(), bson.M{})
	if err != nil {
		return fmt.Errorf("CountDocuments error")
	}

	if count >= cMAX_DOCUMENT_NUM {
		v := &VMCData{}
		err := coll.First(bson.M{"vmc_id": vmc_data.VMCID}, v)
		if err != nil {
			glog.Infof("Cannot find vmc_id: %d", vmc_data.VMCID)
			err = coll.First(bson.M{}, v)
			if err != nil {
				glog.Errorf("cannot find any document, maybe the db is empty")
				return fmt.Errorf("cannot find any document, maybe the db is empty")
			}
		}
		glog.Infof("new v is: %+v", v)
		err = coll.Delete(v)
		if err != nil {
			return fmt.Errorf("delete vmc_data failed")
		}
	}
	return coll.Create(vmc_data)
}

func (vmc_data *VMCData) CollectVMCData(vmc_id int32) error {
	if vmc_data == nil {
		return fmt.Errorf("vcm data is nil need make one")
	}

	vmcs, err := vmc_data.GetVMCList(vmc_id)

	if len(vmcs) < 1 {
		return fmt.Errorf("vcm data is empty")
	} else {
		*vmc_data = *vmcs[0]
	}

	return err
	//return mgm.CollectionByName(config.CommonConfig.DBVMCDataTableName).First(bson.M{}, vmc_data, &options.FindOneOptions{Sort: bson.M{"_id": -1}})
}

func CollectDeviceData(vmc_id int32, device_type string) ([]*DeviceData, error) {
	vmc_data := &VMCData{}
	err := vmc_data.CollectVMCData(vmc_id)

	device_data := []*DeviceData{}
	switch device_type {
	case "cpu":
		device_data = vmc_data.CPUSet
	case "gpu":
		device_data = vmc_data.GPUSet
	case "fpga":
		device_data = vmc_data.FPGASet
	case "dsp":
		device_data = vmc_data.DSPSet
	}

	return device_data, err
}

func CollectAppInfo(vmc_id int32) ([]*App, error) {
	vmc_data := &VMCData{}
	err := vmc_data.CollectVMCData(vmc_id)

	return vmc_data.APPInfo, err
}

func (vmc_data *VMCData) GetVMCList(vmcid int32) ([]*VMCData, error) {
	vmcs := make([]*VMCData, 0)
	if vmc_data == nil {
		return vmcs, fmt.Errorf("vcm data is nil need make one")
	}

	findOptions := options.Find()
	findOptions.SetSort(bson.D{{Key: "updated_at", Value: -1}})
	ret := mgm.CollectionByName(config.CommonConfig.DBVMCDataTableName).SimpleFind(&vmcs, bson.M{"vmc_id": vmcid}, findOptions)
	if len(vmcs) < 1 {
		return vmcs, fmt.Errorf("vcms is empty")
	}
	return vmcs, ret
}
