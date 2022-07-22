package main

import (
	"log"

	"example.com/go-fiber/database"
	routes "example.com/go-fiber/routes"
	"github.com/gofiber/fiber/v2"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to an Awesome API")
}

func setupRoutes(app *fiber.App) {
	app.Get("/", welcome)
	app.Get("/api", welcome)
	routes.InitRoutes(app)
}
func main() {
	database.ConnectDb()

	app := fiber.New()
	setupRoutes(app)
	log.Fatal(app.Listen(":3000"))

}
