package api

import (
	"boss-payback/internal/api/handlers"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	app.Get("/", handlers.Home)
}