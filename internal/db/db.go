package db

import (
	"fmt"
	"log"
	"os"
	"render2/internal/db/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// InitializeDatabase initializes the PostgreSQL database connection and performs auto-migration
func InitializeDatabase() (*gorm.DB, error) {
	// Load env file
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Fetch environment variables
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		host, user, password, dbname, port)

	// Connect to the PostgreSQL database using GORM
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Auto Migrate the Models
	err = db.AutoMigrate(&models.User{}, &models.Role{}, &models.Expense{}, &models.ApprovalWorkflow{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
