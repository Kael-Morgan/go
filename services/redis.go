package services

import (
	"context"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

var (
	redisClient *redis.Client
)

func InitializeRedisClient(ctx context.Context) {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0, // use default DB
	})

	// Ping the Redis server to check the connection
	_, err := redisClient.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}
}

func GetRedisClient() *redis.Client {
	return redisClient
}
