package basicemitter

import (
	"sync"

	"github.com/boxcolli/go-transistor/emitter"
	"github.com/boxcolli/go-transistor/io"
	"github.com/boxcolli/go-transistor/types"
)

type basicEmitter struct {
	m    	chan *types.Message
	mmx		sync.RWMutex
	stop	chan bool
	wg		sync.WaitGroup
	wgmx	sync.Mutex
}

func NewBasicEmitter(qsiz int) emitter.Emitter {
	return &basicEmitter{
		m:    	make(chan *types.Message, qsiz),
		mmx:	sync.RWMutex{},
		wg: 	sync.WaitGroup{},
		stop:	make(chan bool),
		wgmx:	sync.Mutex{},
	}
}

// Work implements emitter.Emitter.
func (e *basicEmitter) Work(w io.StreamWriter) error {
	{
		e.wgmx.Lock()
		e.wg.Add(1)
		e.wgmx.Unlock()
		defer e.wg.Done()
	}

	for {
		select {
		case <- e.stop:

			// Drain the channel
			e.mmx.Lock()
			close(e.m)
			for m := range e.m {
				err := w.Write(m)
				if err != nil { break }
			}
			e.m = make(chan *types.Message)
			e.mmx.Unlock()

			return nil

		case m, ok := <-e.m:

			if !ok { return emitter.ErrClosed }
			err := w.Write(m)
			if err != nil { return err }

		}
	}
}

// Stop implements emitter.Emitter.
func (e *basicEmitter) Stop() {
	e.wgmx.Lock()
	defer e.wgmx.Unlock()

	close(e.stop)				// send stop signal
	e.wg.Wait()					// wait for workers
	e.stop = make(chan bool)	// replace stop with new channel
}

// Emit implements emitter.Emitter.
func (e *basicEmitter) Emit(m *types.Message) {
	e.mmx.RLock()
	e.m <- m
	e.mmx.RUnlock()
}
