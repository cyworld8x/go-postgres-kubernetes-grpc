package configuration

import (
	"github.com/hibiken/asynq"
)

type Configuration asynq.RedisClientOpt
