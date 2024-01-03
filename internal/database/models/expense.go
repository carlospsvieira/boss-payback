package models

import (
	"gorm.io/gorm"
)

type Expense struct {
	gorm.Model
	UserID      uint    `json:"userId"` // Foreign key for User model
	Description string  `json:"description" gorm:"not null"`
	Amount      float64 `json:"amount" gorm:"not null"`
	Status      string  `json:"status" gorm:"default:'pending'"` // e.g., pending, approved, rejected
	ApproverID  uint    `json:"approverId"`                      // Foreign key for User model who approves
}
