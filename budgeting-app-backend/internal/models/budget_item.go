// internal/models/budget_item.go
package models

import "time"

// BudgetItem represents a budget item
// @Description BudgetItem represents a budget item
type BudgetItem struct {
	ID        uint       `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" sql:"index"`
	// Name is the name of the budget item
	// @Description Name is the name of the budget item
	Name string `json:"name"`
	// Amount is the amount allocated for the budget item
	// @Description Amount is the amount allocated for the budget item
	Amount float64 `json:"amount"`
}
