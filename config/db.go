package config

import (
	"fmt"
	"log"
	"togolist_gin/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func ConnectDb(){
	var err error
	dns := "postgres://uzvwnvlxdxdppv:dfbcad8e8c49a162b338d27025f2a2c9e078eeb867e20b169dd075b72aebd7e6@ec2-18-214-238-28.compute-1.amazonaws.com:5432/d6nm3gnvi54q8r"
	Db, err = gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connect to database")
	Db.AutoMigrate(&models.TodoModel{}, &models.UserModel{})
}