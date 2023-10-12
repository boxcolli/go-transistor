package basic

import (
	"github.com/boxcolli/go-transistor/collector"
	"github.com/boxcolli/go-transistor/io"
)

type basicCollector struct {
	
}

// Work implements collector.Collector.
func (*basicCollector) Work(r *io.StreamReader, call func(e error)) {
	panic("unimplemented")
}

// Stop implements collector.Collector.
func (*basicCollector) Stop(r *io.StreamReader) {
	panic("unimplemented")
}

// StopAll implements collector.Collector.
func (*basicCollector) StopAll() {
	panic("unimplemented")
}

func NewBasicCollector() collector.Collector {
	return &basicCollector{}
}
