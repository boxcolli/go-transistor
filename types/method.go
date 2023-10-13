package types

type Method byte

const (
	MethodUnspecified Method = iota
    MethodEmpty
    MethodCreate
    MethodUpdate
    MethodDelete
)
