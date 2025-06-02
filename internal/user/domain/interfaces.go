package domain

import (
	"context"
	"time"

	postgres "github.com/cyworld8x/go-postgres-kubernetes-grpc/internal/user/infrastructure/repository/postgres"
)

type (
	UserRepository interface {
		postgres.Querier
	}
	UserDynamoDBRepository interface {
		CreateUser(ctx context.Context, user *User) error
		GetLogin(ctx context.Context, userName string) (*User, error)
		ChangePassword(ctx context.Context, id string, username string, password string) error
		UpdateUser(ctx context.Context, user *User) error
		DeleteUser(ctx context.Context, id string) error
		GetUser(ctx context.Context, id string) (*User, error)
	}

	SessionRepository interface {
		GetSession(ctx context.Context, sessionID string) (*Session, error)
		DeleteSession(ctx context.Context, sessionID string) error
		UpdateSession(ctx context.Context, session *Session) error
		ListSessions(ctx context.Context) ([]*Session, error)
		GetSessionByUserID(ctx context.Context, username string) ([]*Session, error)
		BlockSession(ctx context.Context, sessionID string) (bool, error)
		GenerateSession(ctx context.Context, username string, token string, duration time.Duration) (*Session, error)
		UnblockSession(ctx context.Context, sessionID string) (bool, error)
		ClearSession(ctx context.Context, token string, username string) error
	}
)
