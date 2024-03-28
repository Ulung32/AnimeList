package responseparser

import (
	"AnimeList/model"
	"time"

	"github.com/google/uuid"
)

type AnimeResponse struct {
	ID         uuid.UUID
	Title      string
	Synopsis   string
	RelaseDate time.Time
}

func (res *AnimeResponse) ParseAnime(anime model.Anime) {
	res.ID = anime.ID
	res.RelaseDate = anime.RelaseDate
	res.Synopsis = anime.Synopsis
	res.Title = anime.Title
}
