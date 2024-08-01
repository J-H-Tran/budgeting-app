// internal/services/account_service.go
package services

import (
	"encoding/json"
	"net/http"

	"github.com/J-H-Tran/budgeting-app-backend/config"
	"github.com/J-H-Tran/budgeting-app-backend/internal/models"
	"github.com/gorilla/mux"
)

// GetAccount retrieves an account by ID
// @Summary Retrieve an account by ID
// @Description Get an account by its ID
// @Tags accounts
// @Accept  json
// @Produce  json
// @Param id path string true "Account ID"
// @Success 200 {object} models.Account
// @Failure 404 {string} string "Not Found"
// @Failure 500 {string} string "Internal Server Error"
// @Router /accounts/{id} [get]
func GetAccount(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var account models.Account
	if err := config.DB.First(&account, params["id"]).Error; err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(account)
}

// GetAccounts retrieves all accounts
// @Summary Retrieve all accounts
// @Description Get a list of all accounts
// @Tags accounts
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Account
// @Failure 500 {string} string "Internal Server Error"
// @Router /accounts [get]
func GetAccounts(w http.ResponseWriter, r *http.Request) {
	var accounts []models.Account
	if err := config.DB.Find(&accounts).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accounts)
}

// CreateAccount creates a new account
// @Summary Create a new account
// @Description Create a new account with the input payload
// @Tags accounts
// @Accept  json
// @Produce  json
// @Param account body models.Account true "Account"
// @Success 201 {object} models.Account
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /accounts [post]
func CreateAccount(w http.ResponseWriter, r *http.Request) {
	var account models.Account
	if err := json.NewDecoder(r.Body).Decode(&account); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := config.DB.Create(&account).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(account)
}

// UpdateAccount updates an existing account by ID
// @Summary Update an existing account
// @Description Update an existing account with the input payload
// @Tags accounts
// @Accept  json
// @Produce  json
// @Param id path string true "Account ID"
// @Param account body models.Account true "Account"
// @Success 200 {object} models.Account
// @Failure 400 {string} string "Bad Request"
// @Failure 404 {string} string "Not Found"
// @Failure 500 {string} string "Internal Server Error"
// @Router /accounts/{id} [put]
func UpdateAccount(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var account models.Account
	if err := config.DB.First(&account, params["id"]).Error; err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&account); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := config.DB.Save(&account).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(account)
}

// DeleteAccount deletes an account by ID
// @Summary Delete an account
// @Description Delete an account by ID
// @Tags accounts
// @Accept  json
// @Produce  json
// @Param id path string true "Account ID"
// @Success 204 {string} string "No Content"
// @Failure 404 {string} string "Not Found"
// @Failure 500 {string} string "Internal Server Error"
// @Router /accounts/{id} [delete]
func DeleteAccount(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var account models.Account
	if err := config.DB.First(&account, params["id"]).Error; err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if err := config.DB.Delete(&account).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}
