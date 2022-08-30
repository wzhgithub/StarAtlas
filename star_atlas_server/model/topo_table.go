package model

import (
	"fmt"
	"star_atlas_server/config"
	"time"

	"github.com/golang/glog"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
)

const (
	CTopoID           = "topo_table"
	CVMCBase          = 10
	CDdeviceTypeShift = 5
)

// var appStatus = []string{"TIMEOUT", "ERROR", "RUN", "ERROR"}
// var vmcStatus = []string{"TIMEOUT", "ERROR", "RUN", "RUN"}
var vmcStatus = []string{"RUN", "RUN", "RUN", "RUN"}

type TransferInfos struct {
	FromId     uint16        `json:"from_id" bson:"from_id"`
	ToId       uint16        `json:"to_id" bson:"to_id"`
	TaskType   uint8         `json:"task_type" bson:"task_type"`
	StartTime  time.Time     `json:"start_time" bson:"start_time"`
	EndTime    time.Time     `json:"end_time" bson:"end_time"`
	DuringTime time.Duration `json:"during_time" bson:"during_time"`
}

type OtherInfos struct {
	Key    string   `json:"key" bson:"key"`
	Values []string `json:"value" bson:"value"`
}

type Nodes struct {
	Id           int64         `json:"id" bson:"id"`
	Name         string        `json:"name" bson:"name"`
	DeviceType   string        `json:"device_type" bson:"device_type"`
	ParentId     uint16        `json:"parent_id" bson:"parent_id"`
	UpstreamId   uint16        `json:"upstream_id" bson:"upstream_id"`
	DeviceStatus string        `json:"device_status" bson:"device_status"`
	DeviceNum    int32         `json:"device_num" bson:"device_num"`
	OtherInfo    []*OtherInfos `json:"other_info" bson:"other_info"`
}

// see https://github.com/Kamva/mgm
type TopoTable struct {
	mgm.DefaultModel `json:",inline" bson:",inline"`
	Id               string           `json:"id" bson:"id"`
	Node             []*Nodes         `json:"node" bson:"node"`
	TransferInfo     []*TransferInfos `json:"transfer_info" bson:"transfer_info"`
}

func NewOtherInfos(key string, vals []string) *OtherInfos {
	return &OtherInfos{key, vals}
}

type pNodesArr []*Nodes

func (v *VMCData) parseCPU(nodes *pNodesArr) {
	if v.CPUNumber == 0 {
		return
	}
	n := &Nodes{
		Id:           int64(v.VMCID)*CVMCBase + int64(v.CPUSet[0].ID>>CDdeviceTypeShift),
		Name:         "SOC-2018B",
		DeviceType:   "cpu",
		ParentId:     uint16(v.VMCID),
		UpstreamId:   0,
		DeviceStatus: vmcStatus[v.Status],
		DeviceNum:    int32(v.CPUNumber),
		OtherInfo:    make([]*OtherInfos, 0),
	}
	cpu_ids := make([]string, 0)
	cpu_names := make([]string, 0)
	cpu_types := make([]string, 0)
	cpu_cores := make([]string, 0)
	for i := 0; i < int(v.CPUNumber); i++ {
		cpu_ids = append(cpu_ids, fmt.Sprintf("%d", v.CPUSet[i].ID))
		cpu_names = append(cpu_names, v.CPUSet[i].Name)
		cpu_types = append(cpu_types, fmt.Sprintf("%d", v.CPUSet[i].Type))
		cpu_cores = append(cpu_cores, fmt.Sprintf("%d", v.CPUSet[i].Num))
	}
	n.OtherInfo = append(n.OtherInfo, NewOtherInfos("cpu_ids", cpu_ids))
	n.OtherInfo = append(n.OtherInfo, NewOtherInfos("cpu_names", cpu_names))
	n.OtherInfo = append(n.OtherInfo, NewOtherInfos("cpu_types", cpu_types))
	n.OtherInfo = append(n.OtherInfo, NewOtherInfos("cpu_cores", cpu_cores))
	*nodes = append(*nodes, n)
}

