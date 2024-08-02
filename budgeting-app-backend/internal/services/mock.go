package services

import (
	"net/http"

	"github.com/J-H-Tran/budgeting-app-backend/internal/models"
)

// BudgetItemService defines the interface for the service
type BudgetItemService interface {
	GetBudgetItem(id int) (*models.BudgetItem, error)
	GetBudgetItems() ([]*models.BudgetItem, error)
}

// BudgetItemHandler handles HTTP requests for budget items
type BudgetItemHandler struct {
	service BudgetItemService
}

// NewBudgetItemHandler creates a new BudgetItemHandler with the given service
func NewBudgetItemHandler(service BudgetItemService) *BudgetItemHandler {
	return &BudgetItemHandler{service: service}
}

// GetBudgetItemHandler handles the GET /api/budget-items/{id} request
func (h *BudgetItemHandler) GetBudgetItemHandler(w http.ResponseWriter, r *http.Request) {
	// Implementation of the handler
}

// GetBudgetItemsHandler handles the GET /api/budget-items request
func (h *BudgetItemHandler) GetBudgetItemsHandler(w http.ResponseWriter, r *http.Request) {
	// Implementation of the handler
}
