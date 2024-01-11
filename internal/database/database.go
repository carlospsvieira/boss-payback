package database

import (
	"boss-payback/internal/database/models"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var DB DbInstance

func ConnectDb() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}

	log.Println("connected")
	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("Running migrations")
	db.AutoMigrate(
		&models.User{},
		&models.Role{},
		&models.Expense{},
		&models.Workflow{},
	)

	if !checkRolesExist(db) {
		insertDefaultRoles(db)
	}

	if !checkAdminExist(db) {
		insertDefaultAdmin(db)
	}

	DB = DbInstance{
		Db: db,
	}
}

func checkRolesExist(db *gorm.DB) bool {
	var count int64
	var role models.Role

	db.Model(&role).Where("name IN ?", []string{"admin", "approver", "employee"}).Count(&count)

	return count == 3
}

func insertDefaultRoles(db *gorm.DB) {
	roles := []models.Role{
		{ID: 1, Name: "admin", Description: "allowed to everything"},
		{ID: 2, Name: "approver", Description: "allowed to approve payback"},
		{ID: 3, Name: "employee", Description: "allowed to request payback"},
	}

	for _, role := range roles {
		result := db.FirstOrCreate(&role, role)
		if result.Error != nil {
			log.Fatalf("Failed to insert default role %s: %v", role.Name, result.Error)
		}
	}
}

func checkAdminExist(db *gorm.DB) bool {
	var count int64
	var user models.User

	db.Model(&user).Where("role_id = ?", 1).Count(&count)

	return count == 1
}

func insertDefaultAdmin(db *gorm.DB) {
	admin := models.User{
		Username: os.Getenv("ADMIN_USERNAME"),
		Password: os.Getenv("ADMIN_PASSWORD"),
		RoleID:   1,
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(admin.Password), bcrypt.DefaultCost)
	if err != nil {
		return
	}

	admin.Password = string(hashedPassword)

	if err := db.FirstOrCreate(&admin).Error; err != nil {
		log.Fatalln("Failed to insert default admin")
	}
}
