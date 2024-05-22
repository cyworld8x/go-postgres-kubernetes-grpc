package crawler

import (
	"context"

	"github.com/cyworld8x/go-postgres-kubernetes-grpc/internal/crawler/domain"
)

// UseCase interface
type UseCase interface {
	Get(ctx context.Context, arg *domain.WebSite) (*domain.Entry, error)
}
