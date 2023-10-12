package io

import "github.com/boxcolli/go-transistor/types"

type StreamWriter interface {
	Write(*types.Message) error
}
