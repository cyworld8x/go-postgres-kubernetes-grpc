package handler

import (
	"net/http"

	"github.com/cyworld8x/go-postgres-kubernetes-grpc/internal/crawler/domain"
	crawler "github.com/cyworld8x/go-postgres-kubernetes-grpc/internal/crawler/usecases/crawler"
	"github.com/cyworld8x/go-postgres-kubernetes-grpc/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func MakeCrawlerHandler(router gin.IRouter, crawlerUseCase crawler.UseCase) {
	router.Use(middleware.GinLogger())
	router.POST("/crawl", Crawl(crawlerUseCase))
}

// Get Crawler godoc
// @Summary      Crawl website
// @Description  Crawl website
// @Tags         Crawler
// @Accept       json
// @Produce      json
// @Param        arg  body domain.WebSite true "Website Info"
// @Success      200  {object}  domain.Entry
// @Failure      400  {object}  error
// @Failure      404  {object}  error
// @Failure      500  {object}  error
// @Router       /crawl [Post]
func Crawl(usecase crawler.UseCase) gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		var req domain.WebSite
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		source, errCreate := usecase.Get(ctx, &req)

		if errCreate != nil {
			ctx.JSON(http.StatusInternalServerError, errCreate)
		}

		ctx.JSON(http.StatusOK, source)
	})
}
