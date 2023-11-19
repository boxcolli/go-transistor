package io

import (
	"errors"

	"github.com/boxcolli/go-transistor/types"
)

var (
	ErrClosed = errors.New("connection closed")
)

type StreamReader interface {
	Read() (*types.Message, error)
}

type StreamWriter interface {
	Write(*types.Message) error
}

type Bus interface {
	Push(m *types.Message) // should be goroutine safe
	Pull() <-chan *types.Message
	Lock()		// turn off the message supply
	Unlock()	// turn on the message supply
}
