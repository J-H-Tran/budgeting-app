// internal/services/budget_service.go
package services

import (
	"encoding/json"
	"net/http"

	"github.com/J-H-Tran/budgeting-app-backend/config"
	"github.com/J-H-Tran/budgeting-app-backend/internal/models"
	"github.com/gorilla/mux"
)

// GetBudgetItem retrieves a single budget by ID
// @Summary Retrieve a single budget item
// @Description Get a budget item by its ID
// @Tags budget_items
// @Accept  json
// @Produce  json
// @Param id path string true "Budget Item ID"
// @Success 200 {object} models.BudgetItem
// @Failure 404 {string} string "Not Found"
// @Failure 500 {string} string "Internal Server Error"
// @Router /budget_items/{id} [get]
func GetBudgetItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var budget models.BudgetItem
	if err := config.DB.First(&budget, params["id"]).Error; err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(budget)
}

// GetBudgets retrieves all budgets
// @Summary Retrieve all budgets
// @Description Get a list of all budgets
// @Tags budget_items
// @Accept  json
// @Produce  json
// @Success 200 {array} models.BudgetItem
// @Failure 500 {string} string "Internal Server Error"
// @Router /budgets [get]
func GetBudgetItems(w http.ResponseWriter, r *http.Request) {
	var budgets []models.BudgetItem
	if err := config.DB.Find(&budgets).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(budgets)
}

// CreateBudget creates a new budget
// @Summary Create a new budget
// @Description Create a new budget with the input payload
// @Tags budget_items
// @Accept  json
// @Produce  json
// @Param budget body models.BudgetItem true "Budget"
// @Success 201 {object} models.BudgetItem
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /budgets [post]
func CreateBudgetItem(w http.ResponseWriter, r *http.Request) {
	var budget models.BudgetItem
	if err := json.NewDecoder(r.Body).Decode(&budget); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := config.DB.Create(&budget).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(budget)
}

// UpdateBudget updates an existing budget by ID
// @Summary Update an existing budget
// @Description Update an existing budget with the input payload
// @Tags budget_items
// @Accept  json
// @Produce  json
// @Param id path string true "Budget ID"
// @Param budget body models.BudgetItem true "Budget"
// @Success 200 {object} models.BudgetItem
// @Failure 400 {string} string "Bad Request"
// @Failure 404 {string} string "Not Found"
// @Failure 500 {string} string "Internal Server Error"
// @Router /budgets/{id} [put]
func UpdateBudgetItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var budget models.BudgetItem
	if err := config.DB.First(&budget, params["id"]).Error; err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&budget); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := config.DB.Save(&budget).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(budget)
}

// DeleteBudget deletes a budget by ID
// @Summary Delete a budget
// @Description Delete a budget by ID
// @Tags budget_items
// @Accept  json
// @Produce  json
// @Param id path string true "Budget ID"
// @Success 204 {string} string "No Content"
// @Failure 404 {string} string "Not Found"
// @Failure 500 {string} string "Internal Server Error"
// @Router /budgets/{id} [delete]
func DeleteBudgetItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var budget models.BudgetItem
	if err := config.DB.First(&budget, params["id"]).Error; err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if err := config.DB.Delete(&budget).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}
