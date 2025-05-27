package users

import (
	"context"

	"github.com/cyworld8x/go-postgres-kubernetes-grpc/internal/user/domain"
	"github.com/google/uuid"
)

// UseCase interface
type UseCase interface {
	CreateUser(ctx context.Context, username string, email string, displayName string, password string, role string) (domain.User, error)
	GetLogin(ctx context.Context, username string) (domain.User, error)
	ChangePassword(ctx context.Context, id uuid.UUID, username string, password string) error
	GetUser(ctx context.Context, Id uuid.UUID) (domain.User, error)
	DeleteUser(ctx context.Context, Id uuid.UUID) error
}
