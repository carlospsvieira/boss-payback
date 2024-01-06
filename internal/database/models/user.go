package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"unique;not null"`
	Password string `json:"password" gorm:"not null"`
	Email    string `json:"email" gorm:"unique;not null"`
	RoleID   uint   `json:"roleid"`
	Role     Role
	Expenses []Expense
	LoggedIn bool `json:"loggedIn" gorm:"default:false"`
}
