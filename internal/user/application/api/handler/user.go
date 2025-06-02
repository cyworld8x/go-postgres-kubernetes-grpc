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
	authRoutes.PUT("/user/:id", changePassword(uc))
	authRoutes.DELETE("/user/:id", deleteUser(uc))
}
func createUser(service users.UseCase) gin.HandlerFunc {
	// Add your code logic here
	return gin.HandlerFunc(func(ctx *gin.Context) {

		var req domain.NewUser
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"status":  "error",
				"message": "Can't create user",
				"error": map[string]interface{}{
					"code":    "user_creation_error",
					"message": err.Error(),
				},
			})
			return
		}
		_, err := service.CreateUser(ctx, req.Username, req.Email, req.DisplayName, req.Password, req.Role)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"status":  "error",
				"message": "Can't create user",
				"error": map[string]interface{}{
					"code":    "user_creation_error",
					"message": err.Error(),
				},
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"message": "User registered successfully",
			"status":  "success",
		})
		return

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
			ctx.JSON(http.StatusOK, gin.H{
				"status": "error",
				"error": map[string]interface{}{
					"message": "Can't login with user name and password.",
					"code":    "LOGIN_ERROR",
				},
			})
			return
		}

		err = utils.CheckPassword(req.Password, user.Password)
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"status": "error",
				"error": map[string]interface{}{
					"message": "Can't login with user name and password.",
					"code":    "LOGIN_ERROR",
				},
			})
			return
		}

		maker, _ := paseto.NewPasetoMaker()
		token, _, err := maker.CreateToken(user.Username, time.Hour)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}

		err = uc.GenerateSession(ctx, user.Username, token, time.Hour)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Can't generate session"})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data": &domain.UserAccount{
				Id:          user.ID.String(),
				Username:    user.Username,
				Email:       user.Email.String,
				DisplayName: user.DisplayName.String,
				Role:        user.Role,
				Token:       token,
			},
			"status": "success",
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
			ctx.JSON(http.StatusOK, gin.H{"message": "User doesn't exist"})
			return
		}
		ctx.JSON(http.StatusOK, user)
	})

}

func changePassword(service users.UseCase) gin.HandlerFunc {
	// Add your code logic here
	return gin.HandlerFunc(func(ctx *gin.Context) {

		var req domain.UserLogin
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, nil)
			return
		}

		id, err := uuid.Parse(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}
		err = service.ChangePassword(ctx, id, req.Username, req.Password)
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{"message": "Can't change password"})

			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "Password changed successfully"})
	})

}

func deleteUser(service users.UseCase) gin.HandlerFunc {
	// Add your code logic here
	return gin.HandlerFunc(func(ctx *gin.Context) {

		id, err := uuid.Parse(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}
		err = service.DeleteUser(ctx, id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Can't delete user successfully. Error:" + err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
	})

}
