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
	apiAuthenticated.Post("logout", controllers.Logout)
	apiAuthenticated.Get("user", controllers.User)
	apiAuthenticated.Put("users/info", controllers.UpdateInfo)
	apiAuthenticated.Put("users/password", controllers.UpdatePassword)
	apiAuthenticated.Get("ambassadors", controllers.Ambassadors)

	apiAuthenticated.Get("products", controllers.Products)
	apiAuthenticated.Post("products", controllers.CreateProducts)
	apiAuthenticated.Post("products/:id", controllers.GetProduct)
	apiAuthenticated.Put("products/:id", controllers.UpdateProduct)
	apiAuthenticated.Delete("products/:id", controllers.DeleteProduct)

	apiAuthenticated.Get("users/:id/links", controllers.Link)

	apiAuthenticated.Get("orders", controllers.Orders)

	ambassador := api.Group("ambassador")
	ambassador.Post("register", controllers.Register)
	ambassador.Post("login", controllers.Login)
	ambassador.Get("products/frontend", controllers.ProductsFrontend)
	ambassador.Get("products/backend", controllers.ProductsBackend)

	ambassadorAuthenticated := ambassador.Use(middlewares.IsAuthenticated)
	ambassadorAuthenticated.Get("user", controllers.User)
	ambassadorAuthenticated.Post("logout", controllers.Logout)
	ambassadorAuthenticated.Put("users/info", controllers.UpdateInfo)
	ambassadorAuthenticated.Put("users/password", controllers.UpdatePassword)
	ambassadorAuthenticated.Post("links", controllers.CreateLink)
	ambassadorAuthenticated.Get("stats", controllers.Stats)

}
