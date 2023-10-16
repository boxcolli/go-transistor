package basic

import (
	"sync"

	"github.com/boxcolli/go-transistor/base"
	"github.com/boxcolli/go-transistor/collector"
	"github.com/boxcolli/go-transistor/io"
)

type basicCollector struct {
	b base.Base

	stop map[io.StreamReader]chan bool
	mx sync.Mutex
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
	stop := false
	for !stop {
		select {
		case <-stopch:
			// Stop signal received.
			stop = true
			
		default:
			// Read message
			m, err := r.Read()
			if err != nil {
				// Stream has an error.
				stop = true
				return err
			} else {
				// Send message
				err := c.b.Flow(m)
				if err != nil {
					// Base has an error
					stop = true
					return err
				}
			}
		}
	}

	return nil
}

// Stop implements collector.Collector.
func (c *basicCollector) Stop(r io.StreamReader) {
	panic("unimplemented")
	delete(c.stop, r)
}

// StopAll implements collector.Collector.
func (c *basicCollector) StopAll() {
	panic("unimplemented")
}

func NewBasicCollector() collector.Collector {
	return &basicCollector{}
}
