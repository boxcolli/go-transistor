package servicecollector

import "github.com/boxcolli/go-transistor/collectors"

type serviceCollector struct {
}

// Load implements collectors.Collector.
func (*serviceCollector) Load(stream collectors.StreamAdapter) {
	panic("unimplemented")
}

func NewServiceCollector() collectors.Collector {
	return &serviceCollector{}
}
