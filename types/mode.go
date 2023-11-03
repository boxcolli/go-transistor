package types

import pb "github.com/boxcolli/go-transistor/api/gen/transistor/v1"

type Mode byte

const (
	ModeUnspecified Mode = iota
	ModeAny
	ModeRoute
	ModeRoot
)

func (m Mode) ToBuf() pb.Mode {
	switch m {
	case ModeAny: return pb.Mode_MODE_ANY
	case ModeRoute: return pb.Mode_MODE_ROUTE
	case ModeRoot: return pb.Mode_MODE_ROOT
	default: return pb.Mode_MODE_UNSPECIFIED
	}
}

func (m *Mode) FromBuf(buf pb.Mode) {
	switch buf {
	case pb.Mode_MODE_ANY: *m = ModeAny
	case pb.Mode_MODE_ROUTE: *m = ModeRoute
	case pb.Mode_MODE_ROOT: *m = ModeRoot
	default: *m = ModeUnspecified
	}
}

func (m Mode) String() string {
	switch m {
	case ModeAny: return "Any"
	case ModeRoute: return "Route"
	case ModeRoot: return "Root"
	default: return "Unspecified"
	}
}