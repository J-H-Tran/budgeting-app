package main

import (
	"log"
	"net/http"

	"github.com/J-H-Tran/budgeting-app-backend/config"
	_ "github.com/J-H-Tran/budgeting-app-backend/docs" // Swagger docs
	"github.com/J-H-Tran/budgeting-app-backend/internal/routes"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger" // Swagger UI
)

// @title Budgeting API
// @version 1.0
// @description This is a sample API for a budgeting application.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @host localhost:8080
// @BasePath /api
func main() {
	config.ConnectDB()
	r := mux.NewRouter()
	routes.BudgetRoutes(r)

	// Swagger UI route
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	r.HandleFunc("/", HomeHandler)

	log.Fatal(http.ListenAndServe(":8080", r))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the Budgeting App Backend"))
}
