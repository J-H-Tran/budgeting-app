// internal/models/account.go
package models

import (
	"time"

	"gorm.io/gorm"
)

// Account represents a bank account
// @Description Account represents a bank account
type Account struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"created_at" example:"2006-01-02T15:04:05Z"`
	UpdatedAt time.Time      `json:"updated_at" example:"2006-01-02T15:04:05Z"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-" sql:"index" example:"2006-01-02T15:04:05Z"`
	// Name is the name of the bank account type (e.g., Checkings, Savings)
	// @Description Name is the name of the bank account type
	Name string `json:"name"`
	// Balance is the current balance of the bank account
	// @Description Balance is the current balance of the bank account
	Balance float64 `json:"balance"`
	// AccountNumber is the unique number of the bank account
	// @Description AccountNumber is the unique number of the bank account
	AccountNumber string `json:"account_number"`
	// Currency is the currency of the bank account balance
	// @Description Currency is the currency of the bank account balance
	Currency string `json:"currency"`
}
