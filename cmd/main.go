package main

import (
	"boss-payback/internal/api"
	"boss-payback/internal/database"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDb()
	app := fiber.New()

	api.Routes(app)

	app.Listen(":3000")
}
