package main

import (
	"context"
	"log"
	"net"

	"github.com/cyworld8x/go-postgres-kubernetes-grpc/api"
	db "github.com/cyworld8x/go-postgres-kubernetes-grpc/db/sqlc"
	"github.com/cyworld8x/go-postgres-kubernetes-grpc/gapi"
	"github.com/cyworld8x/go-postgres-kubernetes-grpc/pkg/pb"
	"github.com/cyworld8x/go-postgres-kubernetes-grpc/util"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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
	go runGinServer(config, store)
	runRPCServer(config, store)
}

func runRPCServer(config util.Configuration, store db.Store) {
	server, err := gapi.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server", err)
	}
	grpcLogger := grpc.UnaryInterceptor(gapi.GrpcLogger)

	grpcServer := grpc.NewServer(grpcLogger)
	pb.RegisterUserServiceServer(grpcServer, server)
	reflection.Register(grpcServer)
	listener, err := net.Listen("tcp", config.GRPCServerAddress)
	if err != nil {
		log.Fatal("cannot create listener", err)
	}
	log.Printf("start gRPC server on %s", config.GRPCServerAddress)
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start gRPC server", err)
	}

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
