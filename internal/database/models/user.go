package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"unique;not null"`
	Password string `json:"password" gorm:"not null"`
	Email    string `json:"email" gorm:"unique;not null"`
	RoleID   uint   `json:"roleid"` // Foreign key for Role model
	Role     Role   // Belongs to Role
}
