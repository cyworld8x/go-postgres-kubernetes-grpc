package pubsub

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type IPublisher interface {
	Publish(ctx context.Context, key string, value interface{}) error
	Close() error
}

type ISubscriber interface {
	Subscribe(ctx context.Context, key string) *redis.PubSub
	Close() error
}
