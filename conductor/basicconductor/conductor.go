package basicconductor

import (
	"sync"

	"github.com/boxcolli/go-transistor/collector"
	"github.com/boxcolli/go-transistor/conductor"
	"github.com/boxcolli/go-transistor/dialer"
	"github.com/boxcolli/go-transistor/io"
	"github.com/boxcolli/go-transistor/plug"
	"github.com/boxcolli/go-transistor/types"
)

type clusterEntry struct {
	stop	chan bool
	cg		chan *types.Change

	m map[*types.Member]*memberEntry
	mmx sync.Mutex
}

// Member state holder
type memberEntry struct {
	stop	chan bool
	cg		chan *types.Change
	st		state
	stmx	sync.Mutex
	op   	types.Operation
	opmx	sync.Mutex
	sr   	io.StreamReader
	srmx	sync.Mutex
}

type basicConductor struct {
	// Class dependency
	p plug.Plug
	d dialer.Dialer
	c collector.Collector

	// Myself
	me *types.Member

	opt	Option

	// Per cluster entries
	ce map[string]*clusterEntry
	cemx sync.RWMutex
}

func NewBasicConductor(p plug.Plug, d dialer.Dialer, c collector.Collector, opt Option) conductor.Conductor {
	return &basicConductor{
		p: p,
		d: d,
		c: c,
		opt: opt,
		ce: make(map[string]*clusterEntry),
		cemx: sync.RWMutex{},
	}
}

// Start implements conductor.Conductor.
func (c *basicConductor) Begin(cname string) error {
	c.cemx.Lock()
	defer c.cemx.Unlock()

	// Check if there's already a worker running
	if _, ok := c.ce[cname]; ok {
		return nil
	}

	// Create a new cluster entry
	var ce *clusterEntry
	{
		c.ce[cname] = &clusterEntry{
			stop: make(chan bool),
			cg: make(chan *types.Change),
			m: make(map[*types.Member]*memberEntry),
			mmx: sync.Mutex{},
		}
	}

	// Start a worker
	ech := make(chan error)	// state signal
	go c.clusterWorker(cname, ce, ech)
	err := <- ech
	close(ech)
	return err
}

func (c *basicConductor) clusterWorker(cname string, ce *clusterEntry, ech chan<- error) {
	// Watch
	var watch <-chan *plug.Event
	{
		var err error
		watch, err = c.p.Watch(c.opt.GetWatchContext(), cname, c.opt.WatchChannelSize)
		if err != nil {
			// Something went wrong.
			ech <- err
			return
		}
	}

	ech <- nil

	// Worker
	for {
		select {
		case <- ce.stop:
			panic("unimplemented")

		case <- ce.cg:
			panic("unimplemented")

		case event, ok := <- watch:
			if !ok {
				// Something went wrong.
				return
			}

			m := event.Data

			switch event.Op {
			case types.OperationAdd:
				ce.mmx.Lock()
				if _, ok := ce.m[m]; ok {
					// Member worker is already running
				} else {
					// Add a new member entry
					ent := &memberEntry{
						stop: make(chan bool),
						cg: make(chan *types.Change),
						st: stateInit,
						stmx: sync.Mutex{},
						op: types.OperationAdd,
						opmx: sync.Mutex{},
						sr: nil,
						srmx: sync.Mutex{},
					}
					ce.m[m] = ent

					// Run a member worker
					go c.memberWorker(ce, m, ent)
				}
				ce.mmx.Unlock()

			case types.OperationDel:
				
			}
		}
	}
}

func (c *basicConductor) memberWorker(ce *clusterEntry, m *types.Member, ent *memberEntry) {
	// Dial
	{
		ent.stmx.Lock()
		ent.st = stateConn
		ent.stmx.Unlock()

		var err error
		ent.sr, err = c.d.Dial(c.opt.GetDialContext(), m, c.opt.DefaultChange)
		if err != nil {
			// Something went wrong.
			return
		}
	}

	// Check Op
	{
		ent.opmx.Lock()
		if ent.op == types.OperationDel {
			// Finish process
			return
		}
	}

	

}

func (c *basicConductor) Apply(cg *types.Change) {

}

// Stop implements conductor.Conductor.
func (c *basicConductor) End(cname string) {
	panic("unimplemented")
}

// StopAll implements conductor.Conductor.
func (c *basicConductor) EndAll() {
	panic("unimplemented")
}
