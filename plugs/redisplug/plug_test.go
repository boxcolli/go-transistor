package redisplug

import (
	"context"
	"testing"
	"time"

	"github.com/boxcolli/go-transistor/types"
	"github.com/stretchr/testify/assert"
	// "github.com/stretchr/testify/assert"
)

func TestWatch(t *testing.T) {
	client := connect(redisAddr)
	formatter := NewBasicRedisFormatter(dbnum, prefix, delim)
	plug := NewRedisPlug(client, formatter)

	// Me
	{
		err := plug.Me(context.Background(), types.OperationAdd, &me)
		assert.NoError(t, err)
		time.Sleep(1 * time.Second)
	}

	// Watch
	watch, err := plug.Watch(context.Background(), cname, qsize)
	if err != nil {
		panic(err)
	}

	t.Log("Loop")
	for {
		e, ok := <- watch
		if !ok {
			break
		}
		t.Log("Event:", *e)
	}

	t.Log("Loop end")
}

// func TestRedisPlug(t *testing.T) {
// 	client := connect(redisAddr)
// 	formatter := NewBasicRedisFormatter(dbnum, prefix, delim)
// 	plug := NewRedisPlug(client, formatter)
// 	events := []plugs.Event{
// 		{
// 			Op: types.OperationAdd,
// 			Data: m0,
// 		},
// 		{
// 			Op: types.OperationAdd,
// 			Data: m1,
// 		},
// 		{
// 			Op: types.OperationDel,
// 			Data: m1,
// 		},
// 	}
// 	es := []plugs.Event{}
// 	wg := sync.WaitGroup{}

// 	// Set m0
// 	{
// 		err := client.Set(context.Background(), formatter.PrintKey(&m0), formatter.PrintValue(&m0), duration).Err()
// 		assert.NoError(t, err, "Set m0")
// 	}

// 	// Watch
// 	var watch <-chan *plugs.Event
// 	{
// 		var err error
// 		watch, err = plug.Watch(context.Background(), cname, qsize)
// 		assert.NoError(t, err, "Watch")

// 		wg.Add(1)
// 		go func(wg *sync.WaitGroup) {
// 			t.Log("goroutine Watch")
// 			for {
// 				e, ok := <- watch
// 				t.Log("goroutine", e, ok, "<- watch")
// 				if !ok {
// 					break
// 				} else {
// 					es = append(es, *e)
// 				}
// 			}
// 			wg.Done()
// 			t.Log("goroutine Done")
// 		}(&wg)
// 	}

// 	// Me Add
// 	{
// 		err := plug.Me(context.Background(), types.OperationAdd, &me)
// 		assert.NoError(t, err, "Me Add")
// 	}

// 	// Set m1
// 	{
// 		err := client.Set(context.Background(), formatter.PrintKey(&m1), formatter.PrintValue(&m1), duration).Err()
// 		assert.NoError(t, err, "Set m1")
// 	}

// 	time.Sleep(1 * time.Second)

// 	// Del m1
// 	{
// 		err := client.Del(context.Background(), formatter.PrintKey(&m1)).Err()
// 		assert.NoError(t, err, "Del m1")
// 	}

// 	t.Log("Sleeping..")
// 	time.Sleep(1 * time.Second)

// 	// Stop, Set m2
// 	{
// 		plug.Stop(cname)
// 		t.Log("Stop")

// 		err := client.Set(context.Background(), formatter.PrintKey(&m2), formatter.PrintValue(&m2), duration).Err()
// 		assert.NoError(t, err, "Set m2")
// 	}

// 	// Check result
// 	{
// 		t.Log("Waiting..")
// 		wg.Wait()

// 		assert.Equal(t, true, reflect.DeepEqual(events, es))
// 		t.Log("events", events)
// 		t.Log("es", es)
// 	}
// }
