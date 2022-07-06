package model

import (
	"star_atlas_server/config"
	"time"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TransferInfos struct {
	FromId      uint16        `json:"from_id" bson:"from_id"`
	toId        uint16        `json:"to_id" bson:"to_id"`
	TaskType    uint8         `json:"task_type" bson:"task_type"`
	StartTime   time.Time     `json:"start_time" bson:"start_time"`
	EndTime     time.Time     `json:"end_time" bson:"end_time"`
	durningTime time.Duration `json:"durning_time" bson:"durning_time"`
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
	TopoId       primitive.ObjectID `json:"topo_id,omitempty" bson:"topo_id,omitempty"`
	Node         []*Nodes           `json:"node" bson:"node"`
	TransferInfo []*TransferInfos   `json:"transfer_info" bson:"transfer_info"`
}

type Devices interface {
	GetDeviceNums()
}

func getDeviceNums(d Devices) uint {
	return d.
}

func NewTransferInfos(v *VMCData) []*TransferInfos {
	obcNum := 3
	obcModule := make([]*TransferInfos, 0)
	for i := 0; i < obcNum; i++ {
		o := &OBCModule{uint8(i), v.CPUNumber, v.DSPNumber, v.GPUNumber, v.FPGANumber}
		obcModule = append(obcModule, o)
	}
	return obcModule
}

func otherInfosHandler(v *VMCData) []*OtherInfos {
	vmcNum := 2
	otherInfos := make([]*OtherInfos, 0)
	for i := 0; i < vmcNum; i++ {
		o := &OtherInfos{
			"proto_type", string(v.protoType),
		}
		otherInfos = append(otherInfos, o)
	}
	return otherInfos
}

func switchModuleHandler(v *VMCData) []*SwitchModule {
	switchNum := 3
	switchModule := make([]*SwitchModule, 0)
	for i := 0; i < switchNum; i++ {
		s := &SwitchModule{uint8(i), v.SwitchID, uint8(i)}
		switchModule = append(switchModule, s)
	}
	return switchModule
}

func rtuModuleHandler(v *VMCData) []*RTUModule {
	rtuNum := 3
	rtuModule := make([]*RTUModule, 0)
	for i := 0; i < rtuNum; i++ {
		r := &RTUModule{Type: uint8(i), Head: uint8(i)}
		rtuModule = append(rtuModule, r)
	}
	return rtuModule
}

func NewTopoTable(v *VMCData) *TopoTable {
	return &TopoTable{
		VMCModule:    vmcModuleHandler(v),
		SwitchModule: switchModuleHandler(v),
		RTUModule:    rtuModuleHandler(v),
	}
}

func (t *TopoTable) CreateOp(v *VMCData) error {
	t = NewTopoTable(v)
	return mgm.CollectionByName(config.CommonConfig.DBTopoTableName).Create(t)
}

func (t *TopoTable) CollectOp(v *VMCData) error {
	t = NewTopoTable(v)
	return mgm.CollectionByName(config.CommonConfig.DBTopoTableName).First(bson.M{}, t)
}
