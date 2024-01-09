package models

import (
	"gorm.io/gorm"
)

type Workflow struct {
	gorm.Model
	ExpenseID  uint   `json:"expenseId"`
	ApproverID uint   `json:"approverId"`
	Status     string `json:"status" gorm:"default:'pending'"` // pending, approved, rejected
	Comments   string `json:"comments"`                        // Optional comments from the approver
}
