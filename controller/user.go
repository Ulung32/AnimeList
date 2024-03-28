package controller

import (
	requestparser "AnimeList/RequestParser"
	"AnimeList/config"
	repository "AnimeList/repository"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	uuid "github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func Register(repo repository.UserRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var registerReq requestparser.Register

		err := ctx.BindJSON(&registerReq)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "error parsing register",
			})
			return
		}

		user, err := repo.GetUser(ctx, registerReq.Email)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "error get user",
			})
			return
		}

		if user.ID != uuid.Nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "email is already used",
			})
			return
		}

		userSave, _ := registerReq.ParseRegister()

		err = repo.CreateUser(ctx, userSave)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "error creating user",
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "success creating user with ID" + userSave.ID.String(),
		})
	}
}

func Login(repo repository.UserRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var loginReq requestparser.LoginRequest

		err := ctx.BindJSON(&loginReq)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "error parsing login",
			})
			return
		}

		user, err := repo.GetUser(ctx, loginReq.Email)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "error geting user",
			})
			return
		}

		if user.ID == uuid.Nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": "invalid email",
			})
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginReq.Password))

		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "invalid password",
			})
			return
		}

		jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"id":  user.ID,
			"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
		})

		tokenString, err := jwtToken.SignedString([]byte(config.Cfg.SecretKey))

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "error creating token",
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"token": tokenString,
		})
	}
}
