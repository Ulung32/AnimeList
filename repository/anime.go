package repository

import (
	"AnimeList/model"
	"context"
)

type AnimeRepository interface {
	GetAllAnime(ctx context.Context) ([]model.Anime, error)
}

func (repo Repository) GetAllAnime(ctx context.Context) ([]model.Anime, error) {
	var animes []model.Anime

	err := repo.DB.Find(&animes)

	if err != nil {
		return nil, err.Error
	}
	return animes, nil
}
