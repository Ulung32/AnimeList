package repository

import "gorm.io/gorm"

type Repository struct {
	DB *gorm.DB
}

type RepositoryLogic interface {
	UserRepository
	AnimeRepository
}

func NewRepository(db *gorm.DB) Repository {
	return Repository{
		DB: db,
	}
}
