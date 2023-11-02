package routeindex

import (
	"testing"

	"github.com/boxcolli/go-transistor/emitter/basicemitter"
	"github.com/boxcolli/go-transistor/index"
	"github.com/boxcolli/go-transistor/types"
	"github.com/stretchr/testify/assert"
)

func TestTearup(t *testing.T) {
	{
		I := newRouteIndex()
		topic := types.Topic{"A0", "B0", "C0"}
		e := basicemitter.NewBasicEmitter(1)
	
		i := I.i
		i.Eset[e] = true
		v := index.NewVnode(nil, i)
		I.v[e] = v
		for x := 0; x < len(topic); x++ {
			t := topic[x]
	
			i.Next[t] = index.NewInode(i)
			i = i.Next[t]
			i.Eset[e] = true
	
			v.Next[t] = index.NewVnode(v, i)
			v = v.Next[t]
		}
	
		printInode(I.i)
		vlast := tearup(v, e, topic, len(topic) - 1)
		if vlast == I.v[e] {
			delete(I.v, e)
		}
	
		assert.Zero(t, len(I.i.Next))
		assert.Zero(t, len(I.v))
		printInode(I.i)
		
	}

	{
		I := newRouteIndex()
		topic := types.Topic{"A0", "B0", "C0"}
		e := basicemitter.NewBasicEmitter(1)
	
		i := I.i
		i.Eset[e] = true
		v := index.NewVnode(nil, i)
		I.v[e] = v
		for x := 0; x < len(topic); x++ {
			t := topic[x]
	
			i.Next[t] = index.NewInode(i)
			i = i.Next[t]
			i.Eset[e] = true
	
			v.Next[t] = index.NewVnode(v, i)
			v = v.Next[t]
		}

		I.i.Next["A1"] = index.NewInode(I.i)
		I.i.Next["A1"].Eset[e] = true
		I.v[e].Next["A1"] = index.NewVnode(I.v[e], I.i)

	
		printInode(I.i)
		vlast := tearup(v, e, topic, len(topic) - 1)
		if vlast == I.v[e] {
			delete(I.v, e)
		}

		assert.NotZero(t, I.v)

		printInode(I.i)
	}
}
