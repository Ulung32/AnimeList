package main

import (
	"AnimeList/config"
	"AnimeList/model"
	repository "AnimeList/repository"
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

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "*"}, // Change to specific origins if needed
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		AllowCredentials: true,
	}))

	db.AutoMigrate(&model.Anime{}, &model.User{})
	repo := repository.NewRepository(db)

	route.Routes(r, repo)

	err = r.Run()
	if err != nil {
		return
	}

}
