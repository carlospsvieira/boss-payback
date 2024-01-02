package main

import (
	"log"
	"render2/internal/api"
	"render2/internal/database"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDb()
	app := fiber.New()

	api.Routes(app)

	log.Fatal(app.Listen(":3000"))
}
