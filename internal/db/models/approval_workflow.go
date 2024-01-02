package models

import (
	"gorm.io/gorm"
)

type ApprovalWorkflow struct {
	gorm.Model
	ExpenseID  uint   // Foreign key for Expense model
	ApproverID uint   // Foreign key for User model who approves
	Status     string `gorm:"default:'pending'"` // e.g., pending, approved, rejected
	Comments   string // Optional comments from the approver
}
