package users

import (
	"context"
	"database/sql"
	"testing"

	"github.com/cyworld8x/go-postgres-kubernetes-grpc/internal/user/domain"
	"github.com/cyworld8x/go-postgres-kubernetes-grpc/internal/user/domain/mock"
	"github.com/cyworld8x/go-postgres-kubernetes-grpc/internal/user/infrastructure/repository/postgres"
	"github.com/cyworld8x/go-postgres-kubernetes-grpc/pkg/utils"
	"github.com/golang/mock/gomock"
	"github.com/jackc/pgx/v5/pgtype"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	userRepository := mock.NewMockUserRepository(ctrl)

	genUser := utils.GenAnUser()
	pwdhash, _ := utils.HashPassword(genUser.Login.Password)
	user := &domain.User{
		Username:    genUser.Login.UserName,
		Email:       sql.NullString{String: genUser.Email, Valid: true},
		DisplayName: sql.NullString{String: genUser.Name.First + " " + genUser.Name.Last, Valid: true},
		Password:    pwdhash,
		Role:        string(domain.Buyer),
	}

	dbUser := &postgres.DbUser{
		Username:    user.Username,
		Email:       pgtype.Text{String: user.Email.String, Valid: true},
		DisplayName: pgtype.Text{String: user.DisplayName.String, Valid: true},
		Password:    user.Password,
		Role:        postgres.RoleBuyer,
	}

	userRepository.EXPECT().
		CreateUser(gomock.Any(), gomock.Any()).
		Times(1).
		Return(*dbUser, nil)

	m := NewService(userRepository, nil)
	u := user
	_, err := m.CreateUser(context.Background(), u.Username, u.Email.String, u.DisplayName.String, u.Password, u.Role)
	assert.Nil(t, err)
	assert.NotEmpty(t, u.Created.IsZero())
}
