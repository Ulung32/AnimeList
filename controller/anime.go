package controller

import (
	requestparser "AnimeList/RequestParser"
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

func AddAnime(repo repository.AnimeRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var reqBody requestparser.AnimeRequest

		err := ctx.BindJSON(&reqBody)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		anime, err := reqBody.ParseRequest()

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		err = repo.AddAnime(ctx, anime)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "success creating anime with ID " + anime.ID.String(),
		})
	}
}
