package basicbase

import (
	"log"
	"sync"
	"testing"
	"time"

	"github.com/boxcolli/go-transistor/base"
	"github.com/boxcolli/go-transistor/emitter"
	"github.com/boxcolli/go-transistor/emitter/basicemitter"
	"github.com/boxcolli/go-transistor/index/routeindex"
	"github.com/boxcolli/go-transistor/io"
	"github.com/boxcolli/go-transistor/io/writer/channelwriter"
	"github.com/boxcolli/go-transistor/types"
	"github.com/stretchr/testify/assert"
)

const (
	qsiz = 10
	csiz = 10
)

var cs = []chan *types.Message {
	make(chan *types.Message, csiz),
	make(chan *types.Message, csiz),
}

var ws = []io.StreamWriter{
	channelwriter.NewChannelWriter(cs[0]),
	channelwriter.NewChannelWriter(cs[1]),
}

var es = []emitter.Emitter{
	basicemitter.NewBasicEmitter(qsiz),
	basicemitter.NewBasicEmitter(qsiz),
}

// var ms = []*types.Message{
// 	{ Topic: types.EmptyTopic },
// 	{ Topic: types.Topic{"A0"} },
// 	{ Topic: types.Topic{"A0", "B0"} },
// }

func TestBasicBase(t *testing.T) {
	log.Println("start test")

	go printChan(cs[0], "e0")
	go printChan(cs[1], "e1")

	// Start emitter
	wg := sync.WaitGroup{}
	{
		for i := 0; i < len(es); i++ {
			// work emitter
			wg.Add(1)
			go func(i int, e emitter.Emitter, w io.StreamWriter) {
				log.Printf("emitter[%d] working\n", i)
				e.Work(w)
				log.Printf("emitter[%d] done\n", i)
				wg.Done()
			} (i, es[i], ws[i])
		}
	}

	// Schedule stop
	stop := make(chan bool)
	go func () {
		<- stop
		time.Sleep(3 * time.Second)
		for i := 0; i < len(es); i++ {
			es[i].Stop()
		}
	} ()

	// base
	base := NewBasicBase(routeindex.NewRouteIndex, qsiz)
	base.Start()

	// test
	for _, cg := range cgs1 {
		base.Apply(es[0], cg)
	}
	for _, cg := range cgs2 {
		base.Apply(es[1], cg)
	}
	for _, cg := range cgs3 {
		base.Apply(es[0], cg)
	}
	for _, cg := range cgs4 {
		base.Apply(es[1], cg)
	}

	for i := 0; i < 10; i++ {
		time.Sleep(time.Second * 2)
		sendMsgOnce(base)
	}

	// close(stop)
	// wg.Wait()

	// assert.Equal(t, 4, len(cs[0]))
	assert.Equal(t, 0, len(cs[1]))
}

func sendMsgOnce(b base.Base) {
	for _, m := range ms {
		b.Flow(m)
	}
	log.Printf("sended --------\n")
}

func printChan(ch <-chan *types.Message, title string) {
	for {
		m := <-ch
		log.Printf("%s : %s\n", title, m.Data)
	}
}
