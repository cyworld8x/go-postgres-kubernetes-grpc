package repository

import (
	"github.com/cyworld8x/go-postgres-kubernetes-grpc/internal/event/infrastructure/repository/postgres"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository interface {
	postgres.Querier
}

// Store provides all functions to execute do queries and transactions
type TicketRepository struct {
	connPool *pgxpool.Pool
	*postgres.Queries
}

// Newstere creates a new Store
func NewRepository(connPool *pgxpool.Pool) Repository {
	return &TicketRepository{
		connPool: connPool,
		Queries:  postgres.New(connPool),
	}
}
