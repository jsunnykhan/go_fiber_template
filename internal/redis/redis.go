package redis

import (
	"context"
	"log"
	"os"
	"strconv"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

var (
	host     = os.Getenv("REDIS_HOST")
	port     = os.Getenv("REDIS_PORT")
	password = os.Getenv("REDIS_PASSWORD")
	db       = os.Getenv("REDIS_DATABASE")
)

type RedisClient interface {
	Set(key string, value string) error
	Get(key string) (string, error)
	Delete(key string) error
}

// NewRedisClient initializes a new Redis client.
func NewRedisClient() RedisClient {
	rdb := redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: password,
		DB:       parseDB(db),
	})

	// Test the connection
	if err := rdb.Ping(ctx).Err(); err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	return &redisService{client: rdb}
}

func parseDB(db string) int {
	if db == "" {
		return 0
	}
	parsedDB, err := strconv.Atoi(db)
	if err != nil {
		log.Fatalf("Invalid REDIS_DB value: %v", err)
	}
	return parsedDB
}
