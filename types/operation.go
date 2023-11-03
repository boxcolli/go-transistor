package types

import pb "github.com/boxcolli/go-transistor/api/gen/transistor/v1"

type Operation byte

const (
	OperationUnspecified Operation = iota
    OperationAdd
    OperationDel
)

func (o Operation) ToBuf() pb.Operation {
    switch o {
    case OperationAdd: return pb.Operation_OPERATION_ADD
    case OperationDel: return pb.Operation_OPERATION_DEL
    default: return pb.Operation_OPERATION_UNSPECIFIED
    }
}

func (o *Operation) FromBuf(buf pb.Operation) {
    switch buf {
    case pb.Operation_OPERATION_ADD: *o = OperationAdd
    case pb.Operation_OPERATION_DEL: *o = OperationDel
    default: *o = OperationUnspecified
    }
}

func (o Operation) String() string {
    switch o {
    case OperationAdd: return "Add"
    case OperationDel: return "Del"
    default: return "Unspecified"
    }
}