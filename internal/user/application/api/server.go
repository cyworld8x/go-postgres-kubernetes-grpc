package api

import (
	"github.com/cyworld8x/go-postgres-kubernetes-grpc/internal/user/application/api/handler"
	"github.com/cyworld8x/go-postgres-kubernetes-grpc/internal/user/usecases/sessions"
	"github.com/cyworld8x/go-postgres-kubernetes-grpc/internal/user/usecases/users"
	"github.com/gin-gonic/gin"
)

// Server serves HTTP requests for our banking service.
type Server struct {
	uc        users.UseCase
	sessionUC sessions.UseCase
	Router    *gin.Engine
}

func NewServer(uc users.UseCase, sessionUC sessions.UseCase) *Server {

	server := &Server{
		uc:        uc,
		sessionUC: sessionUC,
	}

	router := gin.Default()
	handler.MakeUserHandler(router, uc)
	handler.MakeSessionHandler(router, sessionUC)
	server.Router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.Router.Run(address)
}
