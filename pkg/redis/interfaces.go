package redis

import (
	"github.com/cyworld8x/go-postgres-kubernetes-grpc/pkg/redis/cache"
	"github.com/cyworld8x/go-postgres-kubernetes-grpc/pkg/redis/pubsub"
)

type IRedis interface {
	NewPublisher() (pubsub.IPublisher, error)
	NewSubscriber() (pubsub.ISubscriber, error)
	Cache() (cache.IRedisCache, error)
	Close() error
}
