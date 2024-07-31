package models

import "gorm.io/gorm"

type BudgetItem struct {
	gorm.Model
	Name       string  `json:"name"`
	Amount     float64 `json:"amount"`
	Recurrence string  `json:"recurrence"`
	Covered    float64 `json:"covered"`
}
