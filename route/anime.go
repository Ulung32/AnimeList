package route

import (
	"AnimeList/controller"
	"AnimeList/middleware"
	"AnimeList/repository"

	"github.com/gin-gonic/gin"
)

func AnimeRoute(r *gin.Engine, repo repository.RepositoryLogic) {
	animeGroup := r.Group("/anime")

	animeGroup.GET("/", middleware.RequireAuth, controller.GetAllAnime(repo))
	animeGroup.GET("/:id", middleware.RequireAuth, controller.GetAnime(repo))
	animeGroup.POST("/", middleware.RequireAuth, controller.AddAnime(repo))
	animeGroup.PUT("/:id", middleware.RequireAuth, controller.EditAnime(repo))
	animeGroup.DELETE("/:id", middleware.RequireAuth, controller.DeleteAnime(repo))
}
