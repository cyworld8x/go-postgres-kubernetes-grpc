package repository

import (
	"github.com/cyworld8x/go-postgres-kubernetes-grpc/internal/user/infrastructure/repository/postgres"
	"github.com/jackc/pgx/v5/pgxpool"
)

type repository interface {
	postgres.Querier
}

// Store provides all functions to execute do queries and transactions
type Repository struct {
	connPool *pgxpool.Pool
	*postgres.Queries
}

// Newstere creates a new Store
func NewRepository(connPool *pgxpool.Pool) repository {
	return &Repository{
		connPool: connPool,
		Queries:  postgres.New(connPool),
	}
}
