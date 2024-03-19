package main

import (
	"context"
	"log"

	"github.com/cyworld8x/go-postgres-kubernetes-grpc/api"
	db "github.com/cyworld8x/go-postgres-kubernetes-grpc/db/sqlc"
	"github.com/cyworld8x/go-postgres-kubernetes-grpc/util"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	config, err := util.LoadConfiguration(".")
	if err != nil {
		log.Fatal("can not load env configuration:", err)
	}

	log.Printf("Load env configuration %s", config)
	conn, err := pgxpool.New(context.Background(), config.DbSource)
	store := db.NewStore(conn)
	if err != nil {
		log.Fatal("can not connect to db:", err)
	}

	runGinServer(config, store)
}

func runGinServer(config util.Configuration, store db.Store) {
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server", err)
	}

	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatal("cannot start server", err)
	}
}
