package config

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func InitRedisConfig() *redis.Client {
	redisdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return redisdb
}

func TestPingRedis(redisdb *redis.Client) string {
	pong, err := redisdb.Ping(ctx).Result()
	if err != nil {
		return err.Error()
	}
	return pong
}
