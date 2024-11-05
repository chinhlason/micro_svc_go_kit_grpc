package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	protobuf "users/protobuf"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/grpclog"
)

var (
	// command-line options:
	// gRPC server endpoint
	grpcServerEndpoint = flag.String("grpc-server-endpoint", "localhost:50052", "gRPC server endpoint")
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	//err := protobuf.RegisterIdentityServiceHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts)
	err := protobuf.RegisterUserServiceHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts)
	if err != nil {
		return err
	}
	fmt.Println("Starting HTTP server gateway at port 8088...")
	// Start HTTP server (and proxy calls to gRPC server endpoint)
	return http.ListenAndServe(":8088", mux)
}

func main() {
	flag.Parse()
	if err := run(); err != nil {
		grpclog.Fatal(err)
	}
}
