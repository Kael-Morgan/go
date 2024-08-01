package services

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

var (
	redisClient *redis.Client
)

func InitializeRedisClient() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "redis-18248.c290.ap-northeast-1-2.ec2.redns.redis-cloud.com:18248",
		Password: "RjQUkQZaUa3X618S0lnNGszX09o6SYeF",
		DB:       0, // use default DB
	})

	// Ping the Redis server to check the connection
	ctx := context.Background()
	_, err := redisClient.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}
}

func GetRedisClient() *redis.Client {
	return redisClient
}
