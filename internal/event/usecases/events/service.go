package events

import (
	"context"

	"github.com/cyworld8x/go-postgres-kubernetes-grpc/internal/event/domain"
	"github.com/cyworld8x/go-postgres-kubernetes-grpc/internal/event/infrastructure/repository"
	"github.com/cyworld8x/go-postgres-kubernetes-grpc/internal/event/infrastructure/repository/postgres"
	_ "github.com/cyworld8x/go-postgres-kubernetes-grpc/pkg/utils"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/rs/zerolog/log"
)

type service struct {
	repo repository.Repository
}

// CloseEventSlots implements UseCase.
func (s *service) CloseEventSlot(ctx context.Context, id uuid.UUID) (*domain.EventSlot, error) {
	eventSlot, error := s.repo.CloseEventSlot(ctx, id)
	if error != nil {
		return nil, error
	}

	eventSlotDto := &domain.EventSlot{
		ID:          eventSlot.ID,
		SlotName:    eventSlot.SlotName,
		Description: eventSlot.Description,
		Price:       eventSlot.Price,
		Capacity:    eventSlot.Capacity,
		Status:      string(eventSlot.Status),
		StartTime:   eventSlot.StartTime.Time,
		EndTime:     eventSlot.EndTime.Time,
		EventID:     eventSlot.EventID,
		Created:     eventSlot.Created.Time,
		Updated:     eventSlot.Updated.Time,
	}

	event, err := s.repo.GetEvent(ctx, eventSlot.EventID)
	if err != nil {
		log.Error().Err(err).Msg("cannot get event")
		return eventSlotDto, err
	}

	eventSlotDto.EventName = event.EventName

	return eventSlotDto, nil

}

// CreateEvent implements UseCase.
func (s *service) CreateEvent(ctx context.Context, eventName string, note string, eventOwnerId uuid.UUID) (*domain.Event, error) {
	event, err := s.repo.CreateEvent(ctx, postgres.CreateEventParams{EventName: eventName, Note: note, EventOwnerID: eventOwnerId})
	if err != nil {
		return nil, err
	}

	return &domain.Event{
		ID:               event.ID,
		EventName:        event.EventName,
		Note:             event.Note,
		Revenue:          event.Revenue,
		Status:           string(event.Status),
		TotalSoldTickets: event.TotalSoldTickets,
		EventOwnerID:     event.EventOwnerID,
		Created:          event.Created.Time,
		Updated:          event.Updated.Time,
	}, nil

}

// GetEvent implements UseCase.
func (s *service) GetEvent(ctx context.Context, id uuid.UUID) (*domain.Event, error) {
	event, err := s.repo.GetEvent(ctx, id)
	if err != nil {
		log.Error().Err(err).Msg("cannot get event")
		return nil, err
	}
	return &domain.Event{
		ID:               event.ID,
		EventName:        event.EventName,
		Note:             event.Note,
		Revenue:          event.Revenue,
		Status:           string(event.Status),
		TotalSoldTickets: event.TotalSoldTickets,
		EventOwnerID:     event.EventOwnerID,
		Created:          event.Created.Time,
		Updated:          event.Updated.Time,
	}, nil
}

// GetEventSlotsByEventId implements UseCase.
func (s *service) GetEventSlotsByEventId(ctx context.Context, eventID uuid.UUID) ([]domain.EventSlot, error) {
	eventSlots, err := s.repo.GetEventSlotsByEventId(ctx, eventID)
	if err != nil {
		return nil, err
	}

	eventSlotsDto := make([]domain.EventSlot, 0, len(eventSlots))
	for _, eventSlot := range eventSlots {
		eventSlotsDto = append(eventSlotsDto, domain.EventSlot{
			ID:          eventSlot.ID,
			SlotName:    eventSlot.SlotName,
			Description: eventSlot.Description,
			Price:       eventSlot.Price,
			Capacity:    eventSlot.Capacity,
			Status:      string(eventSlot.Status),
			StartTime:   eventSlot.StartTime.Time,
			EndTime:     eventSlot.EndTime.Time,
			EventID:     eventSlot.EventID,
			Created:     eventSlot.Created.Time,
			Updated:     eventSlot.Updated.Time,
		})
	}

	return eventSlotsDto, nil
}

