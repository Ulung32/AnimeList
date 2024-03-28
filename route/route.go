package route

import (
	repository "AnimeList/repository"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine, repo repository.Repository) {
	AnimeRoute(r, repo)
	UserRoute(r, repo)
}
