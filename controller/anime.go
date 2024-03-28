package controller

import (
	requestparser "AnimeList/RequestParser"
	responseparser "AnimeList/ResponseParser"
	repository "AnimeList/repository"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetAnime(repo repository.AnimeRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ID := ctx.Param("id")

		anime, err := repo.GetAnime(ctx, ID)
		if err != nil {
			println(err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "couldnt retrive data"})
			return
		}

		var res responseparser.AnimeResponse

		res.ParseAnime(anime)

		ctx.JSON(http.StatusOK, gin.H{
			"data": res,
		})
	}
}

func GetAllAnime(repo repository.AnimeRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userID, exists := ctx.Get("userid")
		if !exists {
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		userIDStr, ok := userID.(string)

		if !ok {
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		animes, err := repo.GetAllAnime(ctx, userIDStr)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "couldnt retrive data"})
			return
		}

		animeResponses := make([]responseparser.AnimeResponse, len(animes))
		for i, anime := range animes {
			var animeResponse responseparser.AnimeResponse
			animeResponse.ParseAnime(anime)
			animeResponses[i] = animeResponse
		}

		ctx.JSON(http.StatusOK, gin.H{"data": animeResponses})
	}
}

func AddAnime(repo repository.AnimeRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var reqBody requestparser.AnimeRequest

		userID, exists := ctx.Get("userid")
		if !exists {
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		userIDStr, ok := userID.(string)
		if !ok {
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		// Parse userIDStr into a UUID type
		userIDUUID, err := uuid.Parse(userIDStr)
		if err != nil {
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		err = ctx.BindJSON(&reqBody)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		anime, err := reqBody.ParseRequest()

		anime.UserID = userIDUUID

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

func EditAnime(repo repository.AnimeRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userID, exists := ctx.Get("userid")
		if !exists {
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		userIDStr, ok := userID.(string)
		if !ok {
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		userIDUUID, err := uuid.Parse(userIDStr)

		if err != nil {
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		ID := ctx.Param("id")

		var reqBody requestparser.AnimeRequest

		err = ctx.BindJSON(&reqBody)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		uuid, err := uuid.Parse(ID)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		anime, err := reqBody.ParseRequestWithID(uuid)
		anime.UserID = userIDUUID

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		err = repo.EditAnime(ctx, anime)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "success updating anime with ID " + anime.ID.String(),
		})

	}
}

func DeleteAnime(repo repository.AnimeRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ID := ctx.Param("id")

		fmt.Println(ID)
		deletedAnime, err := repo.DeleteAnime(ctx, ID)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		var res responseparser.AnimeResponse

		res.ParseAnime(deletedAnime)

		ctx.JSON(http.StatusOK, gin.H{
			"data": res,
		})
	}
}
