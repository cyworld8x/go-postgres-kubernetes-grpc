package db

import (
	"context"
	"testing"

	"github.com/cyworld8x/go-postgres-kubernetes-grpc/util"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {

	genUser := util.GenAnUser()
	pwd, err := util.HashPassword(genUser.Login.Password)
	require.NoError(t, err, "Should not be error in creating a Hash Password")
	user := CreateUserParams{

		Username: pgtype.Text{String: genUser.Login.UserName, Valid: true},
		Email:    pgtype.Text{String: genUser.Email, Valid: true},
		Fullname: pgtype.Text{String: genUser.Name.First + " " + genUser.Name.Last, Valid: true},
		Password: pgtype.Text{String: pwd, Valid: true},
		Role:     pgtype.Text{String: "User", Valid: true},
	}

	userCreated, err := testStore.CreateUser(context.Background(), user)
	require.NotEqual(t, userCreated.ID, 0, "The user words should be the same.")
	require.NoError(t, err, "Should not be error in creating a new user")
	require.Equal(t, user.Username, userCreated.Username, "The user words should be the same.")
}
