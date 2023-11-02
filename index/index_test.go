package index

import (
	"log"
	"testing"

	"github.com/boxcolli/go-transistor/emitter/basicemitter"
	"github.com/boxcolli/go-transistor/io/writer/channelwriter"
	"github.com/boxcolli/go-transistor/types"
)

const (
	qsiz = 10
)

var (
	c = make(chan *types.Message, qsiz)
	w = channelwriter.NewChannelWriter(c)
	e = basicemitter.NewBasicEmitter(qsiz)
)

func TestInode(t *testing.T) {
	go func() {
		e.Work(w)
	} ()

	i := NewInode(nil)
	i.Eset[e] = true

	i.Emit(&types.Message{})

	log.Println("received", *<-c)
}
