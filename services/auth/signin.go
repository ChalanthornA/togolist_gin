package authservice

import (
	"togolist_gin/config"
	"togolist_gin/models"
	"golang.org/x/crypto/bcrypt"
	// "time"
	// "github.com/golang-jwt/jwt/v4"
)

func compare_password(userPassword, password string) bool{
	err := bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(password))
	return err == nil
}

// func GenerateToken(user models.UserModel)(string, error){
// 	token := jwt.New(jwt.SigningMethodHS256)
// 	claims := token.Claims.(jwt.MapClaims)
// 	claims["id"] = user.ID
//     claims["username"] = user.Username
//     claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
// 	secret := "Secret"
// 	accessToken, err := token.SignedString([]byte(secret))
// 	return accessToken, err
// }

func Signin(signinUser models.SigninModel)(models.UserModel, bool){
	db := config.Db
	user := models.UserModel{}
	if err := db.Where("username = ?", signinUser.Username).First(&user).Error; err != nil{
		return user, false
	}
	if compare_result := compare_password(user.Password, signinUser.Password); !compare_result{
		return models.UserModel{}, false
	}
	return user, true
}