package db

import "star_atlas_server/model"

type OBCModule struct {
	OBCID      uint8
	CPUNumber  uint8
	DSPNumber  uint8
	GPUNumber  uint8
	FPAGNumber uint8
}

type VMCModule struct {
	VMCName   string
	VMCID     uint8
	OBCModule []*OBCModule
	Next      uint8
}

type SwitchModule struct {
	Head     uint8
	SwitchID uint8
	Next     uint8
}

type RTUModule struct {
	Type uint8
	Head uint8
}

// see https://github.com/Kamva/mgm
type TopoTable struct {
	TopoName     string `json:"topo_name" bson:"topo_name"`
	VMCModule    []*VMCModule
	SwitchModule []*SwitchModule
	RTUModule    []*RTUModule
}

func WriteTopoTable(v *model.VMCData) error {

	return nil
}
