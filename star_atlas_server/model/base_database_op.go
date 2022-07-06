package model

// Mongodb base operation: CRUD
type IDatabaseOperation interface {
	CreateOp(v *VMCData) error
	CollectOp(v *VMCData) error
}
