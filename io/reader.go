package io

import (
	"errors"

	"github.com/boxcolli/go-transistor/types"
)

var (
	ErrClosed = errors.New("connection closed.")
)

type StreamReader interface {
	Read() (*types.Message, error)
}
