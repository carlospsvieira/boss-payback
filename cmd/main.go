package main

import (
	"boss-payback/internal/api"
	"boss-payback/internal/database"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDb()
	app := fiber.New()

	api.Routes(app)

	log.Fatal(app.Listen(":3000"))
}
