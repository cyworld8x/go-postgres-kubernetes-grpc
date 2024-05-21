package repository

import (
	"github.com/cyworld8x/go-postgres-kubernetes-grpc/internal/crawler/infrastructure/repository/postgres"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository interface {
	postgres.Querier
}

// Store provides all functions to execute do queries and transactions
type SourceRepository struct {
	connPool *pgxpool.Pool
	*postgres.Queries
}

// Newstere creates a new Store
func NewRepository(connPool *pgxpool.Pool) Repository {
	return &SourceRepository{
		connPool: connPool,
		Queries:  postgres.New(connPool),
	}
}
