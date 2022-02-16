package routes

import (
	"go-fiber-demo/src/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	api := app.Group("api")

	api.Post("register", controllers.Register)
}
