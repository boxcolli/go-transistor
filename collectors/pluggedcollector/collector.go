package pluggedcollector

import (
	"github.com/boxcolli/pepperlink/collectors"
	"github.com/boxcolli/pepperlink/plugs"
)

type pluggedCollector struct {
	cname string
	plug  plugs.Plug
}

// Load implements Collector.
func (*pluggedCollector) Load(stream collectors.StreamAdapter) {
	panic("unimplemented")
}

func NewPluggedCollector(cname string, plug plugs.Plug) collectors.Collector {
	return &pluggedCollector{
		cname: cname,
		plug:  plug,
	}
}
