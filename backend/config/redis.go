package config

import (
	"context"
	"fmt"
	"log"
	"os"

	redisclient "github.com/redis/go-redis/v9"
)

// RDB is the global redis client
var RDB *redisclient.Client // Use the alias here

// RedisContext is used for Redis operations
var RedisContext = context.Background()

// ConnectRedis initializes the Redis client
func ConnectRedis() {
	// 1. Get host address from environment variables
	// Fallback to internal Docker Compose service name if env var is missing
	redisHost := os.Getenv("REDIS_HOST")
	if redisHost == "" {
		redisHost = "redis:6379"
	}

	// 2. Initialize the client
	RDB = redisclient.NewClient(&redisclient.Options{
		Addr:     redisHost,
		Password: "", // No password set in docker-compose for simplicity
		DB:       0,  // Use default DB
	})

	// 3. Ping to verify connection
	_, err := RDB.Ping(RedisContext).Result()
	if err != nil {
		log.Fatal("Error connecting to Redis:", err)
	}

	fmt.Println("Successfully connected to Redis broker!")
}