// GetEventSlotsById implements UseCase.
func (s *service) GetEventSlotsById(ctx context.Context, id uuid.UUID) (*domain.EventSlot, error) {
	eventSlot, err := s.repo.GetEventSlotsById(ctx, id)
	if err != nil {
		return nil, err
	}

	eventSlotDto := &domain.EventSlot{
		ID:          eventSlot.ID,
		SlotName:    eventSlot.SlotName,
		Description: eventSlot.Description,
		Price:       eventSlot.Price,
		Capacity:    eventSlot.Capacity,
		Status:      string(eventSlot.Status),
		StartTime:   eventSlot.StartTime.Time,
		EndTime:     eventSlot.EndTime.Time,
		EventID:     eventSlot.EventID,
		Created:     eventSlot.Created.Time,
		Updated:     eventSlot.Updated.Time,
	}

	event, err := s.repo.GetEvent(ctx, eventSlot.EventID)
	if err != nil {
		log.Error().Err(err).Msg("cannot get event")
		return eventSlotDto, err
	}

	eventSlotDto.EventName = event.EventName
	return eventSlotDto, nil

}

// OpenEventSlots implements UseCase.
func (s *service) OpenEventSlots(ctx context.Context, arg *domain.CreateEventSlot) (*domain.EventSlot, error) {
	eventSlot, err := s.repo.OpenEventSlots(ctx, postgres.OpenEventSlotsParams{
		SlotName:    arg.SlotName,
		Description: arg.Description,
		Price:       arg.Price,
		Capacity:    arg.Capacity,
		StartTime:   pgtype.Timestamp{Time: arg.StartTime, Valid: true},
		EndTime:     pgtype.Timestamp{Time: arg.EndTime, Valid: true},
		EventID:     arg.EventID,
	})
	if err != nil {
		return nil, err
	}
	return &domain.EventSlot{
		ID:          eventSlot.ID,
		SlotName:    eventSlot.SlotName,
		Description: eventSlot.Description,
		Price:       eventSlot.Price,
		Capacity:    eventSlot.Capacity,
		Status:      string(eventSlot.Status),
		StartTime:   eventSlot.StartTime.Time,
		EndTime:     eventSlot.EndTime.Time,
		EventID:     eventSlot.EventID,
		Created:     eventSlot.Created.Time,
		Updated:     eventSlot.Updated.Time,
	}, nil
}

// StartEvent implements UseCase.
func (s *service) StartEvent(ctx context.Context, id uuid.UUID) (bool, error) {
	eventSlot, err := s.repo.StartEvent(ctx, id)
	if err != nil && eventSlot.Status != "Open" {
		return false, err
	}
	return true, nil
}

// UpdateEventSlots implements UseCase.
func (s *service) UpdateEventSlots(ctx context.Context, id uuid.UUID, slotName string) (*domain.EventSlot, error) {
	eventSlot, err := s.repo.UpdateEventSlots(ctx, postgres.UpdateEventSlotsParams{SlotName: slotName, ID: id})
	if err != nil {
		return nil, err
	}
	return &domain.EventSlot{
		ID:          eventSlot.ID,
		SlotName:    eventSlot.SlotName,
		Description: eventSlot.Description,
		Price:       eventSlot.Price,
		Capacity:    eventSlot.Capacity,
		Status:      string(eventSlot.Status),
		StartTime:   eventSlot.StartTime.Time,
		EndTime:     eventSlot.EndTime.Time,
		EventID:     eventSlot.EventID,
		Created:     eventSlot.Created.Time,
		Updated:     eventSlot.Updated.Time,
	}, nil
}

func NewService(repo repository.Repository) UseCase {
	return &service{
		repo: repo,
	}
}
