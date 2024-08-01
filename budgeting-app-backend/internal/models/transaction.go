package models

import (
	"time"

	"gorm.io/gorm"
)

// Transaction represents a financial transaction
// swagger:model
type Transaction struct {
	// The ID of the transaction
	// required: true
	ID uint `gorm:"primaryKey" json:"id"`
	// The time the transaction was created
	// required: true
	CreatedAt time.Time `json:"created_at"`
	// The time the transaction was last updated
	// required: true
	UpdatedAt time.Time `json:"updated_at"`
	// The time the transaction was deleted
	// required: false
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	// The ID of the account associated with the transaction
	// required: true
	AccountID uint `json:"account_id"`
	// The amount of the transaction
	// required: true
	Amount float64 `json:"amount"`
	// The date of the transaction
	// required: true
	Date time.Time `json:"date"`
	// The category of the transaction
	// required: true
	Category string `json:"category"`
	// Additional notes about the transaction
	// required: false
	Note string `json:"note"`
}
