package routes

import "github.com/gofiber/fiber/v2"

func InitRoutes(app *fiber.App) {
	SetOrderRoutes(app)
	SetUserRoutes(app)
	SetProductRoutes(app)

}
