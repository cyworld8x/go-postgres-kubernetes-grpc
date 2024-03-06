package social

import (
	"context"
	"log"
	"os"
	"social/util"
	"testing"

	"github.com/jackc/pgx/v5"
	_ "github.com/lib/pq"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	config, err := util.LoadConfiguration("../..")

	if err != nil {
		log.Fatal("can not load env configuration:", err)
	}

	conn, err := pgx.Connect(context.Background(), config.DbSource)
	if err != nil {
		log.Fatal("can not connect to db:", err)
	}

	testQueries = New(conn)
	os.Exit(m.Run())

}
