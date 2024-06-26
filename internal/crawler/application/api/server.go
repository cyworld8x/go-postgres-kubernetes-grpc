package api

import (
	"github.com/cyworld8x/go-postgres-kubernetes-grpc/internal/crawler/application/api/handler"
	docs "github.com/cyworld8x/go-postgres-kubernetes-grpc/internal/crawler/application/api/swagger/docs"
	"github.com/cyworld8x/go-postgres-kubernetes-grpc/internal/crawler/usecases/crawler"
	"github.com/cyworld8x/go-postgres-kubernetes-grpc/internal/crawler/usecases/prompt"
	"github.com/cyworld8x/go-postgres-kubernetes-grpc/internal/crawler/usecases/sources"
	"github.com/cyworld8x/go-postgres-kubernetes-grpc/pkg/metric"
	"github.com/cyworld8x/go-postgres-kubernetes-grpc/pkg/middleware"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Server serves HTTP requests for our banking service.
type Server struct {
	sourceUC      sources.UseCase
	crawlerUC     crawler.UseCase
	promptUC      prompt.UseCase
	metricService metric.MetricService
	Router        *gin.Engine
}

// NewServer creates a new gRPC server and set up routing.
// @title           Swagger Crawler API
// @version         1.0
// @description     This is a Crawler API server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func NewServer(sourceUC sources.UseCase, crawlerUC crawler.UseCase, promptUC prompt.UseCase, metricService metric.MetricService) *Server {

	server := &Server{
		sourceUC:  sourceUC,
		crawlerUC: crawlerUC,
		promptUC:  promptUC,
	}

	router := gin.Default()
	router.Handle("GET", "/metrics", gin.WrapH(promhttp.Handler()))
	docs.SwaggerInfo.BasePath = "/api/v1"
	routerGroup := router.Group(docs.SwaggerInfo.BasePath)
	routerGroup.Use(middleware.Metrics(metricService))
	handler.MakeSourceHandler(routerGroup, sourceUC)
	handler.MakeCrawlerHandler(routerGroup, crawlerUC)
	handler.MakePromptHandler(routerGroup, promptUC)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	server.Router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.Router.Run(address)
}
