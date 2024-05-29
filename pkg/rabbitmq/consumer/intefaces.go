package consumer

import (
	"context"

	amqp "github.com/rabbitmq/amqp091-go"
)

type TaskConsumer func(ctx context.Context, messages <-chan amqp.Delivery)

type EventConsumer interface {
	Configure(...Option) error
	Consume(fn TaskConsumer) error
}
