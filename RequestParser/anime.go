package requestparser

import (
	"AnimeList/model"
	"time"

	"github.com/google/uuid"
)

type AnimeRequest struct {
	Title       string
	Synopsis    string
	ReleaseDate string
}

func (req AnimeRequest) ParseRequest() (model.Anime, error) {
	date, err := time.Parse("2006-01-02", req.ReleaseDate)

	if err != nil {
		return model.Anime{}, err
	}
	return model.Anime{
		ID:          uuid.New(),
		Title:       req.Title,
		Synopsis:    req.Synopsis,
		ReleaseDate: date,
	}, nil
}

func (req AnimeRequest) ParseRequestWithID(ID uuid.UUID) (model.Anime, error) {
	date, err := time.Parse("2006-01-02", req.ReleaseDate)

	if err != nil {
		return model.Anime{}, err
	}
	return model.Anime{
		ID:          ID,
		Title:       req.Title,
		Synopsis:    req.Synopsis,
		ReleaseDate: date,
	}, nil
}
