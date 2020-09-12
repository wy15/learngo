package main

import (
	"context"
	"flag"
	"log"
	"net/http"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	gw "github.com/wy15/learngo/build/gen/helloworld"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	grpcServerEndpoint = "127.0.0.1:50051"
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	creds, err := credentials.NewClientTLSFromFile(
		"tls/rootCA.pem", "",
	)
	if err != nil {
		log.Fatal(err)
	}
	opts := []grpc.DialOption{grpc.WithTransportCredentials(creds)}
	err = gw.RegisterGreeterHandlerFromEndpoint(ctx, mux, grpcServerEndpoint, opts)
	if err != nil {
		return err
	}

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	return http.ListenAndServe(":8081", mux)
}

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}
