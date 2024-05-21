package main

import (
	"context"

	api "github.com/cyworld8x/go-postgres-kubernetes-grpc/internal/crawler/application/api"
	"github.com/cyworld8x/go-postgres-kubernetes-grpc/pkg/postgres"
	utils "github.com/cyworld8x/go-postgres-kubernetes-grpc/pkg/utils"
	"github.com/rs/zerolog/log"
	"go.uber.org/automaxprocs/maxprocs"
)

func main() {
	// set GOMAXPROCS
	_, err := maxprocs.Set()
	if err != nil {
		log.Err(err).Msg("failed set max procs")
	}

	ctx, cancel := context.WithCancel(context.Background())

	log.Info().Msg("⚡ init ticket api app")

	config, err := utils.LoadConfiguration(".")
	if err != nil {
		log.Fatal().Err(err).Msg("can not load env configuration")
	}

	log.Printf("Load env configuration %s", config)
	app, err := api.Init(postgres.DBConnString(config.DbSource))
	if err != nil {
		log.Fatal().Err(err).Msg("failed init app")
		cancel()
	}
	app.Server.Start(config.CrawlerAPIServerAddress)

	if err != nil {
		log.Fatal().Err(err).Msg("failed start ticket API")
		cancel()
		<-ctx.Done()
	}

	// 	// gRPC Server.
	// 	address := fmt.Sprintf("0.0.0.0:20242")
	// 	network := "tcp"

	// 	l, err := net.Listen(network, address)
	// 	if err != nil {
	// 		log.Err(err).Msg(fmt.Sprintf("failed to listen to address %s : %s ", network, address))
	// 		cancel()
	// 	}

}
