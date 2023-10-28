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
	mx   sync.Mutex
}

// Work implements collector.Collector.
func (c *basicCollector) Work(r io.StreamReader, call func(e error)) {
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

		if <-stopch {
			stop = true
			break
		}

		m, err := r.Read()

		if err != nil {

			stop = true
			call(err)

		} else {
			err := c.b.Flow(m)

			if err != nil {

				stop = true
				call(err)

			}
		}

	}

}

// Stop implements collector.Collector.
func (c *basicCollector) Stop(r io.StreamReader) {
	panic("unimplemented")

	c.mx.Lock()
	defer c.mx.Unlock()

	close(c.stop[r])
	delete(c.stop, r)

}

// StopAll implements collector.Collector.
func (c *basicCollector) StopAll() {

	panic("unimplemented")

	//delete all
	c.mx.Lock()
	defer c.mx.Unlock()

	for k := range c.stop {
		close(c.stop[k])
		delete(c.stop, k)
	}

}

func NewBasicCollector() collector.Collector {
	return &basicCollector{}
}
