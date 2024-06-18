package pubsub

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	*redis.Client
}

// Publish implements IPublisher.
// Subtle: this method shadows the method (*Client).Publish of RedisClient.Client.
func (r *RedisClient) Publish(ctx context.Context, key string, value interface{}) error {
	return r.Client.Publish(ctx, key, value).Err()
}

// Subscribe implements IConsumer.
// Subtle: this method shadows the method (*Client).Subscribe of RedisClient.Client.
func (r *RedisClient) Subscribe(ctx context.Context, key string) *redis.PubSub {
	return r.Client.Subscribe(ctx, key)
}

func NewPublisher(configuration redis.Options) IPublisher {
	return &RedisClient{
		Client: redis.NewClient(&configuration),
	}
}

func NewSubscriber(configuration redis.Options) ISubscriber {
	return &RedisClient{
		Client: redis.NewClient(&configuration),
	}
}

func Close(client *RedisClient) error {
	return client.Close()
}
