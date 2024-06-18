package redis

import (
	"errors"

	"github.com/cyworld8x/go-postgres-kubernetes-grpc/pkg/redis/cache"
	"github.com/cyworld8x/go-postgres-kubernetes-grpc/pkg/redis/pubsub"
	"github.com/redis/go-redis/v9"
)

type Option func(*RedisCache)

type RedisCache struct {
	Client        *redis.Client
	Configuration redis.Options
}

func New(options ...Option) *RedisCache {
	c := &RedisCache{}
	for _, option := range options {
		option(c)
	}
	c.Client = redis.NewClient(&c.Configuration)
	return c
}

func (c *RedisCache) Configure(opts ...Option) RedisCache {
	for _, opt := range opts {
		opt(c)
	}

	return *c
}

func (redis *RedisCache) NewSubscriber() (pubsub.ISubscriber, error) {
	if redis.Client == nil {
		return nil, errors.New("Redis client of Subscriber is nil")
	}
	return &pubsub.RedisClient{
		Client: redis.Client,
	}, nil
}

func (redis *RedisCache) NewPublisher() (pubsub.IPublisher, error) {
	if redis.Client == nil {
		return nil, errors.New("Redis client of Publisher is nil")
	}
	return &pubsub.RedisClient{
		Client: redis.Client,
	}, nil
}

func (redis *RedisCache) Cache() (cache.IRedisCache, error) {
	if redis.Client == nil {
		return nil, errors.New("Redis client of Publisher is nil")
	}
	return &cache.RedisClient{
		Client: redis.Client,
	}, nil
}

func (r *RedisCache) Close() error {
	return r.Client.Close()
}

func WithNetwork(network string) Option {
	return func(c *RedisCache) {
		c.Configuration.Network = network
	}
}

func WithAddress(address string) Option {
	return func(c *RedisCache) {
		c.Configuration.Addr = address
	}
}

func WithDB(db int) Option {
	return func(c *RedisCache) {
		c.Configuration.DB = db
	}
}

func WithPassword(password string) Option {
	return func(c *RedisCache) {
		c.Configuration.Password = password
	}
}

func WithUserName(userName string) Option {
	return func(c *RedisCache) {
		c.Configuration.Username = userName
	}
}
