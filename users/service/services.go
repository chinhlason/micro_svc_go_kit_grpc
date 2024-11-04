package service

import (
	"context"
	"users/data"
)

type IdentityServices interface {
	GetUser(context.Context, string) (data.GetRes, error)
}

type IUserServices interface {
	InsertUser(context.Context, data.SyncUser) error
}

type userService struct {
	db data.IRepository
}

func NewUserService(db *data.Repository) IUserServices {
	return &userService{
		db: db,
	}
}

func (us *userService) InsertUser(ctx context.Context, user data.SyncUser) error {
	err := us.db.Insert(ctx, user.Username, user.Password)
	if err != nil {
		return err
	}
	return nil
}
