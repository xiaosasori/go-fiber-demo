package main

import (
	"go-fiber-demo/src/database"
	"go-fiber-demo/src/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Connect()
	database.AutoMigrate()

	app := fiber.New()

	routes.Setup(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("hello weird")
	})

	app.Listen(":8000")
}
