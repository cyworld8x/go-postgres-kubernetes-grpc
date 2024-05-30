package taskasync_test

import (
	"context"
	"testing"

	"github.com/cyworld8x/go-postgres-kubernetes-grpc/pkg/taskasync"
	"github.com/cyworld8x/go-postgres-kubernetes-grpc/pkg/taskasync/worker"
	"github.com/hibiken/asynq"
)

func TestNewTaskAsync(t *testing.T) {
	taskAsync := taskasync.NewTaskAsync("redis://localhost:6379/0")

	taskAsync.Configure(
		taskasync.WithPoolSize(10),
	)

	taskAsync.NewWorker().Map(
		worker.WithMap("key", func(ctx context.Context, t *asynq.Task) error {
			return nil
		}),
		worker.WithMap("key2", func(ctx context.Context, t *asynq.Task) error {
			return nil
		}),
	)

}
