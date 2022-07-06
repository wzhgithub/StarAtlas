package model

import (
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
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
	TopoId       primitive.ObjectID `json:"topo_id,omitempty" bson:"topo_id,omitempty"`
	Node         []*Nodes           `json:"node" bson:"node"`
	TransferInfo []*TransferInfos   `json:"transfer_info" bson:"transfer_info"`
}

func NewOtherInfos(key string, val string) *OtherInfos {
	return &OtherInfos{key, val}
}

func (v *VMCData) parseCPU(nodes []*Nodes) {
	for i := 0; i < int(v.CPUNumber); i++ {
		n := &Nodes{
			uint16(i), "cpu" + strconv.Itoa(i),
			"cpu", uint16(v.VMCID), 0, "RUN", nil,
		}
		NewOtherInfos("cpu_type", string(v.CPUSet[i].Type))
		nodes = append(nodes, n)
	}
}

func (v *VMCData) parseGPU(nodes []*Nodes) {
	for i := 0; i < int(v.GPUNumber); i++ {
		n := &Nodes{
			uint16(i), "gpu" + strconv.Itoa(i),
			"gpu", uint16(v.VMCID), 0, "RUN", nil,
		}
		// NewOtherInfos("cpu_type", string(v.CPUSet[i].Type))
		nodes = append(nodes, n)
	}
}

func (v *VMCData) parseObc(nodes []*Nodes) {
	v.parseCPU(nodes)
	v.parseGPU(nodes)
}

func (v *VMCData) parseVmc(nodes []*Nodes) {

}

func (v *VMCData) parseSwitch(nodes []*Nodes) {

}

func (v *VMCData) parseRtu(nodes []*Nodes) {

}

func NewNodes(v *VMCData) []*Nodes {
	nodes := make([]*Nodes, 0)
	v.parseObc(nodes)
	v.parseVmc(nodes)
	v.parseSwitch(nodes)
	v.parseRtu(nodes)
	return nodes
}

func NewTopoTable(v *VMCData) *TopoTable {
	return &TopoTable{
		Node: NewNodes(v),
	}
}

// func (t *TopoTable) CreateOp(v *VMCData) error {
// 	t = NewTopoTable(v)
// 	return mgm.CollectionByName(config.CommonConfig.DBTopoTableName).Create(t)
// }

// func (t *TopoTable) CollectOp(v *VMCData) error {
// 	t = NewTopoTable(v)
// 	return mgm.CollectionByName(config.CommonConfig.DBTopoTableName).First(bson.M{}, t)
// }
