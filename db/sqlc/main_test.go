package social

import (
	"context"
	"log"
	"os"
	"social/util"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
)

var testStore Store

func TestMain(m *testing.M) {
	config, err := util.LoadConfiguration("../..")

	if err != nil {
		log.Fatal("can not load env configuration:", err)
	}

	conn, err := pgxpool.New(context.Background(), config.DbSource)
	if err != nil {
		log.Fatal("can not connect to db:", err)
	}

	testStore = NewStore(conn)
	os.Exit(m.Run())

}
