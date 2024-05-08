package grpcserver

import (
	"github.com/cyworld8x/go-postgres-kubernetes-grpc/internal/user/usecases/users"
	"github.com/cyworld8x/go-postgres-kubernetes-grpc/pkg/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Server serves HTTP requests for our banking service.
type Server struct {
	pb.UnimplementedUserServiceServer
	uc users.UseCase
}

// NewServer creates a new gRPC server and set up routing.
func NewServer(grpcServer *grpc.Server, uc users.UseCase) pb.UserServiceServer {

	server := &Server{
		uc: uc,
	}

	pb.RegisterUserServiceServer(grpcServer, server)

	reflection.Register(grpcServer)

	return server
}
