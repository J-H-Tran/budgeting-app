// internal/models/account.go
package models

import "time"

// Account represents a user account
// @Description Account represents a user account
type Account struct {
	ID        uint       `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" sql:"index"`
	// Name is the name of the account holder
	// @Description Name is the name of the account holder
	Name string `json:"name"`
	// Email is the email of the account holder
	// @Description Email is the email of the account holder
	Email string `json:"email"`
}
