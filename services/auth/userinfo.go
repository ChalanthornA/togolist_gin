package authservice

import (
	"togolist_gin/config"
	"togolist_gin/models"
)

func UserInfo(id float64)(models.UserModel, error){
	db := config.Db
	user := models.UserModel{}
	if err := db.Where("id = ?", int(id)).First(&user).Error; err != nil{
		return user, err
	}
	return user, nil
}