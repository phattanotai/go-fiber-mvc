package routes

import (
	"example.com/go-fiber/controller"
	"github.com/gofiber/fiber/v2"
)

func SetUserRoutes(app *fiber.App) {
	// User endpoints
	app.Post("/api/users", controller.CreateUser)
	app.Get("/api/users", controller.GetUsers)
	app.Get("/api/users/:id", controller.GetUser)
	app.Delete("/api/users/:id", controller.DeleteUser)
}
