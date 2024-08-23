package services

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

var (
	redisClient *redis.Client
)

func InitializeRedisClient(ctx context.Context) {
	redisURL := os.Getenv("REDIS_URL")
	redisOptions, err := redis.ParseURL(redisURL)
	if err != nil {
		fmt.Println(err.Error())
	}

	redisClient = redis.NewClient(redisOptions)
	_, err = redisClient.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}
}

func GetRedisClient() *redis.Client {
	return redisClient
}
