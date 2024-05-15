package tickets

import (
	"context"
	"fmt"

	"github.com/cyworld8x/go-postgres-kubernetes-grpc/internal/ticket/domain"
	"github.com/cyworld8x/go-postgres-kubernetes-grpc/internal/ticket/infrastructure/repository"
	"github.com/cyworld8x/go-postgres-kubernetes-grpc/internal/ticket/infrastructure/repository/postgres"
	"github.com/cyworld8x/go-postgres-kubernetes-grpc/pkg/utils"
	_ "github.com/cyworld8x/go-postgres-kubernetes-grpc/pkg/utils"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) UseCase {
	return &service{
		repo: repo,
	}
}

// CreateTicketsByEventSlot implements UseCase.
func (s *service) CreateTicketsByEventSlot(ctx context.Context, eventSlotID uuid.UUID, price float64, total int64) (totalIssued int64, err error) {

	return s.repo.CreateTicketsByEventSlot(ctx, eventSlotID.String(), price, total)

}

func (s *service) GetTotalTicketByEventSlot(ctx context.Context, eventSlotID uuid.UUID) (int64, error) {

	return s.repo.GetTotalTicketByEventSlot(ctx, eventSlotID)

}

// GetTicketsByEventSlotId implements UseCase.
func (s *service) GetTicketsByEventSlotId(ctx context.Context, eventSlotID uuid.UUID) ([]domain.UserTicket, error) {
	tickets, err := s.repo.GetTicketsByEventSlotId(ctx, eventSlotID)
	if err != nil {
		return nil, err
	}

	userTickets := make([]domain.UserTicket, 0, len(tickets))
	for _, userticket := range tickets {
		ticket := domain.UserTicket{
			ID:          userticket.ID,
			Code:        userticket.Code,
			EventSlotID: userticket.EventSlotID,
			Status:      string(userticket.Status),
			Price:       userticket.Price,
			Issued:      userticket.Issued.Time,
		}
		if (userticket.BuyerID != pgtype.UUID{}) {
			ticket.BuyerID = userticket.BuyerID.Bytes
		}
		userTickets = append(userTickets, ticket)
	}
	return userTickets, nil
}

// GetTicketsByUserId implements UseCase.
func (s *service) GetTicketsByUserId(ctx context.Context, buyerID uuid.UUID) ([]domain.UserTicket, error) {
	if (buyerID != uuid.UUID{}) {
		tickets, err := s.repo.GetTicketsByUserId(ctx, pgtype.UUID{Bytes: buyerID, Valid: true})
		if err != nil {
			return nil, err
		}

		userTickets := make([]domain.UserTicket, 0, len(tickets))
		for _, userticket := range tickets {
			ticket := domain.UserTicket{
				ID:          userticket.ID,
				Code:        userticket.Code,
				EventSlotID: userticket.EventSlotID,
				Status:      string(userticket.Status),
				Price:       userticket.Price,
				Issued:      userticket.Issued.Time,
			}
			if (userticket.BuyerID != pgtype.UUID{}) {
				ticket.BuyerID = userticket.BuyerID.Bytes
			}
			userTickets = append(userTickets, ticket)
		}
		return userTickets, nil
	}
	return nil, fmt.Errorf("buyerID is empty")
}

// SellTicket implements UseCase.
func (s *service) SellTicket(ctx context.Context, eventSlotID uuid.UUID, buyerID uuid.UUID) (domain.UserTicket, error) {
	sellTicketParams := postgres.SellTicketParams{
		Column2: pgtype.UUID{Bytes: eventSlotID, Valid: true},
		BuyerID: pgtype.UUID{Bytes: buyerID, Valid: true},
	}
	ticket, err := s.repo.SellTicket(ctx, sellTicketParams)
	if err != nil {
		return domain.UserTicket{}, err
	}
	buyerid := utils.PgGuid{
		Bytes: ticket.BuyerID.Bytes,
	}

	return domain.UserTicket{
		ID:          ticket.ID,
		Code:        ticket.Code,
		EventSlotID: ticket.EventSlotID,
		Status:      string(ticket.Status),
		BuyerID:     buyerid.ToUUID(),
		Price:       ticket.Price,
		Issued:      ticket.Issued.Time,
	}, nil
}

// CheckIn implements UseCase.
func (s *service) CheckIn(ctx context.Context, code string) (domain.UserTicket, error) {
	ticket, err := s.repo.CheckIn(ctx, code)
	if err != nil {
		return domain.UserTicket{}, err
	}

	buyer := utils.PgGuid{
		Bytes: ticket.BuyerID.Bytes,
	}

	return domain.UserTicket{
		ID:          ticket.ID,
		Code:        ticket.Code,
		EventSlotID: ticket.EventSlotID,
		BuyerID:     buyer.ToUUID(),
		Status:      string(ticket.Status),
		Price:       ticket.Price,
		Issued:      ticket.Issued.Time,
	}, nil
}
