package config

import (
	"fmt"
	"github.com/go-redis/redis"
)

var redisClient *redis.Client

func init() {
	//redis
	redisClient = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})
	fmt.Println("redis connected")
}

func GetRedisClient() *redis.Client {
	return redisClient
}
