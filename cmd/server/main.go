package main

import (
	"context"
	"google.golang.org/grpc"
	"homework-2/internal/app"
	"homework-2/internal/db"
	"homework-2/internal/midware"
	"homework-2/internal/repository"
	pb "homework-2/pkg/api"
	"log"
	"net"
	"time"
)

func main() {

	ctx := context.Background()

	adp, err := db.New(ctx)
	if err != nil {
		log.Fatal(err)
	}

	newServer := app.New(repository.New(adp))
	lis, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	opts = []grpc.ServerOption{
		grpc.UnaryInterceptor(midware.LogInterceptor),
	}

	grpcServer := grpc.NewServer(opts...)
	pb.RegisterAwesomeBotIIIServer(grpcServer, newServer)
	err = grpcServer.Serve(lis)
	if err != nil {
		panic(err)
	}
	for {
		time.Sleep(time.Second)
	}
}
