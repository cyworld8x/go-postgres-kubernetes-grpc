package api

import (
	"context"

	"github.com/cyworld8x/go-postgres-kubernetes-grpc/internal/user/infrastructure/repository"
	"github.com/cyworld8x/go-postgres-kubernetes-grpc/internal/user/infrastructure/repository/dynamo"
	"github.com/cyworld8x/go-postgres-kubernetes-grpc/internal/user/usecases/sessions"
	"github.com/cyworld8x/go-postgres-kubernetes-grpc/internal/user/usecases/users"
	"github.com/cyworld8x/go-postgres-kubernetes-grpc/pkg/dynamodb"
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
	dynamoEndPoint string,
) (*App, error) {

	connPool, _ := pgxpool.New(context.Background(), string(dbConnStr))
	repository := repository.NewRepository(connPool)
	dbClient, err := dynamodb.NewDynamoDB(dynamoEndPoint)
	if err != nil {
		return nil, err
	}

	dynamoDB := dynamo.UserDynamoDBRepository(dbClient.GetDB())
	sessionRepo := dynamo.SessionRepository(dbClient.GetDB())

	uc := users.NewService(repository, dynamoDB, sessionRepo)
	sessionUC := sessions.NewService(sessionRepo)
	apiServer := NewServer(uc, sessionUC)
	app := New(apiServer)
	return app, nil

}
