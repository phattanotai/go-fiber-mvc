package routes

import (
	"example.com/go-fiber/controller"
	"github.com/gofiber/fiber/v2"
)

func SetProductRoutes(app *fiber.App) {
	// Product endpoints
	app.Post("/api/products", controller.CreateProduct)
	app.Get("/api/products", controller.GetProducts)
	app.Get("/api/products/:id", controller.GetProduct)
	app.Put("/api/products/:id", controller.UpdateProduct)
}
