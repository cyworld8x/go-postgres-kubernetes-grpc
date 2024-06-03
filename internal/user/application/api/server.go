package api

import (
	"github.com/cyworld8x/go-postgres-kubernetes-grpc/internal/user/application/api/handler"
	"github.com/cyworld8x/go-postgres-kubernetes-grpc/internal/user/usecases/users"
	"github.com/gin-gonic/gin"
)

// Server serves HTTP requests for our banking service.
type Server struct {
	uc     users.UseCase
	Router *gin.Engine
}

func NewServer(uc users.UseCase) *Server {

	server := &Server{
		uc: uc,
	}

	router := gin.Default()
	handler.MakeUserHandler(router, uc)
	server.Router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.Router.Run(address)
}
