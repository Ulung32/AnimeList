package middleware

import (
	"AnimeList/config"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func RequireAuth(ctx *gin.Context) {
	tokenString, err := ctx.Cookie("Authorization")

	println(tokenString)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "error parsing token",
		})
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return config.Cfg.SecretKey, nil
	})

	if !token.Valid {
		fmt.Errorf("Invalid token")
	}

	// Type assert the token claims to jwt.MapClaims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		fmt.Errorf("Error parsing token claims")
	}

	if float64(time.Now().Unix()) > claims["exp"].(float64) {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "expired token",
		})
		return
	}

	ctx.Set("userid", claims["id"])

	ctx.Next()
}
