package authservice

import (
	"fmt"
	"time"
	"togolist_gin/config"
	"togolist_gin/models"

	"github.com/dgrijalva/jwt-go"
)

var SECRET string = config.LoadENV("SECRET")

func GenerateToken(user models.UserModel) (string, error) {
	claims := jwt.MapClaims{}
	claims["user_id"] = user.ID
	claims["username"] = user.Username
	claims["exp"] = time.Now().Add(time.Hour).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(SECRET))
}

func ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("invalid token")
		}
		return []byte(SECRET), nil
	})

}