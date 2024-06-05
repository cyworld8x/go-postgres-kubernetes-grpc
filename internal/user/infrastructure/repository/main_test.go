package repository

import (
	"context"
	"os"
	"testing"

	"github.com/rs/zerolog/log"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
)

var testStore repository

func TestMain(m *testing.M) {
	dbSource := "postgresql://postgres:postgres@localhost:20241/socialdb?sslmode=disable"

	conn, err := pgxpool.New(context.Background(), dbSource)
	if err != nil {
		log.Fatal().Err(err).Msg("can not connect to db.")
	}

	testStore = NewRepository(conn)
	os.Exit(m.Run())

}
