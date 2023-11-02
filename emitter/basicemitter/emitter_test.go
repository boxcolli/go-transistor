package basicemitter

import (
	"log"
	"reflect"
	"sync"
	"testing"
	"time"

	"github.com/boxcolli/go-transistor/io/writer/channelwriter"
	"github.com/boxcolli/go-transistor/types"
	"github.com/stretchr/testify/assert"
)

const (
	qsiz = 10
)

var ms = []*types.Message{
	{Topic: types.Topic{"A0"}},
	{Topic: types.Topic{"A1"}},
}

func TestBasicEmitter(t *testing.T) {
	e := NewBasicEmitter(qsiz)
	mch := make(chan *types.Message, qsiz)
	w := channelwriter.NewChannelWriter(mch)

	// Emit messages
	for _, m := range ms {
		e.Emit(m)
	}

	// Timer
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		time.Sleep(1 * time.Second)
		e.Stop()
		wg.Done()
	}()

	e.Work(w)

	// Must receive exact number of messages sent to Emitter.
	for i := 0; i < len(ms); i++ {
		m := <-mch
		assert.Equal(t, true, reflect.DeepEqual(*ms[i], *m))
		log.Println("received:", m)
	}

	wg.Wait()
}

func TestBasicEmitter2(t *testing.T) {
	e := NewBasicEmitter(qsiz)
	mch := make(chan *types.Message, qsiz)
	w := channelwriter.NewChannelWriter(mch)

	// Emit messages
	for _, m := range ms {
		e.Emit(m)
	}

	// Timer
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		time.Sleep(1 * time.Second)
		e.Stop()
		wg.Done()
	}()

	e.Work(w)

	wg.Wait()
}
