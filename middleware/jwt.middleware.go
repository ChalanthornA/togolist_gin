package middleware

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"togolist_gin/services/auth"
	"github.com/dgrijalva/jwt-go"
)

func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer "
		authHeader := c.GetHeader("Authorization")
		tokenString := authHeader[len(BEARER_SCHEMA):]
		token, err := authservice.ValidateToken(tokenString)
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			c.Set("user_id", claims["user_id"])
			c.Set("username", claims["username"])
		} else {
			fmt.Println(err)
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}