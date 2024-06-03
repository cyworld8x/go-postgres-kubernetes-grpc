package repository

import (
	"context"
	"testing"

	"github.com/cyworld8x/go-postgres-kubernetes-grpc/internal/user/infrastructure/repository/postgres"
	"github.com/cyworld8x/go-postgres-kubernetes-grpc/pkg/utils"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {

	genUser := utils.GenAnUser()
	pwd, err := utils.HashPassword(genUser.Login.Password)
	require.NoError(t, err, "Should not be error in creating a Hash Password")
	user := postgres.CreateUserParams{

		Username:    genUser.Login.UserName,
		Email:       pgtype.Text{String: genUser.Email, Valid: true},
		DisplayName: pgtype.Text{String: genUser.Name.First + " " + genUser.Name.Last, Valid: true},
		Password:    pwd,
		Role:        postgres.Role(postgres.RoleBuyer),
	}

	userCreated, err := testStore.CreateUser(context.Background(), user)
	require.NotEqual(t, userCreated.ID, 0, "The user words should be the same.")
	require.NoError(t, err, "Should not be error in creating a new user")
	require.Equal(t, user.Username, userCreated.Username, "The user words should be the same.")
}
