package handlers

import "github.com/gofiber/fiber/v2"

func SetUpRoutes(app *fiber.App) {
	app.Get("/", WelcomeHandler)
	users := app.Group("/users")

	users.Post("/", CreateUser)
	users.Get("/", GetUsers)
	users.Get("/:id", GetUser)
	users.Put("/:id", UpdateUser)
	users.Delete("/:id", DeleteUser)

	products := app.Group("/products")

	products.Post("/", CreateProduct)
	products.Get("/", GetProducts)
	products.Get("/:id", GetProduct)
	products.Put("/:id", UpdateProduct)
	products.Delete("/:id", DeleteProduct)

	orders := app.Group("/orders")

	orders.Post("/", CreateOrder)
	orders.Get("/:id", GetOrder)
	orders.Get("/", GetOrders)
}
