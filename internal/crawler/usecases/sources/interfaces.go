package sources

import (
	"context"

	"github.com/cyworld8x/go-postgres-kubernetes-grpc/internal/crawler/domain"
	"github.com/google/uuid"
)

// UseCase interface
type UseCase interface {
	CreateSource(ctx context.Context, arg *domain.Source) (*domain.Source, error)
	GetSource(ctx context.Context, id uuid.UUID) (*domain.Source, error)
	GetSources(ctx context.Context) ([]domain.Source, error)
	UpdateSource(ctx context.Context, arg *domain.Source) (*domain.Source, error)
}
