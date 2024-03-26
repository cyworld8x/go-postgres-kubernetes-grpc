package user

import (
	"testing"
	"time"

	mockdb "github.com/cyworld8x/go-postgres-kubernetes-grpc/db/mock"
	entity "github.com/cyworld8x/go-postgres-kubernetes-grpc/db/sqlc"
	"github.com/cyworld8x/go-postgres-kubernetes-grpc/util"
	"github.com/golang/mock/gomock"
	"github.com/jackc/pgx/v5/pgtype"

	"github.com/stretchr/testify/assert"
)

func newFixtureUser() *entity.User {
	return &entity.User{
		ID:        util.RandomId(),
		Username:  pgtype.Text{String: "phambchung", Valid: true},
		Email:     pgtype.Text{String: "phambchung@gmail.com", Valid: true},
		Password:  pgtype.Text{String: "password", Valid: true},
		Fullname:  pgtype.Text{String: "Chung Pham", Valid: true},
		Role:      pgtype.Text{String: "Engineer", Valid: true},
		CreatedAt: pgtype.Timestamp{Valid: true, Time: time.Now()},
	}
}

func Test_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	store := mockdb.NewMockStore(ctrl)

	store.EXPECT().
		CreateUser(gomock.Any(), gomock.Any()).
		Return(*newFixtureUser(), nil)

	m := NewService(store)
	u := newFixtureUser()
	_, err := m.CreateUser(u.Username.String, u.Email.String, u.Fullname.String, u.Password.String, u.Role.String)
	assert.Nil(t, err)
	assert.False(t, u.CreatedAt.Time.IsZero())
}
