package api

import (
	"context"

	"github.com/cyworld8x/go-postgres-kubernetes-grpc/internal/event/infrastructure/repository"
	events "github.com/cyworld8x/go-postgres-kubernetes-grpc/internal/event/usecases/events"
	"github.com/cyworld8x/go-postgres-kubernetes-grpc/pkg/postgres"
	"github.com/jackc/pgx/v5/pgxpool"
)

type App struct {
	Server *Server
}

func New(
	// cfg *config.Config,

	Server *Server,
) *App {
	return &App{
		// Cfg:               cfg,
		Server: Server,
	}
}

func Init(
	dbConnStr postgres.DBConnString,
) (*App, error) {

	connPool, _ := pgxpool.New(context.Background(), string(dbConnStr))
	repository := repository.NewRepository(connPool)
	uc := events.NewService(repository)
	apiServer := NewServer(uc)
	app := New(apiServer)
	return app, nil

}
