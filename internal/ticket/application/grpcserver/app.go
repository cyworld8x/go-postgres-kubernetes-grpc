package grpcserver

import (
	"context"

	"github.com/cyworld8x/go-postgres-kubernetes-grpc/internal/ticket/infrastructure/repository"
	tickets "github.com/cyworld8x/go-postgres-kubernetes-grpc/internal/ticket/usecases/tickets"
	"github.com/cyworld8x/go-postgres-kubernetes-grpc/pkg/pb"
	"github.com/cyworld8x/go-postgres-kubernetes-grpc/pkg/postgres"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc"
)

type App struct {
	// Cfg               *config.Config
	UC         tickets.UseCase
	GRPCServer pb.TicketServiceServer
}

func New(
	// cfg *config.Config,
	uc tickets.UseCase,
	gRPCServer pb.TicketServiceServer,
) *App {
	return &App{
		// Cfg:               cfg,
		UC:         uc,
		GRPCServer: gRPCServer,
	}
}

func Init(
	dbConnStr postgres.DBConnString,
	grpcServer *grpc.Server,
) (*App, error) {

	connPool, _ := pgxpool.New(context.Background(), string(dbConnStr))
	repository := repository.NewRepository(connPool)
	uc := tickets.NewService(repository)
	server := NewServer(grpcServer, uc)
	app := New(uc, server)

	return app, nil

}
