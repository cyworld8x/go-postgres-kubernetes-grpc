package sessions

import (
	"context"
	"time"

	"github.com/cyworld8x/go-postgres-kubernetes-grpc/internal/user/domain"
)

// UseCase interface for session management
type UseCase interface {
	GetSession(ctx context.Context, sessionID string) (*domain.Session, error)
	DeleteSession(ctx context.Context, sessionID string) error
	UpdateSession(ctx context.Context, session *domain.Session) error
	ListSessions(ctx context.Context) ([]*domain.Session, error)
	GetSessionByUserID(ctx context.Context, username string) ([]*domain.Session, error)
	BlockSession(ctx context.Context, sessionID string) (bool, error)
	GenerateSession(ctx context.Context, username string, token string, duration time.Duration) (*domain.Session, error)
	UnblockSession(ctx context.Context, sessionID string) (bool, error)
	ClearSession(ctx context.Context, token string, username string) error
}
