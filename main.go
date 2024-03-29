package main

import (
	"AnimeList/config"
	"AnimeList/model"
	"AnimeList/repository"
	"AnimeList/route"
	"fmt"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := config.LoadConfig()
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	fmt.Println(config.Cfg.DBUrl)

	db, err := gorm.Open(postgres.Open(config.Cfg.DBUrl), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	r := gin.Default()

	// CORS configuration
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:5173"}
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	corsConfig.AllowHeaders = []string{"Origin", "Content-Type"}
	corsConfig.AllowCredentials = true

	r.Use(cors.New(corsConfig))

	// Auto migration of database models
	db.AutoMigrate(&model.Anime{}, &model.User{})

	// Initialize repository
	repo := repository.NewRepository(db)

	// Setup routes
	route.Routes(r, repo)

	// Start Gin server
	err = r.Run()
	if err != nil {
		return
	}
}
