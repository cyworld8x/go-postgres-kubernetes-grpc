package middleware_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	middleware "github.com/cyworld8x/go-postgres-kubernetes-grpc/api/middleware"
	servertest "github.com/cyworld8x/go-postgres-kubernetes-grpc/api/test"
	"github.com/cyworld8x/go-postgres-kubernetes-grpc/pkg/paseto"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAuthMiddleware_ValidToken(t *testing.T) {
	// Create a new Gin router
	serverTest, _ := servertest.NewTestServer(t)

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

func TestAuthMiddleware_InvalidToken(t *testing.T) {
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

func TestAuthMiddleware_MissingToken(t *testing.T) {
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
