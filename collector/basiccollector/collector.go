package basiccollector

import (
	"sync"

	"github.com/boxcolli/go-transistor/base"
	"github.com/boxcolli/go-transistor/collector"
	"github.com/boxcolli/go-transistor/io"
	"github.com/boxcolli/go-transistor/types"
)

type entry struct {
	stop	chan bool
	m		chan *types.Message
	e		chan error
}
func newEntry(msiz int) *entry {
	return &entry{
		stop:	make(chan bool),
		m:		make(chan *types.Message),
		e:		make(chan error),
	}
}

type basicCollector struct {
	b	base.Base

	mqs	int
	ent	map[io.StreamReader]*entry
	mx  sync.Mutex
}

func NewBasicCollector(b base.Base, msgQueueSize int) collector.Collector {
	return &basicCollector{
		b:		b,
		mqs:	msgQueueSize,
		ent:	make(map[io.StreamReader]*entry),
		mx:		sync.Mutex{},
	}
}

// Work implements collector.Collector.
func (c *basicCollector) Work(r io.StreamReader) error {
	ent := newEntry(c.mqs)
	{
		c.mx.Lock()
		if _, ok := c.ent[r]; ok {
			return nil	// 1 thread per stream reader is allowed
		}
		c.ent[r] = ent
		c.mx.Unlock()
	}

	// Async read
	go func(ent *entry) {
		for {
			// Read message
			m, err := r.Read()
			if err != nil {
				ent.e <- err
				return
			}

			ent.m <- m
			
			// Check stop signal
			select {
			case <- ent.stop:
				return
			default:
			}
		}
	} (ent)

	// Work
	for {
		select {
		case <- ent.stop:
			// Stop signal received
			return nil

		case m := <- ent.m:
			// New message
			// fmt.Printf("collector received: %v\n", *m)
			c.b.Flow(m)

		case err := <- ent.e:
			// There is a problem with stream reader
			return err
		}
	}
}

// Stop implements collector.Collector.
func (c *basicCollector) Stop(r io.StreamReader) {
	c.mx.Lock()
	defer c.mx.Unlock()

	close(c.ent[r].stop)
	delete(c.ent, r)
}

// StopAll implements collector.Collector.
func (c *basicCollector) StopAll() {
	//delete all
	c.mx.Lock()
	defer c.mx.Unlock()

	for _, ent := range c.ent {
		close(ent.stop)
	}

	// wipe out
	c.ent = make(map[io.StreamReader]*entry)
}
