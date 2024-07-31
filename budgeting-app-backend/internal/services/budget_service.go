package services

import (
	"encoding/json"
	"net/http"

	"github.com/J-H-Tran/budgeting-app-backend/config"
	"github.com/J-H-Tran/budgeting-app-backend/internal/models"
	"github.com/gorilla/mux"
)

// @Summary Create a new budget item
// @Description Create a new budget item in the database
// @Tags Budget Items
// @Accept json
// @Produce json
// @Param item body models.BudgetItem true "Budget Item"
// @Success 201 {object} models.BudgetItem
// @Router /api/budget-items [post]
func CreateBudgetItem(w http.ResponseWriter, r *http.Request) {
	var item models.BudgetItem
	json.NewDecoder(r.Body).Decode(&item)
	config.DB.Create(&item)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(item)
}

// @Summary Get all budget items
// @Description Retrieve all budget items from the database
// @Tags Budget Items
// @Accept json
// @Produce json
// @Success 200 {array} models.BudgetItem
// @Router /api/budget-items [get]
func GetBudgetItems(w http.ResponseWriter, r *http.Request) {
	var items []models.BudgetItem
	config.DB.Find(&items)
	json.NewEncoder(w).Encode(items)
}

// @Summary Get a specific budget item
// @Description Retrieve a specific budget item by ID
// @Tags Budget Items
// @Accept json
// @Produce json
// @Param id path int true "Budget Item ID"
// @Success 200 {object} models.BudgetItem
// @Router /api/budget-items/{id} [get]
func GetBudgetItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var item models.BudgetItem
	config.DB.First(&item, params["id"])
	json.NewEncoder(w).Encode(item)
}

// @Summary Update a budget item
// @Description Update a budget item in the database
// @Tags Budget Items
// @Accept json
// @Produce json
// @Param id path int true "Budget Item ID"
// @Param item body models.BudgetItem true "Budget Item"
// @Success 200 {object} models.BudgetItem
// @Router /api/budget-items/{id} [put]
func UpdateBudgetItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var item models.BudgetItem
	config.DB.First(&item, params["id"])
	json.NewDecoder(r.Body).Decode(&item)
	config.DB.Save(&item)
	json.NewEncoder(w).Encode(item)
}

// @Summary Delete a budget item
// @Description Delete a budget item from the database
// @Tags Budget Items
// @Accept json
// @Produce json
// @Param id path int true "Budget Item ID"
// @Success 200 {string} string "Item deleted"
// @Router /api/budget-items/{id} [delete]
func DeleteBudgetItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var item models.BudgetItem
	config.DB.Delete(&item, params["id"])
	json.NewEncoder(w).Encode("Item deleted")
}
