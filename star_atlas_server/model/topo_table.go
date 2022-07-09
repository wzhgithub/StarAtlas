package model

import (
	"star_atlas_server/config"
	"time"

	"github.com/golang/glog"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
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
	Id           uint16        `json:"id" bson:"id"`
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
			Id:           uint16(v.CPUSet[i].ID),
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
			Id:           uint16(v.GPUSet[i].ID),
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
			Id:           uint16(v.DSPSet[i].ID),
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
			Id:           uint16(v.FPGASet[i].ID),
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
		Id:           uint16(v.VMCID),
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
			Id:           uint16(v.SwitchDeviceSet[i].SwitchOrder),
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
			Id:           uint16(v.RemoteUnitSet[i].RemoteUnitOrder),
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
		Node:         NewNodes(v),
		TransferInfo: NewTransferInfos(v),
	}
}

func (t *TopoTable) CreateOp(v *VMCData) error {
	t = NewTopoTable(v)
	return mgm.CollectionByName(config.CommonConfig.DBTopoTableName).Create(t)
}

func (t *TopoTable) CollectOp() error {
	return mgm.CollectionByName(config.CommonConfig.DBTopoTableName).First(bson.M{}, t)
}

func (t *TopoTable) InsertOp(node *Nodes) error {
	err := mgm.CollectionByName(config.CommonConfig.DBTopoTableName).First(bson.M{}, t)
	if err != nil {
		glog.Error("[InsertOp] Find error")
	}
	t.Node = append(t.Node, node)
	return mgm.CollectionByName(config.CommonConfig.DBTopoTableName).Create(t)
}

func (t *TopoTable) DeleteOp(id uint16) error {
	err := mgm.CollectionByName(config.CommonConfig.DBTopoTableName).First(bson.M{}, t)
	if err != nil {
		glog.Error("[DeleteOp] Find error")
	}
	var index int
	for idx, node := range t.Node {
		if node.Id == id {
			index = idx
			break
		}
	}
	t.Node = append(t.Node[:index], t.Node[index+1:]...)
	return mgm.CollectionByName(config.CommonConfig.DBTopoTableName).Create(t)
}
