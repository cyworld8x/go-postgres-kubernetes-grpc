// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package social

import (
	"context"
)

type Querier interface {
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeleteUser(ctx context.Context, id int32) error
	FollowUser(ctx context.Context, arg FollowUserParams) (Follow, error)
	GetFollowers(ctx context.Context, followedUserID int32) ([]Follow, error)
	GetUser(ctx context.Context, id int32) (User, error)
	ListUsers(ctx context.Context) ([]User, error)
	UnfollowUser(ctx context.Context, arg UnfollowUserParams) error
	UpdateUser(ctx context.Context, arg UpdateUserParams) error
}

var _ Querier = (*Queries)(nil)
