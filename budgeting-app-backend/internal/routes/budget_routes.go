package routes

import (
	"github.com/J-H-Tran/budgeting-app-backend/internal/services"
	"github.com/gorilla/mux"
)

func BudgetRoutes(r *mux.Router) {
	r.HandleFunc("/api/budget-items", services.CreateBudgetItem).Methods("POST")
	r.HandleFunc("/api/budget-items", services.GetBudgetItems).Methods("GET")
	r.HandleFunc("/api/budget-items/{id}", services.GetBudgetItem).Methods("GET")
	r.HandleFunc("/api/budget-items/{id}", services.UpdateBudgetItem).Methods("PUT")
	r.HandleFunc("/api/budget-items/{id}", services.DeleteBudgetItem).Methods("DELETE")
}
