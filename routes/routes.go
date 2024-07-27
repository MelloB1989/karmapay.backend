package routes

import (
	"github.com/gofiber/fiber/v2"
	user "karmapay/handlers/users"
	customer "karmapay/handlers/customers"
	order "karmapay/handlers/orders"
	middlewares "karmapay/middlewares"
)

func Users() *fiber.App {
	app := fiber.New()
	v1 := app.Group("/v1")

	//User routes
	users := v1.Group("/users")
	users.Post("/create", middlewares.IsKarmaAdmin, user.CreateUser)

	//Customer routes
	customers := v1.Group(("/customer"))
	customers.Post("/register", middlewares.KPAPI, customer.CreateCustomer)

	//Orders routes
	orders := v1.Group("/orders")
	orders.Post("/create", middlewares.KPAPI, order.CreateOrder)

	return app
}