package basicbase

import (
	"sync"

	"github.com/boxcolli/go-transistor/base"
	"github.com/boxcolli/go-transistor/emitter"
	"github.com/boxcolli/go-transistor/index"
	"github.com/boxcolli/go-transistor/types"
)

type task struct {
	e  emitter.Emitter
	cg *types.Change
}

type basicBase struct {
	i		index.Index
	icopy	index.Index
	imx 	sync.RWMutex

	started	bool
	stop	chan bool
	tch 	chan *task
	tq		[]*task
	wg		sync.WaitGroup
	mx		sync.Mutex
}

func NewBasicBase(buildIndex func() index.Index, msgQueueSize int) base.Base {
	return &basicBase{
		i:		buildIndex(),
		icopy:	buildIndex(),
		imx:	sync.RWMutex{},

		started:	false,
		stop:	make(chan bool),
		tch:	make(chan *task, msgQueueSize),
		tq:		make([]*task, 0),
		wg:		sync.WaitGroup{},
		mx:		sync.Mutex{},
	}
}

// Start implements base.Base.
func (b *basicBase) Start() {
	b.mx.Lock()
	defer b.mx.Unlock()

	if b.started {
		return
	}

	b.wg.Add(1)
	go b.start()
}

func (b *basicBase) start() {
	defer b.wg.Done()

	try, get, stop := make(chan bool, 1), make(chan bool), make(chan bool, 1)
	go b.asyncLock(try, get, stop)

	if len(b.tq) != 0 {
		// resolve dirty index copy first
		stopped := b.dirty(try, get, stop)
		if stopped {
			return
		}
	}

	for {
		select {
		case <- b.stop:
			return
		
		case t := <- b.tch:
			
			// Try to apply
			if b.runTask(t) {
				b.tq = append(b.tq, t)
			} else {
				continue	// the request was not valid
			}

			// Begin lock
			try <- true

			// 
			stopped := b.dirty(try, get, stop)
			if stopped {
				return
			}
		}
	}
}

func (b *basicBase) dirty(try chan bool, get chan bool, stop chan bool) bool {
	// The icopy is in dirty state
	for {
		select {
		case <- b.stop:
			stop <- true
			return true

		case <- get:

			// swap index
			b.i, b.icopy = b.icopy, b.i
			b.imx.Unlock()

			// drain dirty tasks
			for _, t := range b.tq {
				b.runTask(t)
			}
			return false

		case t := <- b.tch:
			if b.runTask(t) {
				b.tq = append(b.tq, t)
			} else {
				continue
			}
		}
	}
}

func (b *basicBase) asyncLock(try <-chan bool, get chan<- bool, stop <-chan bool) {
	for {
		select {
		case <- try:
			b.imx.Lock()
			get <- true

		case <- stop:
			b.imx.Unlock()
			return
		}
	}
}

func (b *basicBase) runTask(t *task) bool {
	switch t.cg.Op {
	case types.OperationAdd:
		return b.icopy.Add(t.e, t.cg.Topic)
	case types.OperationDel:
		return b.icopy.Del(t.e, t.cg.Topic)
	}
	return false
}

// Stop implements base.Base.
func (b *basicBase) Stop() {
	b.mx.Lock()
	defer b.mx.Unlock()

	b.stop <- true
	b.started = false
	b.wg.Wait()			// blocking function; wait for graceful stop
}

// Flow implements base.Base.
func (b *basicBase) Flow(m *types.Message) {
	b.imx.RLock()
	defer b.imx.RUnlock()
	
	// fmt.Printf("base received: %v\n", *m)
	b.i.Flow(m)
}

// Apply implements base.Base.
func (b *basicBase) Apply(e emitter.Emitter, cg *types.Change) {
	b.tch <- &task{ e, cg }
}

// Delete implements base.Base.
func (b *basicBase) Delete(e emitter.Emitter) {
	b.tch <- &task{ e, &types.Change{
		Op: types.OperationDel,
		Topic: types.EmptyTopic,
	}}
}
