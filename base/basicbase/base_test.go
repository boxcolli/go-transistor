package basicbase

import (
	"fmt"
	"testing"
	"time"

	"github.com/boxcolli/go-transistor/base"
	"github.com/boxcolli/go-transistor/emitter/basicemitter"
	"github.com/boxcolli/go-transistor/io/writer/channel"
	"github.com/boxcolli/go-transistor/types"
)

func TestBasicBase(t *testing.T) {
	fmt.Println("start test")

	// base
	base := NewBasicBase()
	go base.Start()

	// Emitter
	ch1 := make(chan *types.Message)
	writer1 := channel.NewChannelStreamWriter(ch1)
	emitter1 := basicemitter.NewBasicEmitter(1024)
	go emitter1.Work(writer1)
	go printChan(ch1, "emitter #1")

	ch2 := make(chan *types.Message)
	writer2 := channel.NewChannelStreamWriter(ch2)
	emitter2 := basicemitter.NewBasicEmitter(1024)
	go emitter2.Work(writer2)
	go printChan(ch2, "emitter #2")

	// test
	base.Apply(emitter1, &types.Change{
		Op: types.OperationAdd,
		Topics: []types.Topic{
			[]string{"KOR", "SEOUL", "1001"},
			[]string{"KOR", "INCHEON", "1002"},
			[]string{"USA", "LA", "1003"},
		},
	})
	base.Apply(emitter2, &types.Change{
		Op: types.OperationAdd,
		Topics: []types.Topic{
			[]string{"KOR", "SEOUL", "1001"},
			[]string{"USA", "LA"},
		},
	})
	base.Apply(emitter1, &types.Change{
		Op: types.OperationDel,
		Topics: []types.Topic{
			[]string{"KOR", "SEOUL", "1001"},
			[]string{"USA", "LA"},
		},
	})
	base.Apply(emitter2, &types.Change{
		Op: types.OperationDel,
		Topics: []types.Topic{
			[]string{},
		},
	})

	go sendMsg(base)

	for {
	}
}

func sendMsg(b base.Base) {
	for {
		time.Sleep(time.Second * 2)
		fmt.Printf("sended --------\n")

		b.Flow(&types.Message{
			Topic:  []string{"KOR", "INCHEON", "1002"},
			Method: types.MethodCreate,
			Data:   "1",
			TP:     time.Now(),
		})
		b.Flow(&types.Message{
			Topic:  []string{"USA"},
			Method: types.MethodCreate,
			Data:   "2",
			TP:     time.Now(),
		})
		b.Flow(&types.Message{
			Topic:  []string{"USA", "LA"},
			Method: types.MethodCreate,
			Data:   "3",
			TP:     time.Now(),
		})
		b.Flow(&types.Message{
			Topic:  []string{"USA", "LA", "1003"},
			Method: types.MethodCreate,
			Data:   "4",
			TP:     time.Now(),
		})
	}
}

func printChan(ch <-chan *types.Message, title string) {
	for {
		select {
		case m := <-ch:
			fmt.Printf("%s : %s\n", title, m.Data)
		}
	}
}
