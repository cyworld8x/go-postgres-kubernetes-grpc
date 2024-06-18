package handler

import (
	"net/http"

	"github.com/cyworld8x/go-postgres-kubernetes-grpc/internal/crawler/usecases/prompt"
	"github.com/cyworld8x/go-postgres-kubernetes-grpc/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func MakePromptHandler(router gin.IRouter, usecase prompt.UseCase) {
	router.Use(middleware.GinLogger())
	router.GET("/prompt", SinglePrompt(usecase))
	// authRoutes.GET("/user/:id", getUser(service))
}

// Get Prompt godoc
// @Summary      exec a prompt
// @Description  exec a prompt
// @Tags         Prompt
// @Accept       json
// @Produce      json
// @Param        message query string false "message"
// @Success      200  {object}  string
// @Failure      400  {object}  error
// @Failure      404  {object}  error
// @Failure      500  {object}  error
// @Router       /prompt [get]
func SinglePrompt(usecase prompt.UseCase) gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		prompt := ctx.Query("message")
		source, errGet := usecase.SinglePrompt(ctx, prompt)

		if errGet != nil {
			ctx.JSON(http.StatusInternalServerError, errGet)
		}

		ctx.JSON(http.StatusOK, source)
	})
}
