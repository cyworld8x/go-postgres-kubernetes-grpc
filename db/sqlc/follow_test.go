package db

import (
	"context"
	"testing"

	"github.com/cyworld8x/go-postgres-kubernetes-grpc/pkg/utils"
	"github.com/rs/zerolog/log"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)

func TestFollow(t *testing.T) {
	genUser := utils.GenAnUser()
	pwd, _ := utils.HashPassword(genUser.Login.Password)
	log.Printf("Test Follow: username/pwd: %s %s ", genUser.Login.UserName, genUser.Login.Password)
	userParams := CreateUserParams{

		Username: pgtype.Text{String: genUser.Login.UserName, Valid: true},
		Email:    pgtype.Text{String: genUser.Email, Valid: true},
		Fullname: pgtype.Text{String: genUser.Name.First + " " + genUser.Name.Last, Valid: true},
		Password: pgtype.Text{String: pwd, Valid: true},
		Role:     pgtype.Text{String: "User", Valid: true},
	}

	user1, errCreateUser1 := testStore.CreateUser(context.Background(), userParams)
	require.NoError(t, errCreateUser1)

	genUser = utils.GenAnUser()
	pwd1, _ := utils.HashPassword(genUser.Login.Password)
	log.Printf("Test Follow: username/pwd: %s %s ", genUser.Login.UserName, genUser.Login.Password)
	userParams = CreateUserParams{

		Username: pgtype.Text{String: genUser.Login.UserName, Valid: true},
		Email:    pgtype.Text{String: genUser.Email, Valid: true},
		Fullname: pgtype.Text{String: genUser.Name.First + " " + genUser.Name.Last, Valid: true},
		Password: pgtype.Text{String: pwd1, Valid: true},
		Role:     pgtype.Text{String: "User", Valid: true},
	}

	user2, errCreateUser2 := testStore.CreateUser(context.Background(), userParams)
	require.NoError(t, errCreateUser2)

	follow, err := testStore.FollowTx(context.Background(), FollowsTransParam{FollowUserParams: FollowUserParams{FollowingUserID: user2.ID, FollowedUserID: user1.ID}})
	require.NoError(t, err)
	require.NotEmpty(t, follow)
}
