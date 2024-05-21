// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package postgres

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	ChangeUserPassword(ctx context.Context, arg ChangeUserPasswordParams) error
	CreateUser(ctx context.Context, arg CreateUserParams) (DbUser, error)
	DeactiveUser(ctx context.Context, id uuid.UUID) error
	GetLogin(ctx context.Context, username string) (DbUser, error)
	GetUser(ctx context.Context, id uuid.UUID) (DbUser, error)
	ListUsers(ctx context.Context) ([]DbUser, error)
	UpdateUserInfo(ctx context.Context, arg UpdateUserInfoParams) error
}

var _ Querier = (*Queries)(nil)