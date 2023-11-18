package routeindex

import (
	"testing"

	"github.com/boxcolli/go-transistor/index"
	"github.com/boxcolli/go-transistor/io/bus/channelbus"
	"github.com/boxcolli/go-transistor/types"
	"github.com/stretchr/testify/assert"
)

func TestTearup(t *testing.T) {
	{
		I := newRouteIndex()
		topic := types.Topic{"A0", "B0", "C0"}
		bus := channelbus.NewChannelBus(10)
	
		i := I.i
		i.Eset[bus] = true
		v := index.NewVnode(nil, i)
		I.v[bus] = v
		for x := 0; x < len(topic); x++ {
			t := topic[x]
	
			i.Next[t] = index.NewInode(i)
			i = i.Next[t]
			i.Eset[bus] = true
	
			v.Next[t] = index.NewVnode(v, i)
			v = v.Next[t]
		}
	
		printInode(I.i)
		vlast := tearup(v, bus, topic, len(topic) - 1)
		if vlast == I.v[bus] {
			delete(I.v, bus)
		}
	
		assert.Zero(t, len(I.i.Next))
		assert.Zero(t, len(I.v))
		printInode(I.i)
		
	}

	{
		I := newRouteIndex()
		topic := types.Topic{"A0", "B0", "C0"}
		bus := channelbus.NewChannelBus(10)
	
		i := I.i
		i.Eset[bus] = true
		v := index.NewVnode(nil, i)
		I.v[bus] = v
		for x := 0; x < len(topic); x++ {
			t := topic[x]
	
			i.Next[t] = index.NewInode(i)
			i = i.Next[t]
			i.Eset[bus] = true
	
			v.Next[t] = index.NewVnode(v, i)
			v = v.Next[t]
		}

		I.i.Next["A1"] = index.NewInode(I.i)
		I.i.Next["A1"].Eset[bus] = true
		I.v[bus].Next["A1"] = index.NewVnode(I.v[bus], I.i)

	
		printInode(I.i)
		vlast := tearup(v, bus, topic, len(topic) - 1)
		if vlast == I.v[bus] {
			delete(I.v, bus)
		}

		assert.NotZero(t, I.v)

		printInode(I.i)
	}
}
