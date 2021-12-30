package authservice

import (
	"togolist_gin/config"
	"togolist_gin/models"
	"golang.org/x/crypto/bcrypt"
)

func compare_password(userPassword, password string) bool{
	err := bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(password))
	return err == nil
}

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