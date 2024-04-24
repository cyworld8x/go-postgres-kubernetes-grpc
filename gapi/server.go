package gapi

import (
	handler "github.com/cyworld8x/go-postgres-kubernetes-grpc/api/handler"
	"github.com/cyworld8x/go-postgres-kubernetes-grpc/pkg/pb"

	db "github.com/cyworld8x/go-postgres-kubernetes-grpc/db/sqlc"
	"github.com/cyworld8x/go-postgres-kubernetes-grpc/usecase/user"
	"github.com/cyworld8x/go-postgres-kubernetes-grpc/util"
	"github.com/gin-gonic/gin"
)

// Server serves HTTP requests for our banking service.
type Server struct {
	pb.UnimplementedUserServiceServer
	store  db.Store
	Router *gin.Engine
}

// NewServer creates a new gRPC server and set up routing.
func NewServer(config util.Configuration, store db.Store) (*Server, error) {

	server := &Server{
		store: store,
	}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()
	userService := user.NewService(server.store)
	handler.MakeUserHandler(router, userService) // Pass userService instead of userService.CreateUser

	server.Router = router
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.Router.Run(address)
}
