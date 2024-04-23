package servertest

import (
	"context"
	"testing"

	api "github.com/cyworld8x/go-postgres-kubernetes-grpc/api"
	"github.com/cyworld8x/go-postgres-kubernetes-grpc/util"
	"github.com/jackc/pgx/v5/pgxpool"

	db "github.com/cyworld8x/go-postgres-kubernetes-grpc/db/sqlc"
)

func NewTestServer(t *testing.T) (*api.Server, error) {
	testConfig := util.Configuration{
		DbSource:          "postgresql://postgres:postgres@localhost:20241/socialdb?sslmode=disable",
		HTTPServerAddress: "localhost:8080",
	}
	conn, err := pgxpool.New(context.Background(), testConfig.DbSource)
	testStore := db.NewStore(conn)
	if err == nil {
		return api.NewServer(testConfig, testStore)
	}

	return nil, err
}
