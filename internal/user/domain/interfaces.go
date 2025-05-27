package domain

import (
	"context"

	postgres "github.com/cyworld8x/go-postgres-kubernetes-grpc/internal/user/infrastructure/repository/postgres"
)

type (
	UserRepository interface {
		postgres.Querier
	}
	DynamoDBRepository interface {
		CreateUser(ctx context.Context, user *User) error
		GetLogin(ctx context.Context, userName string) (*User, error)
		ChangePassword(ctx context.Context, id string, username string, password string) error
		UpdateUser(ctx context.Context, user *User) error
		DeleteUser(ctx context.Context, id string) error
		GetUser(ctx context.Context, id string) (*User, error)
	}
)
