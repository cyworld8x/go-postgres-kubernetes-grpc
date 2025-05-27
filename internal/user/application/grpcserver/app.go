package grpcserver

import (
	"context"

	"github.com/cyworld8x/go-postgres-kubernetes-grpc/internal/user/infrastructure/repository"
	"github.com/cyworld8x/go-postgres-kubernetes-grpc/internal/user/infrastructure/repository/dynamo"
	users "github.com/cyworld8x/go-postgres-kubernetes-grpc/internal/user/usecases/users"
	"github.com/cyworld8x/go-postgres-kubernetes-grpc/pkg/dynamodb"
	"github.com/cyworld8x/go-postgres-kubernetes-grpc/pkg/pb"
	"github.com/cyworld8x/go-postgres-kubernetes-grpc/pkg/postgres"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc"
)

type App struct {
	// Cfg               *config.Config
	UC         users.UseCase
	GRPCServer pb.UserServiceServer
}

func New(
	// cfg *config.Config,
	uc users.UseCase,
	gRPCServer pb.UserServiceServer,
) *App {
	return &App{
		// Cfg:               cfg,
		UC:         uc,
		GRPCServer: gRPCServer,
	}
}

func Init(
	dbConnStr postgres.DBConnString,
	dynamoEndPoint string,
	grpcServer *grpc.Server,
) (*App, error) {

	connPool, _ := pgxpool.New(context.Background(), string(dbConnStr))
	repository := repository.NewRepository(connPool)
	dbClient, err := dynamodb.NewDynamoDB(dynamoEndPoint)
	if err != nil {
		return nil, err
	}
	dynamoDB := dynamo.NewUserRepository(dbClient.GetDB())
	uc := users.NewService(repository, dynamoDB)
	server := NewServer(grpcServer, uc)
	app := New(uc, server)

	return app, nil

}

func dbEngineFunc(url postgres.DBConnString) (postgres.DBEngine, func(), error) {
	db, err := postgres.NewPostgresDB(url)
	if err != nil {
		return nil, nil, err
	}
	return db, func() { db.Close() }, nil
}
