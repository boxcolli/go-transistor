package types

type Operation byte

const (
	OperationUnspecified Operation = iota
    OperationAdd
    OperationDel
)
