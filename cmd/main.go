package main

import (
	"fmt"
	"log"
	"render2/internal/db"
	"render2/internal/db/models"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	database, err := db.InitializeDatabase()
	if err != nil {
		panic("Failed to initialize database!")
	}

	// Example: Fetching a user from the database by ID (just an illustration)
	var user models.User
	result := database.First(&user, 1) // Assuming user ID 1 exists
	if result.Error != nil {
		panic("Failed to fetch user!")
	}
	fmt.Printf("Fetched User: %+v\n", user)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString(user.Username)
	})

	log.Fatal(app.Listen(":3000"))
}
