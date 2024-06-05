package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	configuration "github.com/cyworld8x/go-postgres-kubernetes-grpc/cmd/ticket/config"
	"github.com/cyworld8x/go-postgres-kubernetes-grpc/internal/ticket/application/grpcserver"
	"github.com/cyworld8x/go-postgres-kubernetes-grpc/pkg/postgres"
	"github.com/rs/zerolog/log"
	"go.uber.org/automaxprocs/maxprocs"
	"google.golang.org/grpc"
)

func main() {
	// set GOMAXPROCS
	_, err := maxprocs.Set()
	if err != nil {
		log.Err(err).Msg("failed set max procs")
	}

	ctx, cancel := context.WithCancel(context.Background())

	log.Info().Msg("⚡ init ticket grpc app")

	// set up logrus

	server := grpc.NewServer()

	go func() {
		defer server.GracefulStop()
		<-ctx.Done()
	}()

	config, err := configuration.Load()
	if err != nil {
		log.Fatal().Err(err).Msg("can not load env configuration")
	}

	log.Printf("Load env configuration %s", config)
	_, err = grpcserver.Init(postgres.DBConnString(config.DbSource), server)
	if err != nil {
		log.Fatal().Err(err).Msg("failed init app")
		cancel()
	}
	//gRPCServerAddress := "0.0.0.0:20242"
	listener, err := net.Listen("tcp", config.GRPCServerAddress)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create listener: " + config.GRPCServerAddress)
		cancel()
	}
	log.Printf("start gRPC server on %s", config.GRPCServerAddress)
	err = server.Serve(listener)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot start gRPC server")
		cancel()
	}

	defer func() {
		if err1 := listener.Close(); err != nil {
			log.Err(err).Msg(fmt.Sprintf("failed to close %s :%s", config.GRPCServerAddress, err1))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	select {
	case v := <-quit:
		log.Err(err).Msg(fmt.Sprintf("signal.Notify %s", v))
	case done := <-ctx.Done():
		log.Err(err).Msg(fmt.Sprintf("ctx.Done %s", done))
	}
}
