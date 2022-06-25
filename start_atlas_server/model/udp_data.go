package model

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

// todo
func parseTask(bytes []byte) (*Task, error) {
	return nil, nil
}

// todo
func parse(bytes []byte) (*VMCData, error) {

	return nil, nil
}

// read bytes from udp
func NewVMCData(str string) (*VMCData, error) {
	return parse([]byte(str))
}
