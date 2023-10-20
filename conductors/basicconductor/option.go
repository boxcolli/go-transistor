package basicconductor

import (
	"context"

	"github.com/boxcolli/go-transistor/types"
)

type Option struct {
	GetWatchContext		func() context.Context
	WatchChannelSize 	int

	GetDialContext		func() context.Context

	DefaultChange		*types.Change
}
