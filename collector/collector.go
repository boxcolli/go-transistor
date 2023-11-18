package collector

import (
	"github.com/boxcolli/go-transistor/base"
	"github.com/boxcolli/go-transistor/io"
)

type Collector interface {
	Work(b base.Base, r io.StreamReader) error
	Stop(r io.StreamReader)
	StopAll()
}