func (v *VMCData) parseGPU(nodes *pNodesArr) {
	if v.GPUNumber == 0 {
		return
	}
	n := &Nodes{
		Id:           int64(v.VMCID)*CVMCBase + int64(v.GPUSet[0].ID>>CDdeviceTypeShift),
		Name:         "NVIDIA-AGX",
		DeviceType:   "gpu",
		ParentId:     uint16(v.VMCID),
		UpstreamId:   0,
		DeviceStatus: vmcStatus[v.Status],
		DeviceNum:    int32(v.GPUNumber),
		OtherInfo:    make([]*OtherInfos, 0),
	}
	gpu_ids := make([]string, 0)
	gpu_names := make([]string, 0)
	gpu_types := make([]string, 0)
	gpu_cores := make([]string, 0)
	for i := 0; i < int(v.GPUNumber); i++ {
		gpu_ids = append(gpu_ids, fmt.Sprintf("%d", v.GPUSet[i].ID))
		gpu_names = append(gpu_names, v.GPUSet[i].Name)
		gpu_types = append(gpu_types, fmt.Sprintf("%d", v.GPUSet[i].Type))
		gpu_cores = append(gpu_cores, fmt.Sprintf("%d", v.GPUSet[i].Num))
	}
	n.OtherInfo = append(n.OtherInfo, NewOtherInfos("gpu_ids", gpu_ids))
	n.OtherInfo = append(n.OtherInfo, NewOtherInfos("gpu_names", gpu_names))
	n.OtherInfo = append(n.OtherInfo, NewOtherInfos("gpu_types", gpu_types))
	n.OtherInfo = append(n.OtherInfo, NewOtherInfos("gpu_cores", gpu_cores))
	*nodes = append(*nodes, n)
}

func (v *VMCData) parseDSP(nodes *pNodesArr) {
	if v.DSPNumber == 0 {
		return
	}
	n := &Nodes{
		Id:           int64(v.VMCID)*CVMCBase + int64(v.DSPSet[0].ID>>CDdeviceTypeShift),
		Name:         "FT-6678",
		DeviceType:   "dsp",
		ParentId:     uint16(v.VMCID),
		UpstreamId:   0,
		DeviceStatus: vmcStatus[v.Status],
		DeviceNum:    int32(v.DSPNumber),
		OtherInfo:    make([]*OtherInfos, 0),
	}
	dsp_ids := make([]string, 0)
	dsp_names := make([]string, 0)
	dsp_types := make([]string, 0)
	dsp_cores := make([]string, 0)
	for i := 0; i < int(v.DSPNumber); i++ {
		dsp_ids = append(dsp_ids, fmt.Sprintf("%d", v.DSPSet[i].ID))
		dsp_names = append(dsp_names, v.DSPSet[i].Name)
		dsp_types = append(dsp_types, fmt.Sprintf("%d", v.DSPSet[i].Type))
		dsp_cores = append(dsp_cores, fmt.Sprintf("%d", v.DSPSet[i].Num))
	}
	n.OtherInfo = append(n.OtherInfo, NewOtherInfos("dsp_ids", dsp_ids))
	n.OtherInfo = append(n.OtherInfo, NewOtherInfos("dsp_names", dsp_names))
	n.OtherInfo = append(n.OtherInfo, NewOtherInfos("dsp_types", dsp_types))
	n.OtherInfo = append(n.OtherInfo, NewOtherInfos("dsp_cores", dsp_cores))
	*nodes = append(*nodes, n)
}

