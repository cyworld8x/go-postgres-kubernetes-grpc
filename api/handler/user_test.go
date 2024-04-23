package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/cyworld8x/go-postgres-kubernetes-grpc/api"
	servertest "github.com/cyworld8x/go-postgres-kubernetes-grpc/api/test"
	mockdb "github.com/cyworld8x/go-postgres-kubernetes-grpc/db/mock"
	entity "github.com/cyworld8x/go-postgres-kubernetes-grpc/db/sqlc"
	"github.com/cyworld8x/go-postgres-kubernetes-grpc/util"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)

func newFixtureUser() *entity.User {
	user := util.GenAnUser()
	return &entity.User{
		ID:        util.RandomId(),
		Username:  pgtype.Text{String: user.Login.UserName, Valid: true},
		Email:     pgtype.Text{String: user.Email, Valid: true},
		Password:  pgtype.Text{String: user.Login.Password, Valid: true},
		Fullname:  pgtype.Text{String: user.Name.First + user.Name.Last, Valid: true},
		Role:      pgtype.Text{String: "Engineer", Valid: true},
		CreatedAt: pgtype.Timestamp{Valid: true, Time: time.Now()},
	}
}

var user = newFixtureUser()

func TestCreateUser(t *testing.T) {

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
				"role":     user.Role,
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
		{
			name: "Bad Request",
			body: gin.H{
				"username": gomock.Nil(),
				"password": user.Password,
				"fullname": user.Fullname,
				"email":    user.Email,
				"role":     user.Role,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateUser(gomock.Any(), *user).
					Times(0).
					Return(*user, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
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

			server, _ := servertest.NewTestServer(t)
			recorder := httptest.NewRecorder()

			// Marshal body data to JSON
			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			url := "/user"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)

			server.Router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)
		})
	}
}

func TestGetLogin(t *testing.T) {
	TestCreateUser(t)
	testCases := []struct {
		name          string
		body          gin.H
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(recoder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"username": user.Username.String,
				"password": user.Password.String,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetLogin(gomock.Any(), gomock.Eq(user.Username)).
					Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name: "BadRequest",
			body: gin.H{
				"username": "#invalid-username",
				"password": user.Password.String,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetLogin(gomock.Any(), gomock.Eq(user.Username)).
					Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "Unauthorized",
			body: gin.H{
				"username": user.Username.String,
				"password": "#invalid-password",
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetLogin(gomock.Any(), gomock.Eq(user.Username)).
					Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
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

			server, _ := servertest.NewTestServer(t)
			recorder := httptest.NewRecorder()

			// Marshal body data to JSON
			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			url := "/login"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)

			server.Router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)
		})
	}
}
