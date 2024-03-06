package social

import (
	"context"
	"social/util"
	"testing"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {

	genUser := util.GenAnUser()

	user := CreateUserParams{

		Username: pgtype.Text{String: genUser.Login.UserName, Valid: true},
		Email:    pgtype.Text{String: genUser.Email, Valid: true},
		Fullname: pgtype.Text{String: genUser.Name.First + " " + genUser.Name.Last, Valid: true},
		Password: pgtype.Text{String: genUser.Login.Password, Valid: true},
		Role:     pgtype.Text{String: "User", Valid: true},
	}

	userCreated, err := testQueries.CreateUser(context.Background(), user)
	require.NotEqual(t, userCreated.ID, 0, "The user words should be the same.")
	require.NoError(t, err, "Should not be error in creating a new user")
	require.Equal(t, user.Username, userCreated.Username, "The user words should be the same.")

}
