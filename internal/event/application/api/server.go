package api

import (
	"github.com/cyworld8x/go-postgres-kubernetes-grpc/internal/event/application/api/handler"
	"github.com/cyworld8x/go-postgres-kubernetes-grpc/internal/event/usecases/events"
	"github.com/gin-gonic/gin"
)

// Server serves HTTP requests for our banking service.
type Server struct {
	uc     events.UseCase
	Router *gin.Engine
}

// NewServer creates a new gRPC server and set up routing.
func NewServer(uc events.UseCase) *Server {

	server := &Server{
		uc: uc,
	}

	router := gin.Default()
	handler.MakeEventHandler(router, uc)
	server.Router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.Router.Run(address)
}
