package repository

import (
	"context"

	"github.com/cyworld8x/go-postgres-kubernetes-grpc/internal/ticket/infrastructure/repository/postgres"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository interface {
	postgres.Querier
	CreateTicketsByEventSlot(ctx context.Context, eventSlotID string, price float64, total int64) (totalIssued int64, err error)
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

func (m *TicketRepository) CreateTicketsByEventSlot(ctx context.Context, eventSlotID string, price float64, total int64) (totalIssued int64, err error) {
	query := `INSERT INTO "user".tickets(code, event_slot_id, status, price) 
				SELECT uuid_generate_v4(), $1, 'New', $2
				FROM generate_series(1, $3) AS gs; `

	res, err := m.connPool.Exec(ctx, query, eventSlotID, price, total)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected(), nil
}
