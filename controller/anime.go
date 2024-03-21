package controller

import (
	repository "AnimeList/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllAnime(repo repository.AnimeRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		animes, err := repo.GetAllAnime(ctx)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "couldnt retrive data"})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"data": animes})
	}
}
