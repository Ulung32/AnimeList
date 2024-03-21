package repository

import (
	"AnimeList/model"
	"context"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user model.User) error
}

func (r Repository) CreateUser(ctx context.Context, user model.User) error {
	err := r.DB.Create(&user)

	if err != nil {
		return err.Error
	}

	return nil
}
