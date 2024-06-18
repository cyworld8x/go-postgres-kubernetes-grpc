package cache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	*redis.Client
}

func New(configuration redis.Options) IRedisCache {
	return &RedisClient{
		Client: redis.NewClient(&configuration),
	}
}

// Delete implements IRedisCache.
func (r *RedisClient) Delete(ctx context.Context, key string) error {
	return r.Client.Del(ctx, key).Err()
}

// Get implements IRedisCache.
func (r *RedisClient) Get(ctx context.Context, key string) (interface{}, error) {
	return r.Client.Get(ctx, key).Result()
}

// Set implements IRedisCache.
func (r *RedisClient) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return r.Client.Set(ctx, key, value, expiration).Err()
}
