package servicecollector

import "github.com/boxcolli/pepperlink/collectors"

type serviceCollector struct {
}

// Load implements collectors.Collector.
func (*serviceCollector) Load(stream collectors.StreamAdapter) {
	panic("unimplemented")
}

func NewServiceCollector() collectors.Collector {
	return &serviceCollector{}
}
