package basicconductor

import "context"

type Option struct {
	WatchContext		context.Context
	WatchChannelSize 	int
}