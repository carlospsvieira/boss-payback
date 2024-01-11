package models

type Role struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `json:"name" gorm:"unique;not null"` // employee, approver or admin
	Description string `json:"description"`
}
