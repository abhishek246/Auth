package redisclient

import (
	"context"
	"sync"

	"github.com/go-redis/redis/v8"
)

var (
	RedisClient *redis.Client
	once        sync.Once
)

func InitRedisClient() *redis.Client {
	once.Do(func() {
		ctx := context.Background()

		RedisClient = redis.NewClient(&redis.Options{
			Addr:     "redis:6379", // Redis server address
			Password: "",           // No password by default
			DB:       0,            // Default DB
		})

		// Ping the Redis server to check the connection
		_, err := RedisClient.Ping(ctx).Result()
		if err != nil {
			panic(err)
		}
	})

	return RedisClient
}
