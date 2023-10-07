package types

type Method int

const (
	MethodUnspecified Method = iota
    MethodEmpty
    MethodCreate
    MethodUpdate
    MethodDelete
)
