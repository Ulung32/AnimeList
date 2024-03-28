package repository

import (
	"AnimeList/model"
	"context"
)

type AnimeRepository interface {
	GetAllAnime(ctx context.Context) ([]model.Anime, error)
	AddAnime(ctx context.Context, anime model.Anime) error
	EditAnime(ctx context.Context, anime model.Anime) error
	DeleteAnime(ctx context.Context, ID string) (model.Anime, error)
	GetAnime(ctx context.Context, ID string) (model.Anime, error)
}

func (repo Repository) GetAnime(ctx context.Context, ID string) (model.Anime, error) {
	var anime model.Anime

	err := repo.DB.First(&anime, "ID = ?", ID).Error

	if err != nil {
		return model.Anime{}, err
	}

	return anime, nil
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

func (repo Repository) EditAnime(ctx context.Context, anime model.Anime) error {
	err := repo.DB.Save(&anime)

	if err != nil {
		return err.Error
	}

	return nil
}

func (repo Repository) DeleteAnime(ctx context.Context, ID string) (model.Anime, error) {
	var deletedAnime model.Anime

	if err := repo.DB.First(&deletedAnime, "id = ?", ID).Error; err != nil {
		return model.Anime{}, err
	}

	if err := repo.DB.Where("id = ?", ID).Delete(&model.Anime{}).Error; err != nil {
		return model.Anime{}, err
	}
	return deletedAnime, nil

}
