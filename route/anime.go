package route

import (
	"AnimeList/controller"
	"AnimeList/repository"

	"github.com/gin-gonic/gin"
)

func AnimeRoute(r *gin.Engine, repo repository.RepositoryLogic) {
	animeGroup := r.Group("/anime")

	animeGroup.GET("/", controller.GetAllAnime(repo))
	animeGroup.POST("/", controller.AddAnime(repo))
}
