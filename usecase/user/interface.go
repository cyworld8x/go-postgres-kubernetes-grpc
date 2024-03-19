package user

import (
	"context"

	entity "github.com/cyworld8x/go-postgres-kubernetes-grpc/db/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
)

// Reader interface
type Reader interface {
	GetUser(ctx context.Context, id int32) (entity.User, error)
	GetLogin(ctx context.Context, username pgtype.Text) (entity.User, error)
	DeleteUser(ctx context.Context, id int32) error
	ListUsers(ctx context.Context) ([]entity.User, error)
}

// Writer book writer
type Writer interface {
	CreateUser(ctx context.Context, arg entity.CreateUserParams) (entity.User, error)
	DeleteUser(ctx context.Context, id int32) error
}

// Repository interface
type Repository interface {
	Reader
	Writer
}

// UseCase interface
type UseCase interface {
	CreateUser(username, email, fullname, password, role string) (entity.User, error)
	GetLogin(username string) (entity.User, error)
	GetUser(Id int32) (entity.User, error)
}
