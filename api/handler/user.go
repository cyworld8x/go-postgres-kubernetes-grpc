package api

import (
	"net/http"
	"strconv"

	presenter "github.com/cyworld8x/go-postgres-kubernetes-grpc/api/presenter"
	db "github.com/cyworld8x/go-postgres-kubernetes-grpc/db/sqlc"
	user "github.com/cyworld8x/go-postgres-kubernetes-grpc/usecase/user"
	"github.com/cyworld8x/go-postgres-kubernetes-grpc/util"

	"github.com/gin-gonic/gin"
)

type createUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Fullname string `json:"fullname"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func newUserResponse(user db.User) presenter.User {
	return presenter.User{
		ID:        user.ID,
		Username:  user.Username.String,
		Email:     user.Email.String,
		Fullname:  user.Fullname.String,
		Password:  user.Password.String,
		Role:      user.Role.String,
		CreatedAt: user.CreatedAt.Time,
	}
}

func MakeUserHandler(router *gin.Engine, service user.UseCase) {
	router.POST("/user", createUser(service))
	router.GET("/user/:id", getUser(service))
	router.POST("/login", getLogin(service))
}
func createUser(service user.UseCase) gin.HandlerFunc {
	// Add your code logic here
	return gin.HandlerFunc(func(ctx *gin.Context) {

		var req createUserRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}
		user, err := service.CreateUser(req.Username, req.Email, req.Fullname, req.Password, req.Role)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
		rsp := newUserResponse(user)
		ctx.JSON(http.StatusOK, rsp)
	})
}

func getLogin(service user.UseCase) gin.HandlerFunc {
	// Add your code logic here
	return gin.HandlerFunc(func(ctx *gin.Context) {
		var req loginRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, nil)
			return
		}

		user, err := service.GetLogin(req.Username)
		if err != nil {
			ctx.JSON(http.StatusOK, "User name doesn't match")
			return
		}

		err = util.CheckPassword(req.Password, user.Password.String)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, "Can't login with user name and password.")
			return
		}

		rsp := newUserResponse(user)
		ctx.JSON(http.StatusOK, rsp)
	})

}

func getUser(service user.UseCase) gin.HandlerFunc {
	// Add your code logic here
	return gin.HandlerFunc(func(ctx *gin.Context) {

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}
		user, err := service.GetUser(int32(id))
		if err != nil {
			ctx.JSON(http.StatusOK, "User doesn't exist")
			return
		}
		rsp := newUserResponse(user)
		ctx.JSON(http.StatusOK, rsp)
	})

}
