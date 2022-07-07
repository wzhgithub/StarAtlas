package model

// Mongodb base operation: CRUD
type IDatabaseOperation interface {
	CreateOp(v *VMCData) error
	CollectOp() error
	CreateData() error
	CollectData() error
}
