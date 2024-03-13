package social

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Store interface {
	Querier
	FollowTx(ctx context.Context, arg FollowsTransParam) (FollowsTransResult, error)
}

// Store provides all functions to execute do queries and transactions
type SQLStore struct {
	connPool *pgxpool.Pool
	*Queries
}

// Newstere creates a new Store
func NewStore(connPool *pgxpool.Pool) Store {
	return &SQLStore{
		connPool: connPool,
		Queries:  New(connPool),
	}
}
