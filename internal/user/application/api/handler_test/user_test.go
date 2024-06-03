package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	api "github.com/cyworld8x/go-postgres-kubernetes-grpc/internal/user/application/api"
	"github.com/cyworld8x/go-postgres-kubernetes-grpc/internal/user/domain/mock"
	"github.com/cyworld8x/go-postgres-kubernetes-grpc/internal/user/infrastructure/repository/postgres"
	pg "github.com/cyworld8x/go-postgres-kubernetes-grpc/pkg/postgres"
	"github.com/cyworld8x/go-postgres-kubernetes-grpc/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)

func newFixtureUser() *postgres.DbUser {
	user := utils.GenAnUser()
	return &postgres.DbUser{
		ID:          uuid.New(),
		Username:    user.Login.UserName,
		Email:       pgtype.Text{String: user.Email, Valid: true},
		Password:    user.Login.Password,
		DisplayName: pgtype.Text{String: user.Name.First + user.Name.Last, Valid: true},
		Role:        postgres.RoleBuyer,
		Created:     pgtype.Timestamp{Valid: true, Time: time.Now()},
		Updated:     pgtype.Timestamp{Valid: true, Time: time.Now()},
	}
}

var user = newFixtureUser()

func TestCreateUser(t *testing.T) {

	testCases := []struct {
		name          string
		body          gin.H
		buildStubs    func(store *mock.MockUserRepository)
		checkResponse func(recoder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"username":     user.Username,
				"password":     user.Password,
				"display_name": user.DisplayName,
				"email":        user.Email,
				"role":         user.Role,
			},
			buildStubs: func(store *mock.MockUserRepository) {
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
				"username":     gomock.Nil(),
				"password":     user.Password,
				"display_name": user.DisplayName,
				"email":        user.Email,
				"role":         user.Role,
			},
			buildStubs: func(store *mock.MockUserRepository) {
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

			store := mock.NewMockUserRepository(ctrl)
			tc.buildStubs(store)

			server, _ := newTestServer(t)
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
		buildStubs    func(store *mock.MockUserRepository)
		checkResponse func(recoder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"username": user.Username,
				"password": user.Password,
			},
			buildStubs: func(store *mock.MockUserRepository) {
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
				"password": user.Password,
			},
			buildStubs: func(store *mock.MockUserRepository) {
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
				"username": user.Username,
				"password": "#invalid-password",
			},
			buildStubs: func(store *mock.MockUserRepository) {
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

			store := mock.NewMockUserRepository(ctrl)
			tc.buildStubs(store)

			server, _ := newTestServer(t)
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

func newTestServer(t *testing.T) (*api.Server, error) {
	testConfig := utils.Configuration{
		DbSource:          "postgresql://postgres:postgres@localhost:20241/socialdb?sslmode=disable",
		HTTPServerAddress: "localhost:8080",
	}

	app, err := api.Init(pg.DBConnString(testConfig.DbSource))
	if err != nil {
		t.Fatalf("failed to init app: %v", err)
	}

	return app.Server, err

}
