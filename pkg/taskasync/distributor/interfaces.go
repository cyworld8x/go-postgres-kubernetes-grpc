package distributor

import (
	"github.com/hibiken/asynq"
)

type Task struct {
	TypeName string
	Payload  []byte
}

type TaskRegister func(task Task, opts []Option) *asynq.Task

type IDistributor interface {
	AddTask(task Task, opts ...asynq.Option) error
	Close() error
}
