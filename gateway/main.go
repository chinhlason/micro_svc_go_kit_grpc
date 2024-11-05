package main

import (
	"context"
	"flag"
	"fmt"
	protobuf "gateway/protobuf"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/grpclog"
)

var (
	// command-line options:
	// gRPC server endpoint
	grpcIdentityServerEndpoint = flag.String("grpc-identity-server-endpoint", "localhost:50051", "gRPC identity server endpoint")
	grpcUserServerEndpoint     = flag.String("grpc-user-service-endpoint", "localhost:50052", "gRPC user service endpoint")
)

func startGateway() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	//err := protobuf.RegisterIdentityServiceHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts)
	err := protobuf.RegisterIdentityServiceHandlerFromEndpoint(ctx, mux, *grpcIdentityServerEndpoint, opts)
	if err != nil {
		return err
	}

	err = protobuf.RegisterUserServiceHandlerFromEndpoint(ctx, mux, *grpcUserServerEndpoint, opts)
	if err != nil {
		return err
	}
	fmt.Println("Starting HTTP server gateway at port 8089...")
	// Start HTTP server (and proxy calls to gRPC server endpoint)
	return http.ListenAndServe(":8089", mux)
}

func main() {
	flag.Parse()
	if err := startGateway(); err != nil {
		grpclog.Fatal(err)
	}
}
