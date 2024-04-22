package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	mockdb "github.com/cyworld8x/go-postgres-kubernetes-grpc/db/mock"
	entity "github.com/cyworld8x/go-postgres-kubernetes-grpc/db/sqlc"
	util "github.com/cyworld8x/go-postgres-kubernetes-grpc/util"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
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

func Test_CreateUser(t *testing.T) {
	user := newFixtureUser()

	testCases := []struct {
		name          string
		body          gin.H
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(recoder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"username": user.Username,
				"password": user.Password,
				"fullname": user.Fullname,
				"email":    user.Email,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateUser(gomock.Any(), *user).
					Times(0).
					Return(*user, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
	}
	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			server, _ := NewTestServer()
			recorder := httptest.NewRecorder()

			// Marshal body data to JSON
			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			url := "/user"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)
		})
	}
}
