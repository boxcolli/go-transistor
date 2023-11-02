package types

import pb "github.com/boxcolli/go-transistor/api/gen/transistor/v1"

type Method byte

const (
	MethodUnspecified Method = iota
	MethodEmpty
	MethodCreate
	MethodUpdate
	MethodDelete
)

func (m Method) ToBuf() pb.Method {
	switch m {
	case MethodEmpty:
		return pb.Method_METHOD_EMPTY
	case MethodCreate:
		return pb.Method_METHOD_CREATE
	case MethodUpdate:
		return pb.Method_METHOD_UPDATE
	case MethodDelete:
		return pb.Method_METHOD_DELETE
	default:
		return pb.Method_METHOD_UNSPECIFIED
	}
}

func (m Method) ToPb(p pb.Method) {
	switch p {
	case pb.Method_METHOD_EMPTY:
		m = MethodEmpty
		//return MethodEmpty
	case pb.Method_METHOD_CREATE:
		m = MethodCreate
	case pb.Method_METHOD_UPDATE:
		m = MethodUpdate
	case pb.Method_METHOD_DELETE:
		m = MethodDelete
	default:
		m = MethodUnspecified
	}
}
