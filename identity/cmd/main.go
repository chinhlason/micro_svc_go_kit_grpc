package main

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/log"
	"google.golang.org/grpc"
	"identity/common"
	"identity/data"
	endpoints2 "identity/endpoints"
	"identity/protobuf"
	"identity/services"
	"identity/transport"
	log2 "log"
	"net"
	"os"
)

func Run() {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	db1, err := common.Connect("postgresql://postgres:postgres@localhost:5432/db1?sslmode=disable")
	if err != nil {
		panic(err)
	}
	op := common.CreateSerializableTxOption()
	ctx := context.WithValue(context.Background(), "db", "db1")
	ss := common.NewSession(db1, op, ctx)
	repo := data.NewRepository(ss)
	svc := services.NewService(repo)
	endpoints := endpoints2.NewEndpoints(svc)
	server := transport.NewGRPCServer(endpoints, logger)
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log2.Fatalf("fail to listen: %v", err)
	}
	s := grpc.NewServer()
	protobuf.RegisterIdentityServiceServer(s, server)
	fmt.Println("server is listening on port 50051...")
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}

func main() {
	Run()
}
