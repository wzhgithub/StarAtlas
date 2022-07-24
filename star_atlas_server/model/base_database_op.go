package model

// Mongodb base operation: CRUD
type IDatabaseOperation interface {
	CreateOp(v *VMCData) error
	CollectOp() error
	InsertOp() error
	UpdateOp() error
	DeleteOp(delIds []int64) error
	CreateData() error
	CollectData() error
}
