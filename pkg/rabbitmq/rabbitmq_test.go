package rabbitmq_test

import (
	"context"
	"testing"
	"time"

	"github.com/cyworld8x/go-postgres-kubernetes-grpc/pkg/rabbitmq"
	"github.com/cyworld8x/go-postgres-kubernetes-grpc/pkg/rabbitmq/consumer"
	"github.com/cyworld8x/go-postgres-kubernetes-grpc/pkg/rabbitmq/publisher"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/stretchr/testify/assert"
)

func TestNewRabbitMQ(t *testing.T) {
	connection, err := rabbitmq.NewRabbitMQ("amqp://rabbitmq:rabbitmq@localhost:5677/")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	if err != nil {
		t.Errorf("Failed to connect to RabbitMQ server: %v", err)
	}

	eventPublisher, err := publisher.NewPublisher(connection)
	if err != nil {
		t.Errorf("Failed to create a new publisher: %v", err)
	}

	eventPublisher.Configure(
		publisher.WithExchangeName("events"),
		publisher.WithType("direct"),
		publisher.WithBindingKey("message:test"),
	)

	eventPublisher.Publish(ctx, []byte("Hello World!"), "text/plain")

	eventConsumer, err := consumer.NewConsumer(connection)
	if err != nil {
		t.Errorf("Failed to create a new consumer: %v", err)
	}
	eventConsumer.Configure(
		consumer.WithQueueName("message"),
		consumer.WithExchangeName("events"),
		consumer.WithBindingKey("message:test"),
		consumer.WithWorkerPoolSize(2),
	)

	taskConsumer := consumer.TaskConsumer(func(ctx context.Context, messages <-chan amqp.Delivery) {
		for message := range messages {
			t.Logf("Received message: %d", message.Body)
			assert.NotEmpty(t, message.Body, "Message body should not be empty")
			err = message.Ack(false)
			if err != nil {
				t.Errorf("Failed to acknowledge message: %v", err)
			}
			assert.Nil(t, err)
		}
	})

	go func() {
		err = eventConsumer.Consume(taskConsumer)
		if err != nil {
			t.Error("failed to start Consumer", err)
			cancel()
		}
	}()

	time.Sleep(10 * time.Second)
	if err != nil {
		t.Errorf("Failed to consume messages: %v", err)
	}
	assert.Nil(t, err)
}
