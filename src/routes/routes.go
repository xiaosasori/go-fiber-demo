package routes

import (
	"go-fiber-demo/src/controllers"
	"go-fiber-demo/src/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	api := app.Group("api")

	api.Post("register", controllers.Register)
	api.Post("login", controllers.Login)

	apiAuthenticated := api.Use(middlewares.IsAuthenticated)
	apiAuthenticated.Get("user", controllers.User)
	apiAuthenticated.Post("logout", controllers.Logout)
}
