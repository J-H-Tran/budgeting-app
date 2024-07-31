package main

import (
	"log"

	"github.com/J-H-Tran/budgeting-app-backend/config"
	_ "github.com/J-H-Tran/budgeting-app-backend/docs" // Swagger docs
	"github.com/J-H-Tran/budgeting-app-backend/internal/routes"
	"github.com/gin-gonic/gin"
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

	r := gin.Default()

	// Swagger UI route
	r.GET("/swagger/*any", gin.WrapH(httpSwagger.WrapHandler))

	// Register routes
	routes.BudgetRoutes(r)

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Welcome to the Budgeting App Backend")
	})

	log.Fatal(r.Run(":8080"))
}
