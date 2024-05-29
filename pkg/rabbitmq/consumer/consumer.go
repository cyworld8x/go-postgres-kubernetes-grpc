package consumer

import (
	"context"
	"errors"

	amqp "github.com/rabbitmq/amqp091-go"
)

type ExchangeConfiguration struct {
	Name       string
	Kind       string
	Durable    bool
	AutoDelete bool
	Internal   bool
	NoWait     bool
}

type QueueConfiguration struct {
	Name       string
	BindingKey string
	Durable    bool
	AutoDelete bool
	Exclusive  bool
	NoWait     bool
}

type ConsumerConfiguration struct {
	AutoAck     bool
	Exclusive   bool
	NoLocal     bool
	NoWait      bool
	ConsumerTag string
}

type Configuration struct {
	WorkerPoolSize int
	Queue          QueueConfiguration
	Exchange       ExchangeConfiguration
	Consumer       ConsumerConfiguration
}

const (
	ExchangeKind  = "direct"
	PrefecthCount = 5
	PrefetchSize  = 0
	Global        = false
)

type Option func(*Consumer)

func WithExchangeName(name string) Option {
	return func(c *Consumer) {
		c.configuration.Exchange.Name = name
	}
}

func WithQueueName(name string) Option {
	return func(c *Consumer) {
		c.configuration.Queue.Name = name
	}
}

func WithBindingKey(bindingKey string) Option {
	return func(c *Consumer) {
		c.configuration.Queue.BindingKey = bindingKey
	}
}

func WithWorkerPoolSize(size int) Option {
	return func(c *Consumer) {
		c.configuration.WorkerPoolSize = size
	}
}

func WithQueueConfiguration(queue QueueConfiguration) Option {
	return func(c *Consumer) {
		c.configuration.Queue = queue
	}
}

func WithExchangeConfiguration(exchange ExchangeConfiguration) Option {
	return func(c *Consumer) {
		c.configuration.Exchange = exchange
	}
}

func WithConsumerConfiguration(consumer ConsumerConfiguration) Option {
	return func(c *Consumer) {
		c.configuration.Consumer = consumer
	}
}

// Consumer is a struct that contains the connection to the RabbitMQ server.
type Consumer struct {
	// Connection is the connection to the RabbitMQ server.
	Connection    *amqp.Connection
	configuration Configuration
}

// NewConsumer creates a new Consumer struct.
func NewConsumer(connection *amqp.Connection) (*Consumer, error) {

	defaultConfig := Configuration{

		WorkerPoolSize: 1,
		Queue: QueueConfiguration{
			Name:       "Default Queue",
			BindingKey: "Default Binding Key",
			Durable:    true,
			AutoDelete: false,
			Exclusive:  false,
			NoWait:     false,
		},
		Exchange: ExchangeConfiguration{
			Name:       "Default Exchange",
			Kind:       ExchangeKind,
			Durable:    true,
			AutoDelete: false,
			Internal:   false,
			NoWait:     false,
		},
		Consumer: ConsumerConfiguration{
			AutoAck:     false,
			Exclusive:   false,
			NoLocal:     false,
			NoWait:      false,
			ConsumerTag: "Default Consumer Tag",
		},
	}

	return &Consumer{
		Connection:    connection,
		configuration: defaultConfig,
	}, nil
}

func (c *Consumer) Configure(options []Option) *Consumer {
	for _, option := range options {
		option(c)
	}

	return c
}

// Start starts the consumer
func (c *Consumer) Consume(task TaskConsumer) error {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	channel, err := c.createChannel()
	if err != nil {
		return errors.New("failed to create channel")
	}
	defer channel.Close()

	messages, err := channel.Consume(
		c.configuration.Queue.Name,
		c.configuration.Consumer.ConsumerTag,
		c.configuration.Consumer.AutoAck,
		c.configuration.Consumer.Exclusive,
		c.configuration.Consumer.NoLocal,
		c.configuration.Consumer.NoWait,
		nil,
	)
	if err != nil {
		return errors.New("failed to consume message")
	}

	workerPool := make(chan bool)

	for i := 0; i < c.configuration.WorkerPoolSize; i++ {
		go task(ctx, messages)
	}

	chanErr := <-channel.NotifyClose(make(chan *amqp.Error))
	<-workerPool

	return chanErr
}

// create Channel from the RabbitMQ server.
func (c *Consumer) createChannel() (*amqp.Channel, error) {
	channel, err := c.Connection.Channel()
	if err != nil {
		return nil, err
	}

	err = channel.ExchangeDeclare(
		c.configuration.Exchange.Name,
		c.configuration.Exchange.Kind,
		c.configuration.Exchange.Durable,
		c.configuration.Exchange.AutoDelete,
		c.configuration.Exchange.Internal,
		c.configuration.Exchange.NoWait,
		nil,
	)
	if err != nil {
		return nil, err
	}

	_, err = channel.QueueDeclare(
		c.configuration.Queue.Name,
		c.configuration.Queue.Durable,
		c.configuration.Queue.AutoDelete,
		c.configuration.Queue.Exclusive,
		c.configuration.Queue.NoWait,
		nil,
	)
	if err != nil {
		return nil, err
	}

	err = channel.QueueBind(
		c.configuration.Queue.Name,
		c.configuration.Queue.BindingKey,
		c.configuration.Exchange.Name,
		c.configuration.Queue.NoWait,
		nil,
	)
	if err != nil {
		return nil, errors.New("failed to bind queue")
	}

	err = channel.Qos(
		PrefecthCount,
		PrefetchSize,
		Global,
	)

	if err != nil {
		return nil, errors.New("failed to set QoS")
	}

	return channel, nil
}
