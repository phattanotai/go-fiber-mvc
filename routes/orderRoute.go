package routes

import (
	"example.com/go-fiber/controller"
	"github.com/gofiber/fiber/v2"
)

func SetOrderRoutes(app *fiber.App) {
	// Order endpoints
	app.Post("/api/orders", controller.CreateOrder)
	app.Get("/api/orders", controller.GetOrders)
	app.Get("/api/orders/:id", controller.GetOrder)
}
