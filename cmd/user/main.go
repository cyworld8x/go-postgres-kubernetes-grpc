package main

import (
	"context"
	"net"

	"github.com/cyworld8x/go-postgres-kubernetes-grpc/internal/user/application/grpcserver"
	"github.com/cyworld8x/go-postgres-kubernetes-grpc/pkg/postgres"
	utils "github.com/cyworld8x/go-postgres-kubernetes-grpc/pkg/utils"
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

	log.Info().Msg("⚡ init user grpc app")

	// set up logrus

	server := grpc.NewServer()

	go func() {
		defer server.GracefulStop()
		<-ctx.Done()
	}()

	config, err := utils.LoadConfiguration(".")
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
		log.Fatal().Err(err).Msg("cannot create listener")
	}
	log.Printf("start gRPC server on %s", config.GRPCServerAddress)
	err = server.Serve(listener)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot start gRPC server")
	}

	// 	// gRPC Server.
	// 	address := fmt.Sprintf("0.0.0.0:20242")
	// 	network := "tcp"

	// 	l, err := net.Listen(network, address)
	// 	if err != nil {
	// 		log.Err(err).Msg(fmt.Sprintf("failed to listen to address %s : %s ", network, address))
	// 		cancel()
	// 	}

	// 	log.Log().Msg(fmt.Sprintf("🌏 start server... %s", address))

	// 	defer func() {
	// 		if err1 := l.Close(); err != nil {
	// 			log.Err(err).Msg(fmt.Sprintf("failed to close %s: %s :%s", network, address, err1))
	// 		}
	// 	}()

	// 	err = server.Serve(l)
	// 	if err != nil {
	// 		log.Err(err).Msg(fmt.Sprintf("failed start gRPC server  %s: %s :%s", network, address, err))
	// 		cancel()
	// 	}

	// 	quit := make(chan os.Signal, 1)
	// 	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	// 	select {
	// 	case v := <-quit:
	// 		log.Err(err).Msg(fmt.Sprintf("signal.Notify %s", v))
	// 	case done := <-ctx.Done():
	// 		log.Err(err).Msg(fmt.Sprintf("ctx.Done %s", done))
	// 	}
}
