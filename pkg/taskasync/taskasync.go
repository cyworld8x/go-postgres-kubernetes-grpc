package taskasync

import (
	config "github.com/cyworld8x/go-postgres-kubernetes-grpc/pkg/taskasync/configuration"
	"github.com/cyworld8x/go-postgres-kubernetes-grpc/pkg/taskasync/distributor"
	"github.com/cyworld8x/go-postgres-kubernetes-grpc/pkg/taskasync/worker"
)

type Option func(*TaskAsync)

// TaskAsync is an interface that can be used to distribute tasks
type TaskAsync struct {
	configuration *config.Configuration
}

func WithNetwork(network string) Option {
	return func(c *TaskAsync) {
		c.configuration.Network = network
	}
}

func WithAddress(address string) Option {
	return func(c *TaskAsync) {
		c.configuration.Addr = address
	}
}

func WithDB(db int) Option {
	return func(c *TaskAsync) {
		c.configuration.DB = db
	}
}

func WithPassword(password string) Option {
	return func(c *TaskAsync) {
		c.configuration.Password = password
	}
}

func WithUserName(userName string) Option {
	return func(c *TaskAsync) {
		c.configuration.Username = userName
	}
}

func WithPoolSize(poolSize int) Option {
	return func(c *TaskAsync) {
		c.configuration.PoolSize = poolSize
	}
}

const (
	// DefaultPoolSize is the default pool size
	_network = "tcp"
	_addr    = "localhost:6379"
	_db      = 0
)

func New(connection string) *TaskAsync {
	return &TaskAsync{
		configuration: &config.Configuration{
			Network: _network,
			Addr:    connection,
			DB:      _db,
		},
	}
}

func (taskAsync *TaskAsync) NewDistributor() distributor.IDistributor {

	return distributor.New(taskAsync.configuration)
}

func (taskAsync *TaskAsync) NewWorker() worker.IWorker {
	return worker.New(taskAsync.configuration)
}

func (c *TaskAsync) Configure(opts ...Option) TaskAsync {
	for _, opt := range opts {
		opt(c)
	}

	return *c
}
