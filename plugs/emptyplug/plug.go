package emptyplug

import (
	"github.com/boxcolli/pepperlink/plugs"
)

type emptyPlug struct {
}

// GetDiscoveryAddr implements plugs.Plug.
func (*emptyPlug) GetDiscoveryAddr() {
	panic("unimplemented")
}

// Destroy implements plugs.Plug.
func (*emptyPlug) Destroy() {
	panic("unimplemented")
}

// WatchPub implements plugs.Plug.
func (*emptyPlug) WatchPub() {
	panic("unimplemented")
}

func NewEmptyCluster() plugs.Plug {
	return &emptyPlug{}
}
