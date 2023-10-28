package redisplug

import (
	"context"
	"testing"
	"time"

	"github.com/boxcolli/go-transistor/types"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
)

const (
	redisAddr = "localhost:6379"
	dbnum = "0"
	prefix = "test"
	delim = "#"
	cname = "c0"
	qsize = 100
	duration = 10 *time.Second
)
var (
	me = types.Member{
		Cname: cname,
		Name: "nme",
	}

	ms = []types.Member{
		{
			Cname: cname,
			Name: "n0",
		},
		{
			Cname: cname,
			Name: "n1",
		},
		{
			Cname: cname,
			Name: "n2",
		},
	}

	m0, m1, m2 = ms[0], ms[1], ms[2]
)

func connect(addr string) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: addr, // e.g. "localhost:6379"
		DB: 0,
	})

	err := client.Do(context.Background(), "FLUSHALL").Err()
	if err != nil {
		panic(err)
	}

	return client
}

func TestConnect(t *testing.T) {
	client := connect(redisAddr)
	key, value := "hello", "world"

	// Set
	{
		err := client.Set(context.Background(), key, value, duration).Err()
		assert.NoError(t, err)
	}

	// Get
	{
		res, err := client.Get(context.TODO(), key).Result()
		assert.NoError(t, err)
		assert.Equal(t, value, res)
	}
}