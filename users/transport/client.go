package transport

import (
	"context"
	"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/ratelimit"
	grpcTransport "github.com/go-kit/kit/transport/grpc"
	"github.com/sony/gobreaker"
	"golang.org/x/time/rate"
	"google.golang.org/grpc"
	"time"
	"users/data"
	"users/endpoints"
	"users/protobuf"
	"users/service"
)

func NewGrpcClient(conn *grpc.ClientConn, logger log.Logger) service.IdentityServices {
	limiter := ratelimit.NewErroringLimiter(rate.NewLimiter(rate.Every(time.Second), 100))
	var options []grpcTransport.ClientOption
	var getUserEndpoint endpoint.Endpoint
	{
		getUserEndpoint = grpcTransport.NewClient(
			conn,
			"pb.IdentityService",
			"GetUser",
			encodeRequest,
			decodeResponse,
			protobuf.GetRes{},
			options...,
		).Endpoint()
		getUserEndpoint = limiter(getUserEndpoint)
		getUserEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
			Name:    "GetUser",
			Timeout: 30 * time.Second,
		}))(getUserEndpoint)
	}

	return endpoints.Endpoints{
		GetUserEndpoint: getUserEndpoint,
	}
}

func encodeRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(data.GetReq)
	return &protobuf.GetReq{
		Username: req.Username,
	}, nil
}

func decodeResponse(_ context.Context, grpcResponse interface{}) (interface{}, error) {
	response := grpcResponse.(*protobuf.GetRes)
	return data.GetRes{
		Id:       response.Id,
		Username: response.Username,
		Password: response.Password,
	}, nil
}
