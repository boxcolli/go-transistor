package redisplug

import (
	"context"
	"strings"
	"testing"

	"github.com/boxcolli/go-transistor/types"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
)

func TestPrintKeyspace(t *testing.T) {
	client := connect(redisAddr)
	formatter := NewBasicRedisFormatter(dbnum, prefix, delim)

	t.Log("keyspace", formatter.PrintKeyspace(cname))
	pubsub := client.PSubscribe(context.Background(), formatter.PrintKeyspace(cname))

	_, err := pubsub.Receive(context.Background())
	if err != nil {
		panic(err)
	}

	t.Log("Loop")
	ch := pubsub.Channel()
	for {
		e, ok := <- ch
		if !ok {
			break
		}
		t.Log("Event:", *e)
	}

	t.Log("Loop end")
}

func TestFormatter(t *testing.T) {
	formatter := NewBasicRedisFormatter(dbnum, prefix, delim)
	
	// Print
	{
		key := formatter.PrintKey(&me)
		t.Log("PrintKey", key)
		
		value := formatter.PrintValue(&me)
		t.Log("PrintValue", value)
	}
	
	// Scan
	{
		var err error
		var msg *redis.Message
		{
			client := connect(redisAddr)
			pubsub := client.PSubscribe(context.Background(), formatter.PrintPSubscribeKeyspace(cname))
			// _, err = pubsub.Receive(context.Background())
			// if err != nil {
			// 	panic(err)
			// }
			ch := pubsub.Channel()
	
			err = client.Set(context.Background(), formatter.PrintKey(&me), formatter.PrintValue(&me), duration).Err()
			assert.NoError(t, err)

			msg = <- ch
	
			t.Log("msg.Channel", msg.Channel)
		}
		
		key := strings.SplitN(msg.Channel, ":", 2)[1]
		t.Log("SplitN(msg.Channel)", key)

		m := types.Member{}
		formatter.ScanKey(key, &m)
		t.Log("ScanKey", m)
	}
}
