package config

import (
	"context"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

// Initialize redis
func RedisInit() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:        "127.0.0.1:6379",
		Password:    "",
		DB:          0,
		DialTimeout: 20 * time.Second,
	})

	// Test the Redis connection
	ctx := context.Background()
	_, err := RedisClient.Ping(ctx).Result()
	if err != nil {
		log.Printf("❌ Failed to connect to Redis: %v", err)
		RedisClient = nil
		return
	}

	log.Println("✅ Redis Successfully connected!")
}
