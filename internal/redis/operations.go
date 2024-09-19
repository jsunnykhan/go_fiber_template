package redis

import (
	"fmt"
	"github.com/go-redis/redis/v8"
)

type redisService struct {
	client *redis.Client
}

func (r *redisService) Set(key string, value string) error {
	return r.client.Set(ctx, key, value, 0).Err() // 0 means no expiration
}

func (r *redisService) Get(key string) (string, error) {
	val, err := r.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", fmt.Errorf("key does not exist")
	}
	if err != nil {
		return "", err
	}
	return val, nil
}

func (r *redisService) Delete(key string) error {
	return r.client.Del(ctx, key).Err()
}
