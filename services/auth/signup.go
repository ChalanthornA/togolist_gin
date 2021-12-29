package authservice

import (
	"togolist_gin/config"
	"togolist_gin/models"

	"golang.org/x/crypto/bcrypt"
)

func hash_password(password string) string{
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes)
}

func Checkusername(username string) bool{
	db := config.Db
	queryUser := models.UserModel{}
	if err := db.Where("username = ?", username).First(&queryUser).Error; err == nil{
		return false
	}
	return true
}

func CreateUser(newUser *models.UserModel){
	db := config.Db
	hashPassword := hash_password(newUser.Password)
	newUser.Password = hashPassword
	db.Create(&newUser)
}