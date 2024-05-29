package publisher

import (
	"context"
	"encoding/json"
	"time"

	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Publisher struct {
	// Configuration is the configuration of the publisher.
	ExchangeName        string
	BindingKey          string
	Connection          *amqp.Connection
	Channel             *amqp.Channel
	PublishType         string
	PublishMandatory    bool
	PublishImmediate    bool
	PublishDeliveryMode uint8
}

type Option func(*Publisher)

func WithExchangeName(exchangeName string) Option {
	return func(p *Publisher) {
		p.ExchangeName = exchangeName
	}
}

func WithType(publishType string) Option {
	return func(p *Publisher) {
		p.PublishType = publishType
	}
}

func WithBindingKey(bindingKey string) Option {
	return func(p *Publisher) {
		p.BindingKey = bindingKey
	}
}

func WithPublishMandatory(publishMandatory bool) Option {
	return func(p *Publisher) {
		p.PublishMandatory = publishMandatory
	}
}

func WithPublishImmediate(publishImmediate bool) Option {
	return func(p *Publisher) {
		p.PublishImmediate = publishImmediate
	}
}

func WithPublishDeliveryMode(publishDeliveryMode uint8) Option {
	return func(p *Publisher) {
		p.PublishDeliveryMode = publishDeliveryMode
	}
}

const (
	_exchangeName = "Default Exchange"
	_bindingKey   = "Default Binding Key"

	_publishMandatory = false
	_publishImmediate = false

	_publishType         = "default"
	_publishDeliveryMode = amqp.Persistent
)

// NewPublisher creates a new Publisher struct.
func NewPublisher(connection *amqp.Connection) (*Publisher, error) {
	channel, err := connection.Channel()

	if err != nil {
		return nil, err
	}

	defer channel.Close()

	return &Publisher{
		Connection:          connection,
		Channel:             channel,
		ExchangeName:        _exchangeName,
		PublishType:         _publishType,
		PublishMandatory:    _publishMandatory,
		PublishImmediate:    _publishImmediate,
		PublishDeliveryMode: _publishDeliveryMode,
	}, nil
}

func (p *Publisher) Configure(options []Option) *Publisher {
	for _, option := range options {
		option(p)
	}

	return p
}

func (p *Publisher) Publish(ctx context.Context, body []byte, contentType string) error {
	channel, err := p.Connection.Channel()
	if err != nil {
		return err
	}

	defer channel.Close()

	err = channel.PublishWithContext(ctx, p.ExchangeName, p.BindingKey, p.PublishMandatory, p.PublishImmediate,
		amqp.Publishing{
			ContentType:  contentType,
			Body:         body,
			DeliveryMode: amqp.Persistent,
			MessageId:    uuid.New().String(),
			Timestamp:    time.Now(),
			Type:         p.PublishType,
		},
	)
	if err != nil {
		return err
	}
	return nil
}

func (p *Publisher) PublishMessages(ctx context.Context, messages [][]byte, contentType string) error {
	for _, message := range messages {
		jsonMessage, err := json.Marshal(message)
		if err != nil {
			return err
		}

		err = p.Publish(ctx, jsonMessage, contentType)
		if err != nil {
			return err
		}
	}
	return nil
}
