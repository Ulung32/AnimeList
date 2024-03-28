package repository

import (
	"AnimeList/model"
	"context"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user model.User) error
	GetUser(ctx context.Context, email string) (model.User, error)
}

func (r Repository) CreateUser(ctx context.Context, user model.User) error {
	err := r.DB.Create(&user)

	if err != nil {
		return err.Error
	}

	return nil
}

func (r Repository) GetUser(ctx context.Context, email string) (model.User, error) {
	var user model.User
	err := r.DB.First(&user, "email = ?", email).Error

	if err != nil {
		return model.User{}, err
	}

	return user, nil
}