func (v *VMCData) parseFPGA(nodes *pNodesArr) {
	if v.FPGANumber == 0 {
		return
	}
	n := &Nodes{
		Id:           int64(v.VMCID)*CVMCBase + int64(v.FPGASet[0].ID>>CDdeviceTypeShift),
		Name:         "V7-690T",
		DeviceType:   "fpga",
		ParentId:     uint16(v.VMCID),
		UpstreamId:   0,
		DeviceStatus: vmcStatus[v.Status],
		DeviceNum:    int32(v.FPGANumber),
		OtherInfo:    make([]*OtherInfos, 0),
	}
	fpga_ids := make([]string, 0)
	fpga_names := make([]string, 0)
	fpga_types := make([]string, 0)
	fpga_cores := make([]string, 0)
	for i := 0; i < int(v.FPGANumber); i++ {
		fpga_ids = append(fpga_ids, fmt.Sprintf("%d", v.FPGASet[i].ID))
		fpga_names = append(fpga_names, v.FPGASet[i].Name)
		fpga_types = append(fpga_types, fmt.Sprintf("%d", v.FPGASet[i].Type))
		fpga_cores = append(fpga_cores, fmt.Sprintf("%d", v.FPGASet[i].Num))
	}
	n.OtherInfo = append(n.OtherInfo, NewOtherInfos("fpga_ids", fpga_ids))
	n.OtherInfo = append(n.OtherInfo, NewOtherInfos("fpga_names", fpga_names))
	n.OtherInfo = append(n.OtherInfo, NewOtherInfos("fpga_types", fpga_types))
	n.OtherInfo = append(n.OtherInfo, NewOtherInfos("fpga_cores", fpga_cores))
	*nodes = append(*nodes, n)
}

func (v *VMCData) parseVMC(nodes *pNodesArr) {
	v.parseCPU(nodes)
	v.parseGPU(nodes)
	v.parseDSP(nodes)
	v.parseFPGA(nodes)
	n := &Nodes{
		Id:           int64(v.VMCID),
		Name:         v.VMCName,
		DeviceType:   "vmc",
		ParentId:     0,
		UpstreamId:   uint16(v.SwitchID),
		DeviceStatus: vmcStatus[v.Status],
		OtherInfo:    make([]*OtherInfos, 0),
	}
	n.OtherInfo = append(n.OtherInfo, NewOtherInfos("proto_type", []string{fmt.Sprintf("%d", v.protoType)}))
	*nodes = append(*nodes, n)
}

func (v *VMCData) parseSwitch(nodes *pNodesArr) {
	for i := 0; i < int(v.SwitchNumber); i++ {
		n := &Nodes{
			Id:           int64(v.SwitchDeviceSet[i].SwitchOrder),
			Name:         v.SwitchDeviceSet[i].SwitchName,
			DeviceType:   "sw",
			ParentId:     0,
			UpstreamId:   uint16(v.SwitchDeviceSet[i].LinkTo),
			DeviceStatus: "RUN",
			OtherInfo:    make([]*OtherInfos, 0),
		}
		(n.OtherInfo) = append(n.OtherInfo, NewOtherInfos("switch_type", []string{fmt.Sprintf("%d", v.SwitchDeviceSet[i].SwitchType)}))
		*nodes = append(*nodes, n)
	}
}

func (v *VMCData) parseRTU(nodes *pNodesArr) {
	for i := 0; i < int(v.RemoteUnitNumber); i++ {
		n := &Nodes{
			Id:           int64(v.RemoteUnitSet[i].RemoteUnitOrder),
			Name:         v.RemoteUnitSet[i].RemoteUnitName,
			DeviceType:   "rtu",
			ParentId:     0,
			UpstreamId:   uint16(v.RemoteUnitSet[i].LinkTo),
			DeviceStatus: "RUN",
			OtherInfo:    make([]*OtherInfos, 0),
		}
		n.OtherInfo = append(n.OtherInfo, NewOtherInfos("rtu_type", []string{fmt.Sprintf("%d", v.RemoteUnitSet[i].RemoteUnitType)}))
		*nodes = append(*nodes, n)
	}
}

func NewNodes(v *VMCData) pNodesArr {
	nodes := make(pNodesArr, 0)
	v.parseVMC(&nodes)
	v.parseSwitch(&nodes)
	v.parseRTU(&nodes)
	return nodes
}

func NewTransferInfos(v *VMCData) []*TransferInfos {
	return nil
}

