package transport

import (
	"context"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	grpcTransport "github.com/go-kit/kit/transport/grpc"
	"users/endpoints"
	pb "users/protobuf"
)

type grpcServer struct {
	syncUser grpcTransport.Handler
	pb.UnimplementedUserServiceServer
}

func NewGRPCServer(endpoints endpoints.Endpoints, logger log.Logger) pb.UserServiceServer {
	options := []grpcTransport.ServerOption{
		grpcTransport.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
	}
	return &grpcServer{
		syncUser: grpcTransport.NewServer(
			endpoints.SyncEndpoint,
			decodeGRPCSyncUserRequest,
			encodeGRPCSyncUserResponse,
			options...,
		),
	}
}

func (s *grpcServer) SyncUser(ctx context.Context, req *pb.SyncReq) (*pb.SyncRes, error) {
	_, resp, err := s.syncUser.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.SyncRes), nil
}

func decodeGRPCSyncUserRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.SyncReq)
	return req, nil
}

func encodeGRPCSyncUserResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.SyncRes)
	return resp, nil
}
