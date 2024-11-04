package services

import (
	"context"
	"identity/data"
	pb "identity/protobuf"
)

type Service interface {
	InsertUser(context.Context, *pb.InsertReq) (*pb.InsertRes, error)
	GetUser(context.Context, *pb.GetReq) (*pb.GetRes, error)
}

type service struct {
	db data.IRepository
	pb.UnimplementedIdentityServiceServer
}

func NewService(db *data.Repository) Service {
	return &service{
		db: db,
	}
}

func (s *service) InsertUser(ctx context.Context, req *pb.InsertReq) (*pb.InsertRes, error) {
	err := s.db.Insert(ctx, req.Username, req.Password)
	if err != nil {
		return nil, err
	}
	return &pb.InsertRes{
		Token: "token is generated",
	}, nil
}

func (s *service) GetUser(ctx context.Context, req *pb.GetReq) (*pb.GetRes, error) {
	user, err := s.db.Select(ctx, req.Username)
	if err != nil {
		return nil, err
	}
	return &pb.GetRes{
		Id:       int64(user.Id),
		Username: user.Username,
		Password: user.Password,
	}, nil
}
