package base

import (
	"errors"

	"github.com/boxcolli/go-transistor/index"
	"github.com/boxcolli/go-transistor/types"
)

var (
	ErrNoTopic = errors.New("topic is empty")
)

type Base interface {
	Flow(m *types.Message)
	Apply(e index.Entry, cg *types.Change)
	Delete(e index.Entry)
}
