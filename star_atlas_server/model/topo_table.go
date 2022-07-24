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
	CTopoID  = "topo_table"
	CVMCBase = int64(10e6)
	CRTUBase = int64(10e3)
)

// var appStatus = []string{"TIMEOUT", "ERROR", "RUN", "ERROR"}
var vmcStatus = []string{"TIMEOUT", "ERROR", "RUN", "RUN"}

type TransferInfos struct {
	FromId     uint16        `json:"from_id" bson:"from_id"`
	ToId       uint16        `json:"to_id" bson:"to_id"`
	TaskType   uint8         `json:"task_type" bson:"task_type"`
	StartTime  time.Time     `json:"start_time" bson:"start_time"`
	EndTime    time.Time     `json:"end_time" bson:"end_time"`
	DuringTime time.Duration `json:"during_time" bson:"during_time"`
}

type OtherInfos struct {
	Key   string `json:"key" bson:"key"`
	Value string `json:"value" bson:"value"`
}

type Nodes struct {
	Id           int64         `json:"id" bson:"id"`
	Name         string        `json:"name" bson:"name"`
	DeviceType   string        `json:"device_type" bson:"device_type"`
	ParentId     uint16        `json:"parent_id" bson:"parent_id"`
	UpstreamId   uint16        `json:"upstream_id" bson:"upstream_id"`
	DeviceStatus string        `json:"device_status" bson:"device_status"`
	OtherInfo    []*OtherInfos `json:"other_info" bson:"other_info"`
}

// see https://github.com/Kamva/mgm
type TopoTable struct {
	mgm.DefaultModel `json:",inline" bson:",inline"`
	Id               string           `json:"id" bson:"id"`
	Node             []*Nodes         `json:"node" bson:"node"`
	TransferInfo     []*TransferInfos `json:"transfer_info" bson:"transfer_info"`
}

func NewOtherInfos(key string, val string) *OtherInfos {
	return &OtherInfos{key, val}
}

type pNodesArr []*Nodes

func (v *VMCData) parseCPU(nodes *pNodesArr) {
	for i := 0; i < int(v.CPUNumber); i++ {
		n := &Nodes{
			Id:           int64(v.VMCID)*CVMCBase + int64(v.CPUSet[i].ID),
			Name:         v.CPUSet[i].Name,
			DeviceType:   "cpu",
			ParentId:     uint16(v.VMCID),
			UpstreamId:   0,
			DeviceStatus: vmcStatus[v.Status],
			OtherInfo:    make([]*OtherInfos, 0),
		}
		n.OtherInfo = append(n.OtherInfo, NewOtherInfos("cpu_type", fmt.Sprintf("%d", v.CPUSet[i].Type)))
		n.OtherInfo = append(n.OtherInfo, NewOtherInfos("cpu_cores", fmt.Sprintf("%d", v.CPUSet[i].Num)))
		*nodes = append(*nodes, n)
	}
}

func (v *VMCData) parseGPU(nodes *pNodesArr) {
	for i := 0; i < int(v.GPUNumber); i++ {
		n := &Nodes{
			Id:           int64(v.VMCID)*CVMCBase + int64(v.GPUSet[i].ID),
			Name:         v.GPUSet[i].Name,
			DeviceType:   "gpu",
			ParentId:     uint16(v.VMCID),
			UpstreamId:   0,
			DeviceStatus: vmcStatus[v.Status],
			OtherInfo:    make([]*OtherInfos, 0),
		}
		n.OtherInfo = append(n.OtherInfo, NewOtherInfos("gpu_type", fmt.Sprintf("%d", v.GPUSet[i].Type)))
		n.OtherInfo = append(n.OtherInfo, NewOtherInfos("gpu_cores", fmt.Sprintf("%d", v.GPUSet[i].Num)))
		*nodes = append(*nodes, n)
	}
}

func (v *VMCData) parseDSP(nodes *pNodesArr) {
	for i := 0; i < int(v.DSPNumber); i++ {
		n := &Nodes{
			Id:           int64(v.VMCID)*CVMCBase + int64(v.DSPSet[i].ID),
			Name:         v.DSPSet[i].Name,
			DeviceType:   "dsp",
			ParentId:     uint16(v.VMCID),
			UpstreamId:   0,
			DeviceStatus: vmcStatus[v.Status],
			OtherInfo:    make([]*OtherInfos, 0),
		}
		n.OtherInfo = append(n.OtherInfo, NewOtherInfos("dsp_type", fmt.Sprintf("%d", v.DSPSet[i].Type)))
		n.OtherInfo = append(n.OtherInfo, NewOtherInfos("dsp_cores", fmt.Sprintf("%d", v.DSPSet[i].Num)))
		*nodes = append(*nodes, n)
	}
}

