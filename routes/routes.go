package routes

import (
	"github.com/gofiber/fiber/v2"
	user "karmapay/handlers/users"
	middlewares "karmapay/middlewares"
)

func Users() *fiber.App {
	app := fiber.New()
	v1 := app.Group("/v1")

	//User routes
	users := v1.Group("/users")
	users.Post("/create", middlewares.IsKarmaAdmin, user.CreateUser)

	return app
}