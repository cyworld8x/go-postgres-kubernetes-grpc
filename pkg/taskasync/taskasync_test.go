package taskasync_test

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/cyworld8x/go-postgres-kubernetes-grpc/pkg/taskasync"
	"github.com/cyworld8x/go-postgres-kubernetes-grpc/pkg/taskasync/distributor"
	"github.com/cyworld8x/go-postgres-kubernetes-grpc/pkg/taskasync/worker"
	"github.com/hibiken/asynq"
	"github.com/stretchr/testify/assert"
)

func TestNewTaskAsync(t *testing.T) {
	taskAsync := taskasync.New("localhost:6379")

	taskAsync.Configure(
		taskasync.WithPoolSize(10),
	)

	taskDistributor := taskAsync.NewDistributor()
	defer taskDistributor.Close()

	go func() {
		for i := 0; i < 2; i++ {
			testMessage := []byte("Test Message " + fmt.Sprint(i) + "!")
			err := taskDistributor.AddTask(distributor.Task{
				TypeName: "task:message",
				Payload:  testMessage,
			}, asynq.MaxRetry(3), asynq.Queue("critical"))
			assert.Nil(t, err)
		}
		funcHandler := worker.FuncHandler(func(ctx context.Context, t *asynq.Task) error {
			p := string(t.Payload())
			log.Println(p)
			return nil
		})

		taskAsync.NewWorker().Map(
			worker.WithMap("task:message", funcHandler),
		).Run()
	}()

	time.Sleep(time.Second * 10)
	err := taskDistributor.Close()
	assert.Nil(t, err)

}
