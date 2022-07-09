package model

import (
	"star_atlas_server/config"
	"time"

	"github.com/golang/glog"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
)

const (
	cTopoID = "topo_table"
)

type TransferInfos struct {
	FromId      uint16        `json:"from_id" bson:"from_id"`
	ToId        uint16        `json:"to_id" bson:"to_id"`
	TaskType    uint8         `json:"task_type" bson:"task_type"`
	StartTime   time.Time     `json:"start_time" bson:"start_time"`
	EndTime     time.Time     `json:"end_time" bson:"end_time"`
	DurningTime time.Duration `json:"durning_time" bson:"durning_time"`
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
			Id:           int64(v.CPUSet[i].ID),
			Name:         v.CPUSet[i].Name,
			DeviceType:   "cpu",
			ParentId:     uint16(v.VMCID),
			UpstreamId:   0,
			DeviceStatus: "RUN",
			OtherInfo:    make([]*OtherInfos, 0),
		}
		n.OtherInfo = append(n.OtherInfo, NewOtherInfos("cpu_type", string(v.CPUSet[i].Type)))
		n.OtherInfo = append(n.OtherInfo, NewOtherInfos("cpu_cores", string(v.CPUSet[i].Num)))
		*nodes = append(*nodes, n)
	}
}

func (v *VMCData) parseGPU(nodes *pNodesArr) {
	for i := 0; i < int(v.GPUNumber); i++ {
		n := &Nodes{
			Id:           int64(v.GPUSet[i].ID),
			Name:         v.GPUSet[i].Name,
			DeviceType:   "gpu",
			ParentId:     uint16(v.VMCID),
			UpstreamId:   0,
			DeviceStatus: "RUN",
			OtherInfo:    make([]*OtherInfos, 0),
		}
		n.OtherInfo = append(n.OtherInfo, NewOtherInfos("gpu_type", string(v.GPUSet[i].Type)))
		n.OtherInfo = append(n.OtherInfo, NewOtherInfos("gpu_cores", string(v.GPUSet[i].Num)))
		*nodes = append(*nodes, n)
	}
}

func (v *VMCData) parseDSP(nodes *pNodesArr) {
	for i := 0; i < int(v.DSPNumber); i++ {
		n := &Nodes{
			Id:           int64(v.DSPSet[i].ID),
			Name:         v.DSPSet[i].Name,
			DeviceType:   "dsp",
			ParentId:     uint16(v.VMCID),
			UpstreamId:   0,
			DeviceStatus: "RUN",
			OtherInfo:    make([]*OtherInfos, 0),
		}
		n.OtherInfo = append(n.OtherInfo, NewOtherInfos("dsp_type", string(v.DSPSet[i].Type)))
		n.OtherInfo = append(n.OtherInfo, NewOtherInfos("dsp_cores", string(v.DSPSet[i].Num)))
		*nodes = append(*nodes, n)
	}
}

func (v *VMCData) parseFPGA(nodes *pNodesArr) {
	for i := 0; i < int(v.FPGANumber); i++ {
		n := &Nodes{
			Id:           int64(v.FPGASet[i].ID),
			Name:         v.FPGASet[i].Name,
			DeviceType:   "fpga",
			ParentId:     uint16(v.VMCID),
			UpstreamId:   0,
			DeviceStatus: "RUN",
			OtherInfo:    make([]*OtherInfos, 0),
		}
		n.OtherInfo = append(n.OtherInfo, NewOtherInfos("fpga_type", string(v.FPGASet[i].Type)))
		n.OtherInfo = append(n.OtherInfo, NewOtherInfos("fpga_cores", string(v.FPGASet[i].Num)))
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
		ParentId:     uint16(v.SwitchID),
		UpstreamId:   0,
		DeviceStatus: "RUN",
		OtherInfo:    make([]*OtherInfos, 0),
	}
	n.OtherInfo = append(n.OtherInfo, NewOtherInfos("proto_type", string(v.protoType)))
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
		(n.OtherInfo) = append(n.OtherInfo, NewOtherInfos("switch_type", string(v.SwitchDeviceSet[i].SwitchType)))
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
		n.OtherInfo = append(n.OtherInfo, NewOtherInfos("rtu_type", string(v.RemoteUnitSet[i].RemoteUnitType)))
		*nodes = append(*nodes, n)
	}
}

func NewNodes(v *VMCData, isFirst bool) pNodesArr {
	nodes := make(pNodesArr, 0)
	v.parseVMC(&nodes)
	if !isFirst {
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
		Id:           cTopoID,
		Node:         NewNodes(v, isFirst),
		TransferInfo: NewTransferInfos(v),
	}
}

func (t *TopoTable) CreateOp(v *VMCData) error {
	err := mgm.CollectionByName(config.CommonConfig.DBTopoTableName).First(bson.M{"id": cTopoID}, t)
	if err != nil {
		glog.Error("[CreateOp] Find error  err:%+v", err)
	}
	glog.Info("[CreateOp] find return: %+v", t)
	if t == nil {
		t = NewTopoTable(v, true)
		return mgm.CollectionByName(config.CommonConfig.DBTopoTableName).Create(t)
	}
	t = NewTopoTable(v, false)
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
		glog.Error("[InsertOp] Find error err:%+v", err)
	}
	t.Node = append(t.Node, node)
	return t.UpdateOp()
}

func (t *TopoTable) DeleteOp(id int64) error {
	err := t.CollectOp()
	if err != nil {
		glog.Error("[DeleteOp] Find error err:%+v", err)
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
