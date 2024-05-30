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
	Configure(...Option) *Distributor
	AddTask(task Task, opts ...asynq.Option) error
}
