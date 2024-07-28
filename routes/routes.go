package routes

import (
	"github.com/gofiber/fiber/v2"
	user "karmapay/handlers/users"
	customer "karmapay/handlers/customers"
	order "karmapay/handlers/orders"
	payment "karmapay/handlers/payment"
	api "karmapay/handlers/api"
	middlewares "karmapay/middlewares"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Users() *fiber.App {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, X-Karma-Admin-Auth",
		AllowMethods: "GET, POST, HEAD, PUT, DELETE, PATCH, OPTIONS",
	}))
	v1 := app.Group("/v1")

	//User routes
	users := v1.Group("/users")
	users.Post("/create", middlewares.IsKarmaAdmin, user.CreateUser)
	users.Post("/login", user.LoginUser)

	//Customer routes
	customers := v1.Group(("/customer"))
	customers.Post("/register", middlewares.KPAPI, customer.CreateCustomer)

	//Orders routes
	orders := v1.Group("/orders")
	orders.Post("/create", middlewares.KPAPI, order.CreateOrder)

	//Payment routes
	payments := v1.Group("/payment")
	payments.Post("/verify", payment.VerifyPayment)
	payments.Post("/success", payment.SuccessPayment)

	//API routes
	apis := v1.Group("/api")
	apis.Post("/create", middlewares.IsKarmaAdmin, api.CreateAPI)

	// app.Use(cors.New())

	return app
}