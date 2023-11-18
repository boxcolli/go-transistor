package basicemitter

import (
	"sync"

	"github.com/boxcolli/go-transistor/emitter"
	"github.com/boxcolli/go-transistor/io"
	"github.com/boxcolli/go-transistor/io/bus/channelbus"
)

type entry struct {
	stop chan bool
	bus  io.Bus
	wg   sync.WaitGroup
}

func newEntry(qs int) *entry {
	return &entry{
		stop: make(chan bool),
		bus:  channelbus.NewChannelBus(qs),
		wg:   sync.WaitGroup{},
	}
}

type basicEmitter struct {
	mqs int
	ent map[io.StreamWriter]*entry
	mx  sync.RWMutex
}

// Enter implements emitter.Emitter.
func (e *basicEmitter) Bus(w io.StreamWriter) (io.Bus, bool) {
	e.mx.Lock()
	defer e.mx.Unlock()
	
	if ent, ok := e.ent[w]; ok {

		// entry already exists
		return ent.bus, false

	} else {

		// create new entry
		ent = newEntry(e.mqs)
		e.ent[w] = ent
		return ent.bus, true

	}
}

func NewBasicEmitter(mqs int) emitter.Emitter {
	return &basicEmitter{
		mqs: mqs,
		ent: make(map[io.StreamWriter]*entry),
		mx:  sync.RWMutex{},
	}
}

func (e *basicEmitter) lookupEntry(w io.StreamWriter) *entry {
	e.mx.RLock()
	defer e.mx.RUnlock()

	if ent, ok := e.ent[w]; !ok {
		return nil
	} else {
		ent.wg.Add(1)	// increase the wait group here because the Mutex is locked
		return ent
	}
}

func (e *basicEmitter) Work(w io.StreamWriter) error {
	ent := e.lookupEntry(w)
	if ent == nil { return nil }
	defer ent.wg.Done()

	pull := ent.bus.Pull()
	for {
		select {
		case <-ent.stop:

			// Drain the bus
			ent.bus.Lock()
			defer ent.bus.Unlock()
			for m := range pull {
				err := w.Write(m)
				if err != nil { break }
			}

			return nil

		case m, ok := <- pull:

			// Write message
			if !ok { return emitter.ErrClosed }
			err := w.Write(m)
			if err != nil { return err }

		}
	}
}

func (e *basicEmitter) Stop(w io.StreamWriter) {
	e.mx.Lock()
	defer e.mx.Unlock()

	ent := e.ent[w]
	close(ent.stop)  // send stop signal
	ent.wg.Wait()    // wait for workers
	delete(e.ent, w) // delete entry
}

func (e *basicEmitter) StopAll() {
	e.mx.Lock()
	defer e.mx.Unlock()

	for _, ent := range e.ent {
		close(ent.stop)
	}

	for _, ent := range e.ent {
		ent.wg.Wait()
	}

	e.ent = make(map[io.StreamWriter]*entry)	// wipe out entries
}
