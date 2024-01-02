package models

import (
	"gorm.io/gorm"
)

type Expense struct {
	gorm.Model
	UserID      uint    // Foreign key for User model
	Description string  `gorm:"not null"`
	Amount      float64 `gorm:"not null"`
	Status      string  `gorm:"default:'pending'"` // e.g., pending, approved, rejected
	ApproverID  uint    // Foreign key for User model who approves
}
