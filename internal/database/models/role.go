package models

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	Name        string `json:"name" gorm:"unique;not null"` //e.g., employee or approver
	Description string `json:"description"`
	Users       []User `json:"users"` // Has many Users
}
