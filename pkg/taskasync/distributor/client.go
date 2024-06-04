package distributor

import (
	config "github.com/cyworld8x/go-postgres-kubernetes-grpc/pkg/taskasync/configuration"
	"github.com/hibiken/asynq"
)

type Option func(*Distributor)

// Distributor is an interface that can be used to distribute tasks
type Distributor struct {
	client        *asynq.Client
	configuration asynq.RedisClientOpt
}

// func WithNetwork(network string) Option {
// 	return func(c *Distributor) {
// 		c.configuration.Network = network
// 	}
// }

// func WithAddress(address string) Option {
// 	return func(c *Distributor) {
// 		c.configuration.Addr = address
// 	}
// }

// func WithDB(db int) Option {
// 	return func(c *Distributor) {
// 		c.configuration.DB = db
// 	}
// }

// func WithPassword(password string) Option {
// 	return func(c *Distributor) {
// 		c.configuration.Password = password
// 	}
// }

// func WithUserName(userName string) Option {
// 	return func(c *Distributor) {
// 		c.configuration.Username = userName
// 	}
// }

// func WithPoolSize(poolSize int) Option {
// 	return func(c *Distributor) {
// 		c.configuration.PoolSize = poolSize
// 	}
// }

func New(configuration *config.Configuration) IDistributor {
	redisClientOpt := asynq.RedisClientOpt{
		Network:  configuration.Network,
		Addr:     configuration.Addr,
		DB:       configuration.DB,
		Password: configuration.Password,
		Username: configuration.Username,
		PoolSize: configuration.PoolSize,
	}
	client := asynq.NewClient(redisClientOpt)

	return &Distributor{
		client:        client,
		configuration: redisClientOpt,
	}
}

func (c *Distributor) AddTask(task Task, opts ...asynq.Option) error {
	taskRegister := asynq.NewTask(task.TypeName, task.Payload, opts...)
	_, err := c.client.Enqueue(taskRegister)
	return err
}

func (c *Distributor) Close() error {
	return c.client.Close()
}
