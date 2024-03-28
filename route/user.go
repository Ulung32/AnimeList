package route

import (
	"AnimeList/controller"
	"AnimeList/repository"

	"github.com/gin-gonic/gin"
)

func UserRoute(r *gin.Engine, repo repository.RepositoryLogic) {
	animeGroup := r.Group("/auth")

	animeGroup.POST("/sign-up", controller.Register(repo))
	animeGroup.POST("sign-in", controller.Login(repo))
}
