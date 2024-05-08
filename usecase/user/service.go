package user

import (
	"context"

	db "github.com/cyworld8x/go-postgres-kubernetes-grpc/db/sqlc"
	password "github.com/cyworld8x/go-postgres-kubernetes-grpc/pkg/utils"
	"github.com/jackc/pgx/v5/pgtype"
)

// Service book usecase
type Service struct {
	repo Repository
}

// NewService create new service
func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

// CreateUser create an user
func (s *Service) CreateUser(username, email, fullname, pwd, role string) (db.User, error) {

	pwd, err := password.HashPassword(pwd)
	if err != nil {
		return db.User{}, err
	}
	user := db.CreateUserParams{
		Username: pgtype.Text{String: username, Valid: true},
		Email:    pgtype.Text{String: email, Valid: true},
		Fullname: pgtype.Text{String: fullname, Valid: true},
		Password: pgtype.Text{String: pwd, Valid: true},
		Role:     pgtype.Text{String: role, Valid: true},
	}
	return s.repo.CreateUser(context.Background(), user)
}

// GetUser get user
func (s *Service) GetUser(id int32) (db.User, error) {
	return s.repo.GetUser(context.Background(), id)
}

// GetLogin get Login
func (s *Service) GetLogin(username string) (db.User, error) {
	return s.repo.GetLogin(context.Background(), pgtype.Text{String: username, Valid: true})
}
