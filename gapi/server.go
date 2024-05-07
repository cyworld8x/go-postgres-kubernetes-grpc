package gapi

import (
	"github.com/cyworld8x/go-postgres-kubernetes-grpc/pkg/pb"

	db "github.com/cyworld8x/go-postgres-kubernetes-grpc/db/sqlc"
	"github.com/cyworld8x/go-postgres-kubernetes-grpc/util"
)

// Server serves HTTP requests for our banking service.
type Server struct {
	pb.UnimplementedUserServiceServer
	store db.Store
}

// NewServer creates a new gRPC server and set up routing.
func NewServer(config util.Configuration, store db.Store) (*Server, error) {

	server := &Server{
		store: store,
	}

	return server, nil
}
