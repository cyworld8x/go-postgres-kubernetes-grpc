package api

import (
	"github.com/cyworld8x/go-postgres-kubernetes-grpc/internal/event/application/api/handler"
	"github.com/cyworld8x/go-postgres-kubernetes-grpc/internal/event/application/api/swagger/docs"
	"github.com/cyworld8x/go-postgres-kubernetes-grpc/internal/event/usecases/events"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Server serves HTTP requests for our banking service.
type Server struct {
	uc     events.UseCase
	Router *gin.Engine
}

// NewServer creates a new gRPC server and set up routing.
// @title           Swagger Event API
// @version         1.0
// @description     This is a Event API server.
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
func NewServer(uc events.UseCase) *Server {

	server := &Server{
		uc: uc,
	}

	router := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"
	routerGroup := router.Group(docs.SwaggerInfo.BasePath)
	handler.MakeEventHandler(routerGroup, uc)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	server.Router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.Router.Run(address)
}
