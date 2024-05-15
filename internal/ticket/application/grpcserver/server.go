package grpcserver

import (
	"github.com/cyworld8x/go-postgres-kubernetes-grpc/internal/ticket/usecases/tickets"
	"github.com/cyworld8x/go-postgres-kubernetes-grpc/pkg/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Server serves HTTP requests for our banking service.
type Server struct {
	pb.UnimplementedTicketServiceServer
	uc tickets.UseCase
}

// NewServer creates a new gRPC server and set up routing.
func NewServer(grpcServer *grpc.Server, uc tickets.UseCase) pb.TicketServiceServer {

	server := &Server{
		uc: uc,
	}

	pb.RegisterTicketServiceServer(grpcServer, server)

	reflection.Register(grpcServer)

	return server
}
