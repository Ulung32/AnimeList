package repository

import (
	"AnimeList/model"
	"context"
)

type AnimeRepository interface {
	GetAllAnime(ctx context.Context) ([]model.Anime, error)
	AddAnime(ctx context.Context, anime model.Anime) error
}

func (repo Repository) GetAllAnime(ctx context.Context) ([]model.Anime, error) {
	var animes []model.Anime

	err := repo.DB.Find(&animes).Error

	if err != nil {
		return nil, err
	}
	return animes, nil
}

func (repo Repository) AddAnime(ctx context.Context, anime model.Anime) error {
	err := repo.DB.Create(&anime)

	if err != nil {
		return err.Error
	}
	return nil
}
