package main

import (
	"AnimeList/config"
	"AnimeList/model"
	repository "AnimeList/repository"
	"AnimeList/route"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}
	fmt.Println(config.DBUrl)

	db, err := gorm.Open(postgres.Open(config.DBUrl), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	r := gin.Default()

	db.AutoMigrate(&model.Anime{})
	repo := repository.NewRepository(db)

	route.Routes(r, repo)

	err = r.Run()
	if err != nil {
		return
	}

}