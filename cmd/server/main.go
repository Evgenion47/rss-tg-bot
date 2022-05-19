package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"homework-2/internal/app"
	"homework-2/internal/db"
	"homework-2/internal/midware"
	"homework-2/internal/repository"
	pb "homework-2/pkg/api"
	"log"
	"net"
	"net/http"
	"time"
)

func runRest() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := pb.RegisterAwesomeBotIIIHandlerFromEndpoint(ctx, mux, "localhost:8080", opts)
	if err != nil {
		panic(err)
	}
	log.Printf("server listening at 8081")
	if err := http.ListenAndServe(":8081", mux); err != nil {
		panic(err)
	}
}

func main() {

	go runRest()

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
