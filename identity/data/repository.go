package data

import (
	"context"
	"fmt"
	"identity/common"
)

type IRepository interface {
	Insert(context.Context, string, string) error
	Select(context.Context, string) (User, error)
}

type Repository struct {
	ss *common.Session
}

func NewRepository(ss *common.Session) *Repository {
	return &Repository{ss: ss}
}

func (r *Repository) Insert(ctx context.Context, username, password string) error {
	_, err := r.ss.ExecQuery(ctx, "INSERT INTO db1 (username, password) VALUES ($1, $2)", username, password)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) Select(ctx context.Context, username string) (User, error) {
	rows, err := r.ss.QueryMultiRows(ctx, "SELECT * FROM db1 WHERE username = $1", username)
	if err != nil {
		return User{}, err
	}
	defer rows.Close()
	var user []User
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.Id, &u.Username, &u.Password); err != nil {
			return User{}, err
		}
		user = append(user, u)
	}
	fmt.Println("user: ", user)
	return user[0], nil
}
