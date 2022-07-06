package model

import (
	"star_atlas_server/config"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
)

type OBCModule struct {
	OBCID      uint8 `json:"obc_id" bson:"obc_id"`
	CPUNumber  uint8 `json:"cpu_num" bson:"cpu_num"`
	DSPNumber  uint8 `json:"dsp_num" bson:"dsp_num"`
	GPUNumber  uint8 `json:"gpu_num" bson:"gpu_num"`
	FPAGNumber uint8 `json:"fpag_num" bson:"fpag_num"`
}

type VMCModule struct {
	VMCName   string       `json:"vmc_name" bson:"vmc_name"`
	VMCID     uint8        `json:"vmc_id" bson:"vmc_id"`
	OBCModule []*OBCModule `json:"obc_module" bson:"obc_module"`
	Next      uint8        `json:"next" bson:"next"`
}

type SwitchModule struct {
	Head     uint8 `json:"head" bson:"head"`
	SwitchID uint8 `json:"switch_id" bson:"switch_id"`
	Next     uint8 `json:"next" bson:"next"`
}

type RTUModule struct {
	Type uint8 `json:"type" bson:"type"`
	Head uint8 `json:"head" bson:"head"`
}

// see https://github.com/Kamva/mgm
type TopoTable struct {
	mgm.DefaultModel `bson:",inline"`
	VMCModule        []*VMCModule    `json:"vmc_module" bson:"vmc_module"`
	SwitchModule     []*SwitchModule `json:"switch_module" bson:"switch_module"`
	RTUModule        []*RTUModule    `json:"rtu_module" bson:"rtu_module"`
}

func obcModuleHandler(v *VMCData) []*OBCModule {
	obcNum := 3
	obcModule := make([]*OBCModule, 0)
	for i := 0; i < obcNum; i++ {
		o := &OBCModule{uint8(i), v.CPUNumber, v.DSPNumber, v.GPUNumber, v.FPAGNumber}
		obcModule = append(obcModule, o)
	}
	return obcModule
}

func vmcModuleHandler(v *VMCData) []*VMCModule {
	vmcNum := 2
	vmcModule := make([]*VMCModule, 0)
	for i := 0; i < vmcNum; i++ {
		v := &VMCModule{v.VMCName, v.VMCID, obcModuleHandler(v), uint8(i)}
		vmcModule = append(vmcModule, v)
	}
	return vmcModule
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
