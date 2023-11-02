package routeindex

import (
	"testing"

	"github.com/boxcolli/go-transistor/types"
)

var cgs1 = []*types.Change{
	{
		Op: types.OperationAdd,
		Topic: types.Topic{"KOR", "SEOUL", "1001"},
	},
	{
		Op: types.OperationAdd,
		Topic: types.Topic{"KOR", "INCHEON", "1002"},
	},
	{
		Op: types.OperationAdd,
		Topic: types.Topic{"USA", "LA", "1003"},
	},
}

var cgs2 = []*types.Change{
	{
		Op: types.OperationAdd,
		Topic: types.Topic{"KOR", "SEOUL", "1001"},
	},
	{
		Op: types.OperationAdd,
		Topic: types.Topic{"USA", "LA"},
	},
}

var cgs3 = []*types.Change{
	{
		Op: types.OperationDel,
		Topic: types.Topic{"KOR", "SEOUL", "1001"},
	},
	{
		Op: types.OperationDel,
		Topic: types.Topic{"USA", "LA"},
	},
}

var cgs4 = []*types.Change{
	{
		Op: types.OperationDel,
		Topic: types.EmptyTopic,
	},
}

func TestChanges(t *testing.T) {
	index := newRouteIndex()

	for _, cg := range cgs1 {
		index.Add(es[0], cg.Topic)
	}
	for _, cg := range cgs2 {
		index.Add(es[1], cg.Topic)
	}
	for _, cg := range cgs3 {
		index.Del(es[0], cg.Topic)
	}
	for _, cg := range cgs4 {
		index.Del(es[1], cg.Topic)
	}

	printInode(index.i)
	printV(index.v)
}
