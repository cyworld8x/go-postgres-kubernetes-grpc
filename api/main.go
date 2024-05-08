package api

import (
	handler "github.com/cyworld8x/go-postgres-kubernetes-grpc/api/handler"

	db "github.com/cyworld8x/go-postgres-kubernetes-grpc/db/sqlc"
	"github.com/cyworld8x/go-postgres-kubernetes-grpc/pkg/utils"
	"github.com/cyworld8x/go-postgres-kubernetes-grpc/usecase/user"
	"github.com/gin-gonic/gin"
)

// Server serves HTTP requests for our banking service.
type Server struct {
	store  db.Store
	Router *gin.Engine
}

// NewServer creates a new HTTP server and set up routing.
func NewServer(config utils.Configuration, store db.Store) (*Server, error) {

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

	// router.POST("/users/login", server.loginUser)
	// router.POST("/tokens/renew_access", server.renewAccessToken)

	// authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))
	// authRoutes.POST("/accounts", server.createAccount)
	// authRoutes.GET("/accounts/:id", server.getAccount)
	// authRoutes.GET("/accounts", server.listAccounts)

	// authRoutes.POST("/transfers", server.createTransfer)

	server.Router = router
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.Router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
