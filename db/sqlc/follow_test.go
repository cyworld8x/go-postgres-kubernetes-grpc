package db

import (
	"context"
	"testing"

	"github.com/cyworld8x/go-postgres-kubernetes-grpc/util"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)

func TestFollow(t *testing.T) {
	genUser := util.GenAnUser()

	userParams := CreateUserParams{

		Username: pgtype.Text{String: genUser.Login.UserName, Valid: true},
		Email:    pgtype.Text{String: genUser.Email, Valid: true},
		Fullname: pgtype.Text{String: genUser.Name.First + " " + genUser.Name.Last, Valid: true},
		Password: pgtype.Text{String: genUser.Login.Password, Valid: true},
		Role:     pgtype.Text{String: "User", Valid: true},
	}

	user1, errCreateUser1 := testStore.CreateUser(context.Background(), userParams)
	require.NoError(t, errCreateUser1)

	genUser = util.GenAnUser()

	userParams = CreateUserParams{

		Username: pgtype.Text{String: genUser.Login.UserName, Valid: true},
		Email:    pgtype.Text{String: genUser.Email, Valid: true},
		Fullname: pgtype.Text{String: genUser.Name.First + " " + genUser.Name.Last, Valid: true},
		Password: pgtype.Text{String: genUser.Login.Password, Valid: true},
		Role:     pgtype.Text{String: "User", Valid: true},
	}

	user2, errCreateUser2 := testStore.CreateUser(context.Background(), userParams)
	require.NoError(t, errCreateUser2)

	follow, err := testStore.FollowTx(context.Background(), FollowsTransParam{FollowUserParams: FollowUserParams{FollowingUserID: user2.ID, FollowedUserID: user1.ID}})
	require.NoError(t, err)
	require.NotEmpty(t, follow)
}
