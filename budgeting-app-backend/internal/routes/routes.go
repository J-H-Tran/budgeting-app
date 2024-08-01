package routes

import (
	"github.com/J-H-Tran/budgeting-app-backend/internal/services"
	"github.com/gorilla/mux"
)

// RegisterRoutes registers all application routes
func RegisterRoutes(r *mux.Router) {
	// Account routes
	r.HandleFunc("/api/accounts/{id}", services.GetAccount).Methods("GET")
	r.HandleFunc("/api/accounts", services.GetAccounts).Methods("GET")
	r.HandleFunc("/api/accounts", services.CreateAccount).Methods("POST")
	r.HandleFunc("/api/accounts/{id}", services.UpdateAccount).Methods("PUT")
	r.HandleFunc("/api/accounts/{id}", services.DeleteAccount).Methods("DELETE")

	// BudgetItem routes
	r.HandleFunc("/api/budget-items/{id}", services.GetBudgetItem).Methods("GET")
	r.HandleFunc("/api/budget-items", services.GetBudgetItems).Methods("GET")
	r.HandleFunc("/api/budget-items", services.CreateBudgetItem).Methods("POST")
	r.HandleFunc("/api/budget-items/{id}", services.UpdateBudgetItem).Methods("PUT")
	r.HandleFunc("/api/budget-items/{id}", services.DeleteBudgetItem).Methods("DELETE")
}
