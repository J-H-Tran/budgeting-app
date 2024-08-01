// internal/services/transaction_service.go
package services

import (
	"encoding/json"
	"net/http"

	"github.com/J-H-Tran/budgeting-app-backend/config"
	"github.com/J-H-Tran/budgeting-app-backend/internal/models"
	"github.com/gorilla/mux"
)

// GetTransaction retrieves a specific transaction by ID
// @Summary Get a transaction by ID
// @Description Get a transaction by ID
// @Tags transactions
// @Produce  json
// @Param id path int true "Transaction ID"
// @Success 200 {object} models.Transaction
// @Failure 404 {object} map[string]string
// @Router /transactions/{id} [get]
func GetTransaction(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var transaction models.Transaction
	if err := config.DB.First(&transaction, params["id"]).Error; err != nil {
		w.WriteHeader(http.StatusNotFound)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(transaction)
}

// GetTransactions retrieves all transactions
// @Summary Get all transactions
// @Description Get all transactions
// @Tags transactions
// @Produce  json
// @Success 200 {array} models.Transaction
// @Router /transactions [get]
func GetTransactions(w http.ResponseWriter, r *http.Request) {
	var transactions []models.Transaction
	if err := config.DB.Find(&transactions).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(transactions)
}

// CreateTransaction creates a new transaction
// @Summary Create a new transaction
// @Description Create a new transaction
// @Tags transactions
// @Accept  json
// @Produce  json
// @Param transaction body models.Transaction true "Transaction"
// @Success 201 {object} models.Transaction
// @Failure 400 {object} map[string]string
// @Router /transactions [post]
func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	var transaction models.Transaction
	if err := json.NewDecoder(r.Body).Decode(&transaction); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := config.DB.Create(&transaction).Error; err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(transaction)
}

// UpdateTransaction updates an existing transaction
// @Summary Update an existing transaction
// @Description Update an existing transaction
// @Tags transactions
// @Accept  json
// @Produce  json
// @Param id path int true "Transaction ID"
// @Param transaction body models.Transaction true "Transaction"
// @Success 200 {object} models.Transaction
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /transactions/{id} [put]
func UpdateTransaction(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var transaction models.Transaction
	if err := config.DB.First(&transaction, params["id"]).Error; err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&transaction); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := config.DB.Save(&transaction).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(transaction)
}

// DeleteTransaction deletes a transaction by ID
// @Summary Delete a transaction
// @Description Delete a transaction by ID
// @Tags transactions
// @Accept  json
// @Produce  json
// @Param id path int true "Transaction ID"
// @Success 204 {string} string "No Content"
// @Failure 404 {string} string "Not Found"
// @Failure 500 {string} string "Internal Server Error"
// @Router /transactions/{id} [delete]
func DeleteTransaction(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var transaction models.Transaction
	if err := config.DB.First(&transaction, params["id"]).Error; err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if err := config.DB.Delete(&transaction).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}
