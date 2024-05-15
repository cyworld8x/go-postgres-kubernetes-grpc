package events

import (
	"context"

	"github.com/cyworld8x/go-postgres-kubernetes-grpc/internal/event/domain"
	"github.com/google/uuid"
)

// UseCase interface
type UseCase interface {
	CloseEventSlot(ctx context.Context, id uuid.UUID) (*domain.EventSlot, error)
	CreateEvent(ctx context.Context, eventName string, note string, eventOwnerId uuid.UUID) (*domain.Event, error)
	GetEvent(ctx context.Context, id uuid.UUID) (*domain.Event, error)
	GetEventSlotsByEventId(ctx context.Context, eventID uuid.UUID) ([]domain.EventSlot, error)
	GetEventSlotsById(ctx context.Context, id uuid.UUID) (*domain.EventSlot, error)
	OpenEventSlots(ctx context.Context, arg *domain.CreateEventSlot) (*domain.EventSlot, error)
	StartEvent(ctx context.Context, id uuid.UUID) (bool, error)
	UpdateEventSlots(ctx context.Context, id uuid.UUID, slotName string) (*domain.EventSlot, error)
}
