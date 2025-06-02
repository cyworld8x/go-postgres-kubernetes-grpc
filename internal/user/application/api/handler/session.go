package handler

import (
	"net/http"
	"strings"
	"time"

	"github.com/cyworld8x/go-postgres-kubernetes-grpc/internal/user/usecases/sessions"
	"github.com/cyworld8x/go-postgres-kubernetes-grpc/pkg/middleware"
	"github.com/cyworld8x/go-postgres-kubernetes-grpc/pkg/paseto"
	"github.com/gin-gonic/gin"
)

func MakeSessionHandler(router gin.IRouter, uc sessions.UseCase) {
	router.Use(middleware.GinLogger())
	router.GET("/session/:id", getSession(uc))
	pasetoMaker, _ := paseto.NewPasetoMaker()
	authRoutes := router.Group("/").Use(middleware.AuthMiddleware(pasetoMaker))
	authRoutes.GET("/sessions/user/:username", getSessionByUserID(uc))
	authRoutes.DELETE("/session/user/:id", deleteSession(uc))
	authRoutes.DELETE("/session/:username", clearSession(uc))
	authRoutes.PUT("/session/block/:id", blockSession(uc))
	authRoutes.POST("/session/renew", renewToken(uc))
	authRoutes.GET("/sessions", listSessions(uc))
}

func getSession(uc sessions.UseCase) gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		sessionID := ctx.Param("id")
		session, err := uc.GetSession(ctx, sessionID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve session"})
			return
		}
		ctx.JSON(http.StatusOK, session)
	})
}

func deleteSession(uc sessions.UseCase) gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		sessionID := ctx.Param("id")
		err := uc.DeleteSession(ctx, sessionID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete session"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "Session deleted successfully"})
	})
}

func listSessions(uc sessions.UseCase) gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		sessions, err := uc.ListSessions(ctx)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list sessions"})
			return
		}
		ctx.JSON(http.StatusOK, sessions)
	})
}

func getSessionByUserID(uc sessions.UseCase) gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		username := ctx.Param("username")
		sessions, err := uc.GetSessionByUserID(ctx, username)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve sessions for user"})
			return
		}
		ctx.JSON(http.StatusOK, sessions)
	})
}
func blockSession(uc sessions.UseCase) gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		sessionID := ctx.Param("id")
		blocked, err := uc.BlockSession(ctx, sessionID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to block session"})
			return
		}
		if blocked {
			ctx.JSON(http.StatusOK, gin.H{"message": "Session blocked successfully"})
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Session could not be blocked"})
		}
	})
}

// clearSession is a placeholder for a function that clears a session.

func clearSession(uc sessions.UseCase) gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		if token == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"status": "error",
				"error": gin.H{
					"message": "Request is invalid",
					"code":    "BAD_REQUEST",
				}})
			return
		}
		prefix := "Bearer "
		if !strings.HasPrefix(token, prefix) {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"status": "error",

				"error": gin.H{
					"message": "Request is invalid",
					"code":    "BAD_REQUEST",
				}})
			return
		}

		username := ctx.Param("username")
		if username == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Request is invalid"})
			return
		}

		err := uc.ClearSession(ctx, strings.TrimPrefix(token, prefix), username)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"status": "error",
				"error": gin.H{
					"message": "Failed to clear session" + err.Error(),
					"code":    "INTERNAL_SERVER_ERROR",
				},
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"status": "success",
		})
	})
}

type renewAccessTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

func renewToken(uc sessions.UseCase) gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		var req renewAccessTokenRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		maker, _ := paseto.NewPasetoMaker()
		token, err := maker.VerifyToken(req.RefreshToken)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		session, err := uc.GetSession(ctx, token.ID.String())
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if session == nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Session not found"})
			return
		}

		if session.ExpiresAt < time.Now().Unix() {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Session expired"})
			return
		}

		if session.IsBlocked {
			ctx.JSON(http.StatusForbidden, gin.H{"error": "Session is blocked"})
			return
		}

		// Generate a new session
		newToken, _, err := maker.CreateToken(session.Username, time.Hour)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create new token"})
			return
		}
		newSession, _ := uc.GenerateSession(ctx, session.Username, newToken, time.Hour)
		if newSession == nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate session"})
			return
		}

		err = uc.DeleteSession(ctx, session.ID) // Optionally delete the old session
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete old session"})
		}
		ctx.JSON(http.StatusOK, newSession)
	})
}
