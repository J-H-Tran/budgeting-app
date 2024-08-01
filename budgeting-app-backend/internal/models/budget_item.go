// internal/models/budget_item.go
package models

import (
	"time"

	"gorm.io/gorm"
)

// BudgetItem represents a budget item
// @Description BudgetItem represents a budget item
type BudgetItem struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	// Name is the name of the budget item
	// @Description Name is the name of the budget item
	Name string `json:"name"`
	// Description is the description of the budget item
	// @Description Description is the description of the budget item
	Description string `json:"description"`
	// Amount is the amount allocated for the budget item
	// @Description Amount is the amount allocated for the budget item
	Amount float64 `json:"amount"`
}
