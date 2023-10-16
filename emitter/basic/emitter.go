package basicemitter

import (
	"github.com/boxcolli/go-transistor/emitter"
	"github.com/boxcolli/go-transistor/io"
	"github.com/boxcolli/go-transistor/types"
)

type basicEmitter struct {
	q    chan *types.Message
	stop chan bool
}

// Work implements emitter.Emitter.
func (e *basicEmitter) Work(w io.StreamWriter) error {
	stop := false
	for !stop {
		select {
		case <-e.stop:
			// Stop working
			stop = true

		default:
			// Work
			m, ok := <-e.q
			if !ok {
				return emitter.ErrClosed
			}
			err := w.Write(m)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// Stop implements emitter.Emitter.
func (e *basicEmitter) Stop() {
	e.stop <- true
}

// Emit implements emitter.Emitter.
func (e *basicEmitter) Emit(m *types.Message) {
	e.q <- m
}

func NewBasicEmitter(qsiz int) emitter.Emitter {
	return &basicEmitter{
		q:    make(chan *types.Message, qsiz),
		stop: make(chan bool),
	}
}
