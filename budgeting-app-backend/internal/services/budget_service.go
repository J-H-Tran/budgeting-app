package services

import (
	"encoding/json"
	"net/http"

	"github.com/J-H-Tran/budgeting-app-backend/config"
	"github.com/J-H-Tran/budgeting-app-backend/internal/models"
	"github.com/gorilla/mux"
)

// @Summary Create a new budget
// @Description Create a new budget with the input payload
// @Tags budgets
// @Accept  json
// @Produce  json
// @Param budget body models.Budget true "Budget"
// @Success 201 {object} models.Budget
// @Failure 400 {object} gin.H{"error": "Bad Request"}
// @Failure 500 {object} gin.H{"error": "Internal Server Error"}
// @Router /budgets [post]
func CreateBudgetItem(w http.ResponseWriter, r *http.Request) {
	var item models.BudgetItem
	json.NewDecoder(r.Body).Decode(&item)
	config.DB.Create(&item)
	json.NewEncoder(w).Encode(item)
}

// @Summary Get all budgets
// @Description Get a list of all budgets
// @Tags budgets
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Budget
// @Failure 400 {object} gin.H{"error": "Bad Request"}
// @Failure 500 {object} gin.H{"error": "Internal Server Error"}
// @Router /budgets [get]
func GetBudgetItems(w http.ResponseWriter, r *http.Request) {
	var items []models.BudgetItem
	config.DB.Find(&items)
	json.NewEncoder(w).Encode(items)
}

// @Summary Get a budget by ID
// @Description Get a budget by its ID
// @Tags budgets
// @Accept  json
// @Produce  json
// @Param id path int true "Budget ID"
// @Success 200 {object} models.Budget
// @Failure 400 {object} gin.H{"error": "Bad Request"}
// @Failure 404 {object} gin.H{"error": "Not Found"}
// @Failure 500 {object} gin.H{"error": "Internal Server Error"}
// @Router /budgets/{id} [get]
func GetBudgetItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var item models.BudgetItem
	config.DB.First(&item, params["id"])
	json.NewEncoder(w).Encode(item)
}

// @Summary Update a budget by ID
// @Description Update a budget by its ID with the input payload
// @Tags budgets
// @Accept  json
// @Produce  json
// @Param id path int true "Budget ID"
// @Param budget body models.Budget true "Budget"
// @Success 200 {object} models.Budget
// @Failure 400 {object} gin.H{"error": "Bad Request"}
// @Failure 404 {object} gin.H{"error": "Not Found"}
// @Failure 500 {object} gin.H{"error": "Internal Server Error"}
// @Router /budgets/{id} [put]
func UpdateBudgetItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var item models.BudgetItem
	config.DB.First(&item, params["id"])
	json.NewDecoder(r.Body).Decode(&item)
	config.DB.Save(&item)
	json.NewEncoder(w).Encode(item)
}

// @Summary Delete a budget by ID
// @Description Delete a budget by its ID
// @Tags budgets
// @Accept  json
// @Produce  json
// @Param id path int true "Budget ID"
// @Success 204 {object} nil
// @Failure 400 {object} gin.H{"error": "Bad Request"}
// @Failure 404 {object} gin.H{"error": "Not Found"}
// @Failure 500 {object} gin.H{"error": "Internal Server Error"}
// @Router /budgets/{id} [delete]
func DeleteBudgetItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var item models.BudgetItem
	config.DB.Delete(&item, params["id"])
	json.NewEncoder(w).Encode("Item deleted")
}
