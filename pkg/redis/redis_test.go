package redis_test

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/cyworld8x/go-postgres-kubernetes-grpc/pkg/redis"
	"github.com/stretchr/testify/assert"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func TestNewRedis(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	redisCache := redis.New(redis.WithAddress("localhost:6379"),
		redis.WithDB(1),
		redis.WithNetwork("tcp"),
	)

	defer redisCache.Client.Close()

	// Test cache
	cache, errCache := redisCache.Cache()
	assert.Nil(t, errCache)
	cache.Set(ctx, "test", "test", 100)
	test, errGet := cache.Get(ctx, "test")
	assert.Nil(t, errGet)
	assert.Equal(t, "test", test)

	// Test pubsub
	publisher, err := redisCache.NewPublisher()
	assert.Nil(t, err)

	for i := 0; i < 2; i++ {
		go func() {
			user := User{
				Name:  "Test User" + fmt.Sprint(i),
				Email: "email@email.com",
			}
			payload, _ := json.Marshal(user)
			err := publisher.Publish(ctx, "task-message", payload)
			assert.Nil(t, err)
		}()
	}

	subscriber, err := redisCache.NewSubscriber()
	sub := subscriber.Subscribe(ctx, "task-message")
	assert.Nil(t, err)
	user := User{}
	go func() {
		for {
			msg, err := sub.ReceiveMessage(ctx)
			if err != nil {
				t.Errorf("Failed to receive message: %v", err)
			}
			err = json.Unmarshal([]byte(msg.Payload), &user)
			if err != nil {
				t.Errorf("Failed to unmarshal message: %v", err)
			}
			t.Logf("Received message: %v", user)
		}

	}()

	time.Sleep(10 * time.Second)
	if err != nil {
		t.Errorf("Failed to subscribe messages: %v", err)
	}
	assert.Nil(t, err)
}
