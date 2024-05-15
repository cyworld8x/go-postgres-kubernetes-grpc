package tickets

import (
	"context"

	"github.com/cyworld8x/go-postgres-kubernetes-grpc/internal/ticket/domain"
	"github.com/google/uuid"
)

// UseCase interface
type UseCase interface {
	CreateTicketsByEventSlot(ctx context.Context, eventSlotID uuid.UUID, price float64, total int64) (totalIssued int64, err error)
	GetTicketsByEventSlotId(ctx context.Context, eventSlotID uuid.UUID) ([]domain.UserTicket, error)
	GetTicketsByUserId(ctx context.Context, buyerID uuid.UUID) ([]domain.UserTicket, error)
	SellTicket(ctx context.Context, eventSlotID uuid.UUID, buyerID uuid.UUID) (domain.UserTicket, error)
	GetTotalTicketByEventSlot(ctx context.Context, eventSlotID uuid.UUID) (int64, error)
}
