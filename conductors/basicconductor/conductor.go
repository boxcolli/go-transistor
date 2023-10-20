package basicconductor

import (
	"sync"

	"github.com/boxcolli/go-transistor/collector"
	"github.com/boxcolli/go-transistor/conductors"
	"github.com/boxcolli/go-transistor/dialers"
	"github.com/boxcolli/go-transistor/io"
	"github.com/boxcolli/go-transistor/plugs"
	"github.com/boxcolli/go-transistor/types"
)

type entry struct {
	state state
	op    types.Operation
	sr    io.StreamReader
}

type basicConductor struct {
	// Dependency
	p plugs.Plug
	d dialers.Dialer
	c collector.Collector

	opt	Option

	// Cluster list
	cnames	map[string]bool
	cmx		sync.Mutex

	stop	map[string]chan bool
	stopmx	sync.Mutex

	// State holder
	e  	map[*types.Member]entry
	emx	sync.Mutex
}

// Start implements conductor.Conductor.
func (c *basicConductor) Begin(cname string) error {
	ech := make(chan error)
	go c.begin(cname, ech)

	select {
	case err := <- ech:
		close(ech)
		return err
	}
}

func (c *basicConductor) begin(cname string, ech chan<- error) {
	// Add the new cluster name
	{
		c.cmx.Lock()
		if c.cnames[cname] {
			c.cmx.Unlock()
			return
		}
		c.cnames[cname] = true
		c.cmx.Unlock()
	}

	// Add stop channel
	stopch := make(chan bool)
	{
		c.stopmx.Lock()
		c.stop[cname] = stopch
		c.stopmx.Unlock()
	}

	// Watch
	var watch <-chan *plugs.Event
	{
		var err error
		watch, err = c.p.Watch(c.opt.WatchContext, cname, c.opt.WatchChannelSize)
		if err != nil {
			// Something went wrong.
			c.cmx.Lock()
			c.cnames[cname] = true
			c.cmx.Unlock()
			ech <- err
		}
	}

	// Loop
	stop := false
	for !stop {
		select {
		case <- stopch:
			//
		case e, ok := <- watch:
			if !ok {

			}
		}
	}
}

// Stop implements conductor.Conductor.
func (c *basicConductor) End(cname string) {
	panic("unimplemented")
}

// StopAll implements conductor.Conductor.
func (c *basicConductor) EndAll() {
	panic("unimplemented")
}

func NewBasicConductor(p plugs.Plug, d dialers.Dialer, c collector.Collector, opt Option) conductors.Conductor {
	return &basicConductor{
		p: p,
		d: d,
		c: c,
		opt: opt,
		cnames: make(map[string]bool),
		cmx: sync.Mutex{},
		stop: make(map[string]chan bool),
		stopmx: sync.Mutex{},
		e: make(map[*types.Member]entry),
		emx: sync.Mutex{},
	}
}
