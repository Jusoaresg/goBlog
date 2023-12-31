package main

import (
	"blogbackend/db"
	"blogbackend/routes"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
  db.Connect()
  err := godotenv.Load("app.env")
  if err != nil{
    log.Fatal("Error on load .env file")
  }
  port := os.Getenv("PORT")
  app := fiber.New()
  routes.Setup(app)
  app.Listen(port)  
}
