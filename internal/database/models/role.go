package models

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	Name        string `json:"name" gorm:"unique;not null"` // employee, approver or admin
	Description string `json:"description"`
	Users       []User `json:"users"`
}
