package middleware

import (
	"AnimeList/config"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func RequireAuth(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	println(authHeader)
	const bearerScheme = "Bearer "
	if !strings.HasPrefix(authHeader, bearerScheme) {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	tokenString := authHeader[len(bearerScheme):]

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected Signing method: %v", token.Header["alg"])
		}
		secretKey := []byte(config.Cfg.SecretKey)

		return secretKey, nil
	})

	if err != nil || token == nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if exp, ok := claims["exp"].(float64); ok {
			if float64(time.Now().Unix()) > exp {
				ctx.AbortWithStatus(http.StatusUnauthorized)
				return
			}
		}

		userID, ok := claims["id"].(string)
		if !ok {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		ctx.Set("userid", userID)

		ctx.Next()
	} else {
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}
}
