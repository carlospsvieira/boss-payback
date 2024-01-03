package models

import (
	"gorm.io/gorm"
)

type ApprovalWorkflow struct {
	gorm.Model
	ExpenseID  uint   `json:"expenseId"`                       // Foreign key for Expense model
	ApproverID uint   `json:"approverId"`                      // Foreign key for User model who approves
	Status     string `json:"status" gorm:"default:'pending'"` // e.g., pending, approved, rejected
	Comments   string `json:"comments"`                        // Optional comments from the approver
}
