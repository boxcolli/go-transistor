package basicbase

import (
	"time"

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

var ms = []*types.Message{
	{
		Topic:  []string{"KOR", "INCHEON", "1002"},
		Method: types.MethodCreate,
		Data:   []byte("1"),
		TP:     time.Now(),
	},
	{
		Topic:  []string{"USA"},
		Method: types.MethodCreate,
		Data:   []byte("2"),
		TP:     time.Now(),
	},
	{
		Topic:  []string{"USA", "LA"},
		Method: types.MethodCreate,
		Data:   []byte("3"),
		TP:     time.Now(),
	},
	{
		Topic:  []string{"USA", "LA", "1003"},
		Method: types.MethodCreate,
		Data:   []byte("4"),
		TP:     time.Now(),
	},
}