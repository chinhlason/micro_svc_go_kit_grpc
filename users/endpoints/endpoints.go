package endpoints

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"users/data"
	"users/protobuf"
	"users/service"
)

type Endpoints struct {
	//endpoints local
	SyncEndpoint endpoint.Endpoint

	//endpoints from another server
	GetUserEndpoint endpoint.Endpoint
}

func NewEndpoints(is service.IdentityServices, us service.IUserServices) Endpoints {
	return Endpoints{
		SyncEndpoint: makeSyncEndpoint(is, us),
	}
}

func makeSyncEndpoint(is service.IdentityServices, us service.IUserServices) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*protobuf.SyncReq)
		user, err := is.GetUser(ctx, req.Username)
		if err != nil {
			return nil, err
		}
		syncReq := data.SyncUser{
			Id:       user.Id,
			Username: user.Username,
			Password: user.Password,
		}
		err = us.InsertUser(ctx, syncReq)
		if err != nil {
			return nil, err
		}
		return &protobuf.SyncRes{
			Message: "success",
		}, nil
	}
}

// GetUser func from another server
func (e Endpoints) GetUser(ctx context.Context, username string) (data.GetRes, error) {
	req := data.GetReq{
		Username: username,
	}
	resp, err := e.GetUserEndpoint(ctx, req)
	if err != nil {
		return data.GetRes{}, err
	}
	return resp.(data.GetRes), nil
}
