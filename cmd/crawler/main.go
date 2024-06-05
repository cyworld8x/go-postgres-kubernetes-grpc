package main

import (
	"context"

	configuration "github.com/cyworld8x/go-postgres-kubernetes-grpc/cmd/crawler/config"
	"github.com/cyworld8x/go-postgres-kubernetes-grpc/internal/crawler/application/api"
	"github.com/cyworld8x/go-postgres-kubernetes-grpc/pkg/postgres"
	"github.com/playwright-community/playwright-go"
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

	config, err := configuration.Load()
	if err != nil {
		log.Fatal().Err(err).Msg("can not load env configuration")
	}

	log.Printf("Load env configuration %s", config)
	pw, err := playwright.Run()
	if err != nil {
		log.Fatal().Err(err).Msgf("could not start playwright: %v", err)
	}
	app, err := api.Init(postgres.DBConnString(config.DbSource), pw)
	app.Server.Start(config.HttpServerAddress)

	if err != nil {
		log.Fatal().Err(err).Msg("failed start ticket API")
		pw.Stop()
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
