package main

import (
	"github/asuzukosi/fiber-api/databases"
	"github/asuzukosi/fiber-api/handlers"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	databases.ConnectDatabase()
	app := fiber.New()

	handlers.SetUpRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
