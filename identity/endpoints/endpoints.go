package endpoints

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	pb "identity/protobuf"
	"identity/services"
)

type Endpoints struct {
	InsertUserEndpoint endpoint.Endpoint
	SelectUserEndpoint endpoint.Endpoint
}

func NewEndpoints(s services.Service) Endpoints {
	return Endpoints{
		InsertUserEndpoint: makeInsertUserEndpoint(s),
		SelectUserEndpoint: makeSelectUserEndpoint(s),
	}
}

func makeInsertUserEndpoint(s services.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*pb.InsertReq)
		res, err := s.InsertUser(ctx, req)
		if err != nil {
			return nil, err
		}
		return res, nil
	}
}

func makeSelectUserEndpoint(s services.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*pb.GetReq)
		res, err := s.GetUser(ctx, req)
		if err != nil {
			return nil, err
		}
		return res, nil
	}
}
