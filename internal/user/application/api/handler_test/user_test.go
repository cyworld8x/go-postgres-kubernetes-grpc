package handler_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	api "github.com/cyworld8x/go-postgres-kubernetes-grpc/internal/user/application/api"
	"github.com/cyworld8x/go-postgres-kubernetes-grpc/internal/user/domain/mock"
	"github.com/cyworld8x/go-postgres-kubernetes-grpc/internal/user/infrastructure/repository/postgres"
	middleware "github.com/cyworld8x/go-postgres-kubernetes-grpc/pkg/middleware"
	"github.com/cyworld8x/go-postgres-kubernetes-grpc/pkg/paseto"
	pg "github.com/cyworld8x/go-postgres-kubernetes-grpc/pkg/postgres"
	"github.com/cyworld8x/go-postgres-kubernetes-grpc/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/assert"
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

	dbSource := "postgresql://postgres:postgres@localhost:20241/socialdb?sslmode=disable"
	dynamodbEndpoint := "http://localhost:8000"
	app, err := api.Init(pg.DBConnString(dbSource), dynamodbEndpoint)
	if err != nil {
		t.Fatalf("failed to init app: %v", err)
	}

	return app.Server, err

}

func TestAuthMiddlewareValidToken(t *testing.T) {
	// Create a new Gin router
	serverTest, _ := newTestServer(t)

	// Create PasetoMaker
	pasetoMaker, _ := paseto.NewPasetoMaker()
	// Add the AuthMiddleware to the router
	serverTest.Router.GET(
		"/auth",
		middleware.AuthMiddleware(pasetoMaker),
		func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{})
		},
	)

	// Create a test request with a valid token
	req := httptest.NewRequest(http.MethodGet, "/auth", nil)
	token, payload, err := pasetoMaker.CreateToken("username", time.Hour)
	require.NoError(t, err)
	require.NotEmpty(t, payload)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	// Create a test response recorder
	res := httptest.NewRecorder()

	// Perform the request
	serverTest.Router.ServeHTTP(res, req)

	// Assert that the response status code is 200 OK
	assert.Equal(t, http.StatusOK, res.Code)
}

func TestAuthMiddlewareInvalidToken(t *testing.T) {
	// Create a new Gin router
	router := gin.New()

	// Create PasetoMaker
	pasetoMaker, _ := paseto.NewPasetoMaker()

	// Add the AuthMiddleware to the router
	router.Use(middleware.AuthMiddleware(pasetoMaker))

	// Create a test request with an invalid token
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("Authorization", "Bearer invalid_token")

	// Create a test response recorder
	res := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(res, req)

	// Assert that the response status code is 401 Unauthorized
	assert.Equal(t, http.StatusUnauthorized, res.Code)
}

func TestAuthMiddlewareMissingToken(t *testing.T) {
	// Create a new Gin router
	router := gin.New()

	// Create PasetoMaker
	pasetoMaker, _ := paseto.NewPasetoMaker()

	// Add the AuthMiddleware to the router
	router.Use(middleware.AuthMiddleware(pasetoMaker))

	// Create a test request without the Authorization header
	req := httptest.NewRequest(http.MethodGet, "/", nil)

	// Create a test response recorder
	res := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(res, req)

	// Assert that the response status code is 401 Unauthorized
	assert.Equal(t, http.StatusUnauthorized, res.Code)
}
