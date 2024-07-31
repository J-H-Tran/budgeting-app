package models

type BudgetItemDTO struct {
	ID         uint    `json:"id"`
	Name       string  `json:"name"`
	Amount     float64 `json:"amount"`
	Recurrence string  `json:"recurrence"`
	Covered    float64 `json:"covered"`
}
