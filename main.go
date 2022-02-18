package main

import (
	"go-fiber-demo/src/database"
	"go-fiber-demo/src/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.Connect()
	database.AutoMigrate()
	database.SetupRedis()
	database.SetupCacheChannel()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	routes.Setup(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("hello weird")
	})

	app.Listen(":8000")
}
