package routeindex

import (
	"fmt"
	"log"
	"sync"
	"testing"
	"time"

	"github.com/boxcolli/go-transistor/emitter/basicemitter"
	"github.com/boxcolli/go-transistor/index"
	"github.com/boxcolli/go-transistor/io"
	"github.com/boxcolli/go-transistor/io/bus/channelbus"
	"github.com/boxcolli/go-transistor/io/writer/channelwriter"
	"github.com/boxcolli/go-transistor/io/writer/slicewriter"
	"github.com/boxcolli/go-transistor/types"
	"github.com/stretchr/testify/assert"
)

const (
	qsiz = 10
	csiz = 10
)

type pair struct {
	e	index.Entry
	t	types.Topic
}

var ps = []pair{
	{
		channelbus.NewChannelBus(qsiz),
		types.EmptyTopic,
	},
	{
		channelbus.NewChannelBus(qsiz),
		types.Topic{"A0"},
	},
	{
		channelbus.NewChannelBus(qsiz),
		types.Topic{"A0", "B0"},
	},
}

var cs = []chan *types.Message {
	make(chan *types.Message, csiz),
	make(chan *types.Message, csiz),
	make(chan *types.Message, csiz),
}

var ws = []io.StreamWriter{
	channelwriter.NewChannelWriter(cs[0]),
	channelwriter.NewChannelWriter(cs[1]),
	channelwriter.NewChannelWriter(cs[2]),
}

var es = []index.Entry{
	ps[0].e,
	ps[1].e,
	ps[2].e,
}

var ms = []*types.Message{
	{ Topic: types.EmptyTopic },
	{ Topic: types.Topic{"A0"} },
	{ Topic: types.Topic{"A0", "B0"} },
}

func TestFlow(t *testing.T) {
	// Start emitter
	wg := sync.WaitGroup{}
	{
		for i := 0; i < len(es); i++ {
			// work emitter
			wg.Add(1)
			go func(i int, e index.Entry, w io.StreamWriter) {
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

	// // Test..
	// {
	// 	for i := 0; i < len(es); i++ {
	// 		es[i].Emit(ms[0])
	// 	}
	// }

	index := newRouteIndex()
	{
		// Add
		for _, p := range ps {
			index.Add(p.e, p.t)
		}
		// printInode(index.i)

		// Flow
		for _, m := range ms {
			index.Flow(m)
			log.Printf("flow %v\n", *m)
		}
	}

	close(stop)

	wg.Wait()

	assert.Equal(t, 1, len(cs[0]))
	assert.Equal(t, 2, len(cs[1]))
	assert.Equal(t, 3, len(cs[2]))
}

func TestOne(t *testing.T) {
	w, sw := slicewriter.NewSliceWriter()
	e := basicemitter.NewBasicEmitter(qsiz)

	// Start emitter
	go e.Work(w)

	index := newRouteIndex()
	index.Add(e, types.EmptyTopic)
	index.Add(e, ps[1].t)
	index.Flow(ms[0])
	index.Flow(ms[1])
	
	for i := 0; i < 2; i++ {
		time.Sleep(1 * time.Second)
		sw.MX.Lock()
		for _, m := range sw.M {
			log.Println("SliceWriter.M[]:", *m)
		}
		sw.MX.Unlock()
	}
}

func TestIndexTree(t *testing.T) {
	index := newRouteIndex()

	// Add
	for _, p := range ps {
		index.Add(p.e, p.t)
	}

	printInode(index.i)
}

func TestDel(t *testing.T) {
	index := newRouteIndex()
	// Add
	for _, p := range ps {
		index.Add(p.e, p.t)
	}

	printInode(index.i)
	index.Del(ps[2].e, ps[2].t)
	printInode(index.i)
	fmt.Println()
	printV(index.v)
}

func printInode(i *index.Inode) {
	_printInode(i, "", 0)
}
func _printInode(i *index.Inode, token string, step int) {
	for x := 0; x < step; x++ { fmt.Print("\t") }
	fmt.Printf("%d [%s]inode.Eset\n", step, token)

	for e := range i.Eset {
		for x := 0; x < step; x++ { fmt.Print("\t") }
		fmt.Printf("- %v\n", &e)
	}

	for token, next := range i.Next {
		_printInode(next, token, step + 1)
	}
}
func printV(V map[index.Entry]*index.Vnode) {
	fmt.Printf("V\n")
	for e, v := range V {
		fmt.Printf("- [%v]\n", &e)
		_printV(v, "", 1)
	}
}
func _printV(v *index.Vnode, token string, step int) {
	for x := 0; x < step; x++ { fmt.Print("\t") }
	fmt.Printf("%d [%s]vnode.Next\n", step, token)

	for t, vnext := range v.Next {
		for x := 0; x < step; x++ { fmt.Print("\t") }
		_printV(vnext, t, step + 1)
	}
}

func newRouteIndex() *routeIndex {
	return &routeIndex{
		i: index.NewInode(nil),
		v:	make(map[index.Entry]*index.Vnode),
	}
}
