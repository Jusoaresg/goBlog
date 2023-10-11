package routes

import (
	"blogbackend/controller"
	"blogbackend/middleware"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App){
  app.Post("/api/register", controller.Register)
  app.Post("/api/login", controller.Login)

  app.Use(middleware.IsAuthenticated)

  app.Post("/api/post", controller.CreatePost)
  app.Delete("/api/post/:id", controller.DeletePost)
  app.Put("/api/allpost/:id", controller.UpdatePost)

  app.Get("/api/allpost", controller.AllPost)
  app.Get("/api/allpost/:id", controller.DetailPost)
  app.Get("/api/uniquepost", controller.UniquePost)

  app.Post("/api/upload-image", controller.Upload)
  app.Static("/api/uploads", "./uploads")
}
