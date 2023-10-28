package basicbase

import (
	"fmt"
	"testing"
	"time"

	"github.com/boxcolli/go-transistor/base"
	"github.com/boxcolli/go-transistor/collector/basiccollector"
	"github.com/boxcolli/go-transistor/emitter/basicemitter"
	"github.com/boxcolli/go-transistor/io/reader/channelreader"
	"github.com/boxcolli/go-transistor/io/writer/channel"
	"github.com/boxcolli/go-transistor/types"
)

func TestBasicBase(t *testing.T) {
	// base
	base := NewBasicBase()
	go base.Start()

	// collector
	ch1 := make(chan *types.Message)
	reader := channelreader.NewChannelStreamReader(ch1)

	collector := basiccollector.NewBasicCollector(base)
	go collector.Work(reader)

	// Emitter
	ch2 := make(chan *types.Message)
	writer := channel.NewChannelStreamWriter(ch2)

	emitter := basicemitter.NewBasicEmitter(1024)
	go emitter.Work(writer)
	go printChan(ch2, "emitter #1")

	// test
	base.Apply(emitter, &types.Change{
		Op: types.OperationAdd,
		Topics: []types.Topic{
			[]string{"KOR", "SEOUL", "1001"},
			[]string{"KOR", "INCHEON", "1002"},
			[]string{"USA", "LA", "1003"},
		},
	})

	base.Apply(emitter, &types.Change{
		Op: types.OperationDel,
		Topics: []types.Topic{
			[]string{"KOR", "SEOUL"},
		},
	})

	go sendMsg(base)

	for {
	}
}

func sendMsg(b base.Base) {
	for {
		time.Sleep(time.Second * 2)

		m := &types.Message{
			Topic:  []string{"KOR", "SEOUL"},
			Method: types.MethodCreate,
			Data:   "Hello, World!",
			TP:     time.Now(),
		}
		b.Flow(m)
	}
}

func printChan(ch <-chan *types.Message, title string) {
	for {
		select {
		case m := <-ch:
			fmt.Printf("%s : %s\n", title, m)
		}
	}
}
