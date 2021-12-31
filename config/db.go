package config

import (
	"fmt"
	"log"
	"os"
	"togolist_gin/models"
	
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func LoadENV(key string) string {
	err := godotenv.Load(".env")
  
	if err != nil {
	  log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
  }

func ConnectDb(){
	var err error
	dns := LoadENV("DNS")
	Db, err = gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connect to database")
	Db.AutoMigrate(&models.TodoModel{}, &models.UserModel{})
}