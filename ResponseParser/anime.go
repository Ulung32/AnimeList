package responseparser

import (
	"AnimeList/model"
	"time"

	"github.com/google/uuid"
)

type AnimeResponse struct {
	ID          uuid.UUID
	Title       string
	Synopsis    string
	ReleaseDate time.Time
}

func (res *AnimeResponse) ParseAnime(anime model.Anime) {
	res.ID = anime.ID
	res.ReleaseDate = anime.ReleaseDate
	res.Synopsis = anime.Synopsis
	res.Title = anime.Title
}
