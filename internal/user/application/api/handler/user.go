package handler

import (
	"net/http"
	"time"

	"github.com/cyworld8x/go-postgres-kubernetes-grpc/internal/user/domain"
	"github.com/cyworld8x/go-postgres-kubernetes-grpc/internal/user/usecases/users"
	"github.com/cyworld8x/go-postgres-kubernetes-grpc/pkg/middleware"
	"github.com/cyworld8x/go-postgres-kubernetes-grpc/pkg/paseto"
	"github.com/cyworld8x/go-postgres-kubernetes-grpc/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func MakeUserHandler(router gin.IRouter, uc users.UseCase) {
	router.Use(middleware.GinLogger())
	router.POST("/user", createUser(uc))
	router.POST("/login", getLogin(uc))
	pasetoMaker, _ := paseto.NewPasetoMaker()
	authRoutes := router.Group("/").Use(middleware.AuthMiddleware(pasetoMaker))
	authRoutes.GET("/user/:id", getUser(uc))
}
func createUser(service users.UseCase) gin.HandlerFunc {
	// Add your code logic here
	return gin.HandlerFunc(func(ctx *gin.Context) {

		var req domain.NewUser
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}
		user, err := service.CreateUser(ctx, req.Username, req.Email, req.DisplayName, req.Password, req.Role)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
		ctx.JSON(http.StatusOK, user)
	})
}

func getLogin(uc users.UseCase) gin.HandlerFunc {
	// Add your code logic here
	return gin.HandlerFunc(func(ctx *gin.Context) {
		var req domain.UserLogin
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, nil)
			return
		}

		user, err := uc.GetLogin(ctx, req.Username)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, "User name doesn't match")
			return
		}

		err = utils.CheckPassword(req.Password, user.Password)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, "Can't login with user name and password.")
			return
		}

		maker, _ := paseto.NewPasetoMaker()
		token, _, err := maker.CreateToken(user.Username, time.Hour)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
		}
		ctx.JSON(http.StatusOK, &domain.UserAccount{
			Username:    user.Username,
			Email:       user.Email.String,
			DisplayName: user.DisplayName.String,
			Role:        user.Role,
			Token:       token,
		})
	})

}

func getUser(service users.UseCase) gin.HandlerFunc {
	// Add your code logic here
	return gin.HandlerFunc(func(ctx *gin.Context) {

		id, err := uuid.Parse(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}
		user, err := service.GetUser(ctx, id)
		if err != nil {
			ctx.JSON(http.StatusOK, "User doesn't exist")
			return
		}
		ctx.JSON(http.StatusOK, user)
	})

}
