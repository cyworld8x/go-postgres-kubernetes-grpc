package rabbitmq

import (
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

const (
	_reTryTimes = 3
)

// RabbitMQ is a struct that contains the connection to the RabbitMQ server.
type RabbitMQ struct {
	// Connection is the connection to the RabbitMQ server.
	Connection *amqp.Connection
}

// NewRabbitMQ creates a new RabbitMQ struct.
func NewRabbitMQ(url string) (*amqp.Connection, error) {

	var (
		connection *amqp.Connection
		tried      int64
		errConnect error
	)

	for tried < _reTryTimes {

		conn, err := amqp.Dial(url)

		if err != nil {
			errConnect = err
			tried++
			// Sleep for 1 second before retrying
			time.Sleep(time.Second)
		} else {
			connection = conn
			break
		}
	}

	if errConnect != nil {
		return nil, errConnect
	}

	return connection, nil
}
