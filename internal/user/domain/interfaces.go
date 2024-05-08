package domain

import (
	postgres "github.com/cyworld8x/go-postgres-kubernetes-grpc/internal/user/infrastructure/repository/postgres"
)

type (
	UserRepository interface {
		postgres.Querier
	}
)
