package main

import (
	"context"
	"fmt"
	"go-beyond/services"
	"time"
)

func main() {
	services.InitializeRedisClient()

	redisClient := services.GetRedisClient()

	ctx := context.Background()

	start := time.Now()
	cart := redisClient.HGetAll(ctx, "initCart").Val()
	ellapsed := time.Since(start)

	fmt.Println(cart["1"], ellapsed)
}
