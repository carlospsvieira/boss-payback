package api

import (
	"render2/internal/api/handlers"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	app.Get("/", handlers.Home)
}
