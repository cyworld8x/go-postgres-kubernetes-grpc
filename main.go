package main

import (
	"context"
	"log"
	"social/util"

	"github.com/jackc/pgx/v5"
)

func main() {
	config, err := util.LoadConfiguration(".")
	if err != nil {
		log.Fatal("can not load env configuration:", err)
	}

	log.Printf("Load env configuration %s", config)
	_, err = pgx.Connect(context.Background(), config.DbSource)
	if err != nil {
		log.Fatal("can not connect to db:", err)
	}
}