func NewTopoTable(v *VMCData) *TopoTable {
	return &TopoTable{
		Id:           CTopoID,
		Node:         NewNodes(v),
		TransferInfo: NewTransferInfos(v),
	}
}

func (t *TopoTable) CreateOp(v *VMCData) error {
	err := mgm.CollectionByName(config.CommonConfig.DBTopoTableName).First(bson.M{"id": CTopoID}, t)
	if err != nil {
		glog.Infof("[CreateOp] Cannot find, create a new topo_table")
	}
	glog.Infof("[CreateOp] topo id = %s", t.Id)
	if t.Id == "" {
		glog.Infof("[CreateOp] new topoTable")
		t = NewTopoTable(v)
		glog.Infof("[CreateOp] new topo: %+v", t)
		return mgm.CollectionByName(config.CommonConfig.DBTopoTableName).Create(t)
	}

	newNodes := make([]*Nodes, 0)
	for _, node := range t.Node {
		if node.Id != int64(v.VMCID) && node.ParentId != uint16(v.VMCID) && node.DeviceType != "sw" && node.DeviceType != "rtu" {
			newNodes = append(newNodes, node)
		}
	}
	newNodes = append(newNodes, NewNodes(v)...)
	t.Node = newNodes
	return t.UpdateOp()
}

func (t *TopoTable) CollectOp() error {
	return mgm.CollectionByName(config.CommonConfig.DBTopoTableName).First(bson.M{}, t)
}

func (t *TopoTable) UpdateOp() error {
	return mgm.CollectionByName(config.CommonConfig.DBTopoTableName).Update(t)
}

func (t *TopoTable) InsertOp(node *Nodes) error {
	err := t.CollectOp()
	if err != nil {
		return fmt.Errorf("[InsertOp] Find error err:%+v", err)
	}
	t.Node = append(t.Node, node)
	return t.UpdateOp()
}

func (t *TopoTable) DeleteOp(delIds []int64) error {
	glog.Infof("[DeleteOp] ids: %+v", delIds)
	newNodes := make([]*Nodes, 0)
	for _, node := range t.Node {
		newAdded := true
		for _, id := range delIds {
			if node.Id == id {
				newAdded = false
				break
			}
		}
		if newAdded {
			newNodes = append(newNodes, node)
		}
	}
	glog.Infof("[DeleteOp] t.node length: %d", len(t.Node))
	glog.Infof("[DeleteOp] newNodes length: %d", len(newNodes))
	t.Node = newNodes
	return t.UpdateOp()
}

func (t *TopoTable) GetMaxVmcId() (int64, error) {
	err := t.CollectOp()
	if err != nil {
		glog.Errorf("[DeleteOp] Find error err:%+v", err)
		return -1, err
	}
	var maxId int64 = -1
	for _, node := range t.Node {
		if node.DeviceType == "vmc" {
			if node.Id > maxId {
				maxId = node.Id
			}
		}
	}
	if maxId == -1 {
		return -1, fmt.Errorf("cannot vmc device in the topo_table")
	}
	return maxId, nil
}

func (t *TopoTable) GetVmcStatus(vmc_id int64) (string, error) {
	for _, node := range t.Node {
		if node.Id == vmc_id {
			return node.DeviceStatus, nil
		}
	}
	return "", fmt.Errorf("vmc_id: %d is not found in topo_table", vmc_id)
}

func (t *TopoTable) GetBackupId(vmc_id int64) ([]int64, error) {
	backupId := make([]int64, 0)
	var switchId uint16
	for _, node := range t.Node {
		if node.Id == vmc_id {
			switchId = node.UpstreamId
		}
	}
	for _, node := range t.Node {
		if node.UpstreamId == switchId &&
			node.DeviceType == "vmc" &&
			node.Id != vmc_id {
			backupId = append(backupId, node.Id)
		}
	}
	if len(backupId) > 0 {
		return backupId, nil
	}
	return backupId, fmt.Errorf("vmc_id: %d doesn't have other vmcs in the same switch: %d", vmc_id, switchId)
}
