package users

import (
	"context"
	"database/sql"

	"github.com/cyworld8x/go-postgres-kubernetes-grpc/internal/user/domain"
	"github.com/cyworld8x/go-postgres-kubernetes-grpc/internal/user/infrastructure/repository/postgres"
	password "github.com/cyworld8x/go-postgres-kubernetes-grpc/pkg/utils"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/rs/zerolog/log"
)

type service struct {
	repo domain.UserRepository
}

func NewService(repo domain.UserRepository) UseCase {
	return &service{
		repo: repo,
	}
}

// CreateUser implements UseCase.
func (s *service) CreateUser(ctx context.Context, username string, email string, displayName string, pwd string, role int) (domain.User, error) {

	pwdhash, err := password.HashPassword(pwd)
	if err != nil {
		return domain.User{}, err
	}
	user := postgres.CreateUserParams{
		Username:    username,
		Email:       pgtype.Text{String: email, Valid: true},
		DisplayName: pgtype.Text{String: displayName, Valid: true},
		Password:    pwdhash,
		Role:        int32(role),
		Code:        "unknown",
	}

	dbUser, err := s.repo.CreateUser(ctx, user)
	if err != nil {
		log.Error().Err(err).Msg("Error creating user")
		return domain.User{}, err
	}
	return domain.User{
		ID:          dbUser.ID,
		Username:    dbUser.Username,
		Email:       sql.NullString{String: dbUser.Email.String, Valid: true},
		DisplayName: sql.NullString{String: dbUser.DisplayName.String, Valid: true},
		Password:    dbUser.Password,
		Role:        dbUser.Role,
		Created:     dbUser.Created.Time,
		Updated:     dbUser.Updated.Time,
	}, nil

}

// GetLogin implements UseCase.
func (s *service) GetLogin(ctx context.Context, username string) (domain.User, error) {
	dbUser, err := s.repo.GetLogin(ctx, username)
	if err != nil {
		log.Error().Err(err).Msg("Error get login user")
		return domain.User{}, err
	}
	return domain.User{
		ID:          dbUser.ID,
		Username:    dbUser.Username,
		Email:       sql.NullString{String: dbUser.Email.String, Valid: true},
		DisplayName: sql.NullString{String: dbUser.DisplayName.String, Valid: true},
		Role:        dbUser.Role,
		Created:     dbUser.Created.Time,
		Updated:     dbUser.Updated.Time,
	}, nil
}

// GetUser implements UseCase.
func (s *service) GetUser(ctx context.Context, id uuid.UUID) (domain.User, error) {
	dbUser, err := s.repo.GetUser(ctx, id)
	if err != nil {
		log.Error().Err(err).Msg("Error get login user")
		return domain.User{}, err
	}
	return domain.User{
		ID:          dbUser.ID,
		Username:    dbUser.Username,
		Email:       sql.NullString{String: dbUser.Email.String, Valid: true},
		DisplayName: sql.NullString{String: dbUser.DisplayName.String, Valid: true},
		Role:        dbUser.Role,
		Created:     dbUser.Created.Time,
		Updated:     dbUser.Updated.Time,
	}, nil
}
