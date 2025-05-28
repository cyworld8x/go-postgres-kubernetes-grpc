package users

import (
	"context"
	"database/sql"
	"time"

	"github.com/cyworld8x/go-postgres-kubernetes-grpc/internal/user/domain"
	password "github.com/cyworld8x/go-postgres-kubernetes-grpc/pkg/utils"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

type service struct {
	postgresDB  domain.UserRepository
	userRepo    domain.UserDynamoDBRepository
	sessionRepo domain.SessionRepository
}

func NewService(repo domain.UserRepository, userRepo domain.UserDynamoDBRepository, sessionRepo domain.SessionRepository) UseCase {
	return &service{
		postgresDB:  repo,
		userRepo:    userRepo,
		sessionRepo: sessionRepo,
	}
}

// CreateUser implements UseCase.
func (s *service) CreateUser(ctx context.Context, username string, email string, displayName string, pwd string, role string) (domain.User, error) {

	pwdhash, err := password.HashPassword(pwd)
	if err != nil {
		return domain.User{}, err
	}
	// user := postgres.CreateUserParams{
	// 	Username:    username,
	// 	Email:       pgtype.Text{String: email, Valid: true},
	// 	DisplayName: pgtype.Text{String: displayName, Valid: true},
	// 	Password:    pwdhash,
	// 	Role:        postgres.Role(role), // Convert role to postgres.Role
	// 	Code:        "unknown",
	// }

	// dbUser, err := s.postgresDB.CreateUser(ctx, user)
	// if err != nil {
	// 	log.Error().Err(err).Msg("Error creating user")
	// 	return domain.User{}, err
	// }
	now := time.Now().UTC()
	dynamouser := domain.User{
		ID:          uuid.New(),
		Username:    username,
		Email:       sql.NullString{String: email, Valid: true},
		DisplayName: sql.NullString{String: displayName, Valid: true},
		Password:    pwdhash,
		Role:        role, // Convert role to postgres.Role
		Code:        "unknown",
		Status:      true,
		Created:     now,
		Updated:     now,
	}
	err = s.userRepo.CreateUser(ctx, &dynamouser)
	if err != nil {
		log.Error().Err(err).Msg("Error creating user")
		return domain.User{}, err
	}
	return dynamouser, nil

}

// GetLogin implements UseCase.
func (s *service) GetLogin(ctx context.Context, username string) (domain.User, error) {
	// dbUser, err := s.postgresDB.GetLogin(ctx, username)
	// if err != nil {
	// 	log.Error().Err(err).Msg("Error get login user")
	// 	return domain.User{}, err
	// }
	// return domain.User{
	// 	ID:          dbUser.ID,
	// 	Username:    dbUser.Username,
	// 	Email:       sql.NullString{String: dbUser.Email.String, Valid: true},
	// 	DisplayName: sql.NullString{String: dbUser.DisplayName.String, Valid: true},
	// 	Role:        string(dbUser.Role),
	// 	Password:    dbUser.Password,
	// 	Created:     dbUser.Created.Time,
	// 	Updated:     dbUser.Updated.Time,
	// }, nil
	dbUser, err := s.userRepo.GetLogin(ctx, username)
	if err != nil {
		log.Error().Err(err).Msg("Error get login user")
		return domain.User{}, err
	}
	return domain.User{
		ID:          dbUser.ID,
		Username:    dbUser.Username,
		Code:        dbUser.Code,
		Email:       sql.NullString{String: dbUser.Email.String, Valid: true},
		DisplayName: sql.NullString{String: dbUser.DisplayName.String, Valid: true},
		Role:        string(dbUser.Role),
		Password:    dbUser.Password,
		Created:     dbUser.Created,
		Updated:     dbUser.Updated,
	}, nil
}

// GetUser implements UseCase.
func (s *service) GetUser(ctx context.Context, id uuid.UUID) (domain.User, error) {
	//dbUser, err := s.postgresDB.GetUser(ctx, id)

	dbUser, err := s.userRepo.GetUser(ctx, id.String())
	if err != nil {
		log.Error().Err(err).Msg("Error get login user")
		return domain.User{}, err
	}
	return domain.User{
		ID:          dbUser.ID,
		Username:    dbUser.Username,
		Code:        dbUser.Code,
		Email:       sql.NullString{String: dbUser.Email.String, Valid: true},
		DisplayName: sql.NullString{String: dbUser.DisplayName.String, Valid: true},
		Role:        string(dbUser.Role),
		Created:     dbUser.Created,
		Updated:     dbUser.Updated,
	}, nil
}

func (s *service) ChangePassword(ctx context.Context, id uuid.UUID, username string, pwd string) error {
	pwdhash, err := password.HashPassword(pwd)
	if err != nil {
		return err
	}
	// err = s.postgresDB.ChangePassword(ctx, postgres.ChangePasswordParams{
	// 	ID:       id,
	// 	Password: pwdhash,
	// })
	// if err != nil {
	// 	log.Error().Err(err).Msg("Error change password")
	// 	return err
	// }
	err = s.userRepo.ChangePassword(ctx, id.String(), username, pwdhash)
	if err != nil {
		log.Error().Err(err).Msg("Error change password")
		return err
	}
	return nil
}

// GetLogin implements UseCase.
func (s *service) DeleteUser(ctx context.Context, id uuid.UUID) error {

	err := s.userRepo.DeleteUser(ctx, id.String())
	if err != nil {
		log.Error().Err(err).Msg("Error delete user")
		return err
	}
	return nil
}

// GenerateSession implements UseCase.
func (s *service) GenerateSession(ctx context.Context, username string, token string, duration time.Duration) error {
	session, err := s.sessionRepo.GenerateSession(ctx, username, token, duration)
	if err != nil {
		log.Error().Err(err).Msg("Error generating session")
		return err
	}

	// Optionally, you can log the session creation or return the session object
	log.Info().Str("session_id", session.ID).Msg("Session created successfully")
	return nil
}
