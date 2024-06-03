package repository

import (
	"context"
	"os"
	"testing"

	"github.com/rs/zerolog/log"

	"github.com/cyworld8x/go-postgres-kubernetes-grpc/pkg/utils"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
)

var testStore repository

func TestMain(m *testing.M) {
	config, err := utils.LoadConfiguration("../../../..")

	if err != nil {
		log.Fatal().Err(err).Msg("can not load env configuration.")
	}

	conn, err := pgxpool.New(context.Background(), config.DbSource)
	if err != nil {
		log.Fatal().Err(err).Msg("can not connect to db.")
	}

	testStore = NewRepository(conn)
	os.Exit(m.Run())

}
