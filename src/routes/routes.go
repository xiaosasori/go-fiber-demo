package routes

import (
	"go-fiber-demo/src/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	api := app.Group("api")

	api.Post("register", controllers.Register)
	api.Post("login", controllers.Login)
	api.Post("logout", controllers.Logout)
	api.Get("user", controllers.User)
}
