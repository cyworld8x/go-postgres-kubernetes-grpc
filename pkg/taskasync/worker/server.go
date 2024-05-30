package worker

import (
	"context"

	config "github.com/cyworld8x/go-postgres-kubernetes-grpc/pkg/taskasync/configuration"
	"github.com/hibiken/asynq"
)

type ConfigMap func(*Worker)

// New creates a new asynq server
type Worker struct {
	server *asynq.Server
	mux    *asynq.ServeMux
}

func New(configuration *config.Configuration) *Worker {
	redisClientOpt := asynq.RedisClientOpt{
		Network:  configuration.Network,
		Addr:     configuration.Addr,
		DB:       configuration.DB,
		Password: configuration.Password,
		Username: configuration.Username,
		PoolSize: configuration.PoolSize,
	}
	server := asynq.NewServer(
		redisClientOpt,
		asynq.Config{
			// Specify how many concurrent workers to use
			Concurrency: 10,
			// Optionally specify multiple queues with different priority.
			Queues: map[string]int{
				"critical": 6,
				"default":  3,
				"low":      1,
			},
			// See the godoc for other configuration options
		},
	)
	return &Worker{
		server: server,
		mux:    asynq.NewServeMux(),
	}
}

type FuncHandler func(ctx context.Context, t *asynq.Task) error

type MiddleWare func(ctx context.Context, t *asynq.Task) asynq.Handler

func WithMap(key string, funcHandler FuncHandler) ConfigMap {
	return func(w *Worker) {
		w.mux.HandleFunc(key, funcHandler)
	}
}

func ProcessTask(ctx context.Context, t *asynq.Task) ConfigMap {
	return func(w *Worker) {
		w.mux.ProcessTask(ctx, t)
	}
}

func Use(middleWare asynq.MiddlewareFunc) ConfigMap {
	return func(w *Worker) {
		w.mux.Use(middleWare)
	}
}

func (p *Worker) Map(options ...ConfigMap) *Worker {
	for _, option := range options {
		option(p)
	}

	return p
}

func (p *Worker) Run() error {
	if err := p.server.Run(p.mux); err != nil {
		return err
	}

	return nil
}
