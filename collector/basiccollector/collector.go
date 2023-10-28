package basiccollector

import (
	"sync"

	"github.com/boxcolli/go-transistor/base"
	"github.com/boxcolli/go-transistor/collector"
	"github.com/boxcolli/go-transistor/io"
)

type basicCollector struct {
	b base.Base

	stop map[io.StreamReader]chan bool
	mx   sync.Mutex
}

func NewBasicCollector(b base.Base) collector.Collector {
	return &basicCollector{
		b: b,
		stop: make(map[io.StreamReader]chan bool),
		mx: sync.Mutex{},
	}
}

// Work implements collector.Collector.
func (c *basicCollector) Work(r io.StreamReader) error {
	// Open a stop signal channel
	stopch := make(chan bool)
	{
		c.mx.Lock()
		c.stop[r] = stopch
		c.mx.Unlock()
	}

	// Work
	for {
		if <-stopch {
			return nil
		}

		// Read message
		m, err := r.Read()
		if err != nil {
			return err
		}

		// Send message
		err = c.b.Flow(m)
		if err != nil {
			return err
		}
	}
}

// Stop implements collector.Collector.
func (c *basicCollector) Stop(r io.StreamReader) {
	c.mx.Lock()
	defer c.mx.Unlock()

	close(c.stop[r])
	delete(c.stop, r)
}

// StopAll implements collector.Collector.
func (c *basicCollector) StopAll() {
	//delete all
	c.mx.Lock()
	defer c.mx.Unlock()

	for k := range c.stop {
		close(c.stop[k])
		delete(c.stop, k)
	}
}
