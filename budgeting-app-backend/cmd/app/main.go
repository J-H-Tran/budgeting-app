package main

import (
	"log"
	"net/http"

	"github.com/J-H-Tran/budgeting-app-backend/config"
	"github.com/J-H-Tran/budgeting-app-backend/internal/routes"
	"github.com/gorilla/mux"
)

func main() {
	config.ConnectDB()
	r := mux.NewRouter()
	routes.BudgetRoutes(r)

	r.HandleFunc("/", HomeHandler)

	log.Fatal(http.ListenAndServe(":8080", r))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the Budgeting App Backend"))
}
