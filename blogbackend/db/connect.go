package db

import (
	"blogbackend/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB


func Connect() {
	err := godotenv.Load("app.env")
	if err != nil {
		log.Fatal("Error load .env file")
	}
  dsn := os.Getenv("DSN")
  database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
  if err != nil{
    log.Fatal("Could not connect to the database")
  }else{
    log.Println("Connect to the database")
  }
  DB = database
  database.AutoMigrate(
    &models.User{},
    &models.Blog{},
    )
}
