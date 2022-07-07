package model

// Mongodb base operation: CRUD
type IDatabaseOperation interface {
	CreateOp(v *VMCData) error
	CollectOp() error
	DeleteOp(id uint16) error
	CreateData() error
	CollectData() error
}
