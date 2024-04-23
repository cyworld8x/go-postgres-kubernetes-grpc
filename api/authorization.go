package api

import (
	"errors"
	"net/http"
	"strings"

	"github.com/cyworld8x/go-postgres-kubernetes-grpc/pkg/paseto"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(am *paseto.PasetoMaker) gin.HandlerFunc {

	return gin.HandlerFunc(func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			// Return unauthorized status code if the header is missing or invalid
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}
		// Get the Authorization header value
		token := strings.TrimPrefix(authHeader, "Bearer ")
		payload, err := am.VerifyToken(token)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, err)
			return
		}
		ctx.Set("authorization_payload", payload)
		ctx.Next()
	})

}
