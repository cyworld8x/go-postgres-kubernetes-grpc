package handler

import (
	"net/http"

	"github.com/cyworld8x/go-postgres-kubernetes-grpc/internal/crawler/domain"
	"github.com/cyworld8x/go-postgres-kubernetes-grpc/internal/crawler/usecases/sources"
	"github.com/cyworld8x/go-postgres-kubernetes-grpc/pkg/middleware"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func MakeSourceHandler(router gin.IRouter, sourcesUseCase sources.UseCase) {
	router.Use(middleware.GinLogger())
	router.POST("/sources", CreateSource(sourcesUseCase))
	router.PUT("/sources", UpdateSource(sourcesUseCase))
	router.GET("/sources/:id", GetSource(sourcesUseCase))
	router.GET("/sources/", GetSources(sourcesUseCase))
	// router.POST("/login", getLogin(service))
	// router.Use(middleware.GinLogger())
	// router.POST("/user", createUser(service))
	// router.POST("/login", getLogin(service))
	// pasetoMaker, _ := paseto.NewPasetoMaker()
	// authRoutes := router.Group("/").Use(middleware.AuthMiddleware(pasetoMaker))
	// authRoutes.GET("/user/:id", getUser(service))
}

// Get Source godoc
// @Summary      Create Source
// @Description  Create Source
// @Tags         Source
// @Accept       json
// @Produce      json
// @Param        arg  body domain.Source true "Source Info"
// @Success      200  {object}  domain.Source
// @Failure      400  {object}  error
// @Failure      404  {object}  error
// @Failure      500  {object}  error
// @Router       /sources [Post]
func CreateSource(usecase sources.UseCase) gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		var req domain.Source
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		source, errCreate := usecase.CreateSource(ctx, &req)

		if errCreate != nil {
			ctx.JSON(http.StatusInternalServerError, errCreate)
		}

		ctx.JSON(http.StatusOK, source)
	})
}

// Get Source godoc
// @Summary      get an source
// @Description  get source by Id
// @Tags         Source
// @Accept       json
// @Produce      json
// @Param        id   path string  true "Id"
// @Success      200  {object}  domain.Source
// @Failure      400  {object}  error
// @Failure      404  {object}  error
// @Failure      500  {object}  error
// @Router       /sources/{id} [get]
func GetSource(usecase sources.UseCase) gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		id, err := uuid.Parse(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}
		source, errGet := usecase.GetSource(ctx, id)

		if errGet != nil {
			ctx.JSON(http.StatusInternalServerError, errGet)
		}

		ctx.JSON(http.StatusOK, source)
	})
}

// Get Source godoc
// @Summary      get sources
// @Description  get source by Id
// @Tags         Source
// @Accept       json
// @Produce      json
// @Success      200  {object}  domain.Source
// @Failure      400  {object}  error
// @Failure      404  {object}  error
// @Failure      500  {object}  error
// @Router       /sources [get]
func GetSources(usecase sources.UseCase) gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		sources, errGet := usecase.GetSources(ctx)

		if errGet != nil {
			ctx.JSON(http.StatusInternalServerError, errGet)
		}

		ctx.JSON(http.StatusOK, sources)
	})
}

// Get Source godoc
// @Summary      Create Source
// @Description  Create Source
// @Tags         Source
// @Accept       json
// @Produce      json
// @Param        arg  body domain.Source true "Source Info"
// @Success      200  {object}  domain.Source
// @Failure      400  {object}  error
// @Failure      404  {object}  error
// @Failure      500  {object}  error
// @Router       /sources [Put]
func UpdateSource(usecase sources.UseCase) gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		var req domain.Source
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		source, errCreate := usecase.UpdateSource(ctx, &req)

		if errCreate != nil {
			ctx.JSON(http.StatusInternalServerError, errCreate)
		}

		ctx.JSON(http.StatusOK, source)
	})
}
