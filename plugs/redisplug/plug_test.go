package redisplug

import (
	"context"

	"github.com/redis/go-redis/v9"
	// "github.com/stretchr/testify/assert"
)

func connect() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB: 0,
	})

	_, err := client.Do(
		context.Background(),
		"CONFIG", "SET", "notify-keyspace-events", "KEA",
	).Result()
	if err != nil {
		panic("unable to set keyspace events")
	}

	return client
}
