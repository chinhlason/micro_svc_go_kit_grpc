package main

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	log2 "log"
	"net"
	"os"
	"users/common"
	"users/data"
	endpoints2 "users/endpoints"
	"users/protobuf"
	"users/service"
	"users/transport"
)

func Run() {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	db1, err := common.Connect("postgresql://postgres:postgres@localhost:5433/db2?sslmode=disable")
	if err != nil {
		panic(err)
	}

	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log2.Fatalf("fail to dial: %v", err)
	}
	op := common.CreateSerializableTxOption()
	ctx := context.WithValue(context.Background(), "db", "db2")
	ss := common.NewSession(db1, op, ctx)
	repo := data.NewRepository(ss)
	userSvc := service.NewUserService(repo)
	identitySvc := transport.NewGrpcClient(conn, logger)
	userEndpoints := endpoints2.NewEndpoints(identitySvc, userSvc)
	server := transport.NewGRPCServer(userEndpoints, logger)
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	protobuf.RegisterUserServiceServer(s, server)
	fmt.Println("server is listening on port 50052...")
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}

func main() {
	Run()
}
