package transport

import (
	"context"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	grpcTransport "github.com/go-kit/kit/transport/grpc"
	"identity/endpoints"
	pb "identity/protobuf"
)

type grpcServer struct {
	insertUser grpcTransport.Handler
	selectUser grpcTransport.Handler
	pb.UnimplementedIdentityServiceServer
}

func NewGRPCServer(endpoints endpoints.Endpoints, logger log.Logger) pb.IdentityServiceServer {
	options := []grpcTransport.ServerOption{
		grpcTransport.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
	}
	return &grpcServer{
		insertUser: grpcTransport.NewServer(
			endpoints.InsertUserEndpoint,
			decodeGRPCInsertUserRequest,
			encodeGRPCInsertUserResponse,
			options...,
		),
		selectUser: grpcTransport.NewServer(
			endpoints.SelectUserEndpoint,
			decodeGRPCGetUserRequest,
			encodeGRPCGetUserResponse,
			options...,
		),
	}
}

func (s *grpcServer) InsertUser(ctx context.Context, req *pb.InsertReq) (*pb.InsertRes, error) {
	_, resp, err := s.insertUser.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.InsertRes), nil
}

func (s *grpcServer) GetUser(ctx context.Context, req *pb.GetReq) (*pb.GetRes, error) {
	_, resp, err := s.selectUser.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.GetRes), nil
}

func decodeGRPCInsertUserRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.InsertReq)
	return req, nil
}

func encodeGRPCInsertUserResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.InsertRes)
	return resp, nil
}

func decodeGRPCGetUserRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.GetReq)
	return req, nil
}

func encodeGRPCGetUserResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.GetRes)
	return resp, nil
}
