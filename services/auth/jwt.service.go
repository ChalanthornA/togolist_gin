package authservice

import (
	"fmt"
	"time"
	"togolist_gin/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"strings"
)

func GenerateToken(user models.UserModel) (string, error) {
	claims := jwt.MapClaims{}
	claims["user_id"] = user.ID
	claims["username"] = user.Username
	claims["exp"] = time.Now().Add(time.Hour).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("Secret"))
}

func ExtractToken(c *gin.Context) string {
	token := c.Query("token")
	if token != "" {
		return token
	}
	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

func ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("invalid token")
		}
		return []byte("Secret"), nil
	})

}