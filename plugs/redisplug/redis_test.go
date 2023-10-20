package redisplug

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRedis(t *testing.T) {
	client := connect(redisAddr)
	keyspace := "__keyspace@0__:"

	t.Log("Subscribe:", keyspace)
	pubsub := client.PSubscribe(context.Background(), keyspace)
	
	_, err := pubsub.Receive(context.Background())
	if err != nil {
		panic(err)
	}

	ch := pubsub.Channel()

	t.Log("Loop")
	for {
		m, ok := <- ch
		if !ok {
			break
		}
		t.Log("<- ch:", m)
	}
}

func TestKeys(t *testing.T) {
	client := connect(redisAddr)
	formatter := NewBasicRedisFormatter(dbnum, prefix, delim)
	res, err := client.Keys(context.Background(), formatter.PrintKeyspace(cname)).Result()
	assert.NoError(t, err, "client.Keys()")

	t.Log("Keyspace", formatter.PrintKeyspace(cname))
	for _, key := range res {
		t.Log("key", key)
	}
}