func (v *VMCData) parseFPGA(nodes *pNodesArr) {
	for i := 0; i < int(v.FPGANumber); i++ {
		n := &Nodes{
			Id:           int64(v.VMCID)*CVMCBase + int64(v.FPGASet[i].ID),
			Name:         v.FPGASet[i].Name,
			DeviceType:   "fpga",
			ParentId:     uint16(v.VMCID),
			UpstreamId:   0,
			DeviceStatus: vmcStatus[v.Status],
			OtherInfo:    make([]*OtherInfos, 0),
		}
		n.OtherInfo = append(n.OtherInfo, NewOtherInfos("fpga_type", fmt.Sprintf("%d", v.FPGASet[i].Type)))
		n.OtherInfo = append(n.OtherInfo, NewOtherInfos("fpga_cores", fmt.Sprintf("%d", v.FPGASet[i].Num)))
		*nodes = append(*nodes, n)
	}
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
	n.OtherInfo = append(n.OtherInfo, NewOtherInfos("proto_type", fmt.Sprintf("%d", v.protoType)))
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
		(n.OtherInfo) = append(n.OtherInfo, NewOtherInfos("switch_type", fmt.Sprintf("%d", v.SwitchDeviceSet[i].SwitchType)))
		*nodes = append(*nodes, n)
	}
}

func (v *VMCData) parseRTU(nodes *pNodesArr) {
	for i := 0; i < int(v.RemoteUnitNumber); i++ {
		n := &Nodes{
			Id:           int64(v.RemoteUnitSet[i].LinkTo)*CRTUBase + int64(v.RemoteUnitSet[i].RemoteUnitOrder),
			Name:         v.RemoteUnitSet[i].RemoteUnitName,
			DeviceType:   "rtu",
			ParentId:     0,
			UpstreamId:   uint16(v.RemoteUnitSet[i].LinkTo),
			DeviceStatus: "RUN",
			OtherInfo:    make([]*OtherInfos, 0),
		}
		n.OtherInfo = append(n.OtherInfo, NewOtherInfos("rtu_type", fmt.Sprintf("%d", v.RemoteUnitSet[i].RemoteUnitType)))
		*nodes = append(*nodes, n)
	}
}

func NewNodes(v *VMCData, isFirst bool) pNodesArr {
	nodes := make(pNodesArr, 0)
	v.parseVMC(&nodes)
	if isFirst {
		v.parseSwitch(&nodes)
		v.parseRTU(&nodes)
	}
	return nodes
}

func NewTransferInfos(v *VMCData) []*TransferInfos {
	return nil
}

func NewTopoTable(v *VMCData, isFirst bool) *TopoTable {
	return &TopoTable{
		Id:           CTopoID,
		Node:         NewNodes(v, isFirst),
		TransferInfo: NewTransferInfos(v),
	}
}

func (t *TopoTable) CreateOp(v *VMCData) error {
	err := mgm.CollectionByName(config.CommonConfig.DBTopoTableName).First(bson.M{"id": CTopoID}, t)
	if err != nil {
		glog.Infof("[CreateOp] Cannot find, create a new topo_table")
	}
	glog.Infof("[CreateOp] find return: %+v", t)
	if t.Id == "" {
		glog.Infof("[CreateOp] new topoTable")
		t = NewTopoTable(v, true)
		return mgm.CollectionByName(config.CommonConfig.DBTopoTableName).Create(t)
	}
	t.Node = append(t.Node, NewNodes(v, false)...)
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

func (t *TopoTable) DeleteOp(id int64) error {
	err := t.CollectOp()
	if err != nil {
		return fmt.Errorf("[DeleteOp] Find error err:%+v", err)
	}
	var index int
	for idx, node := range t.Node {
		if node.Id == id {
			index = idx
			break
		}
	}
	t.Node = append(t.Node[:index], t.Node[index+1:]...)
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
