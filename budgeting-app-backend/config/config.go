package config

import (
	"log"

	"github.com/J-H-Tran/budgeting-app-backend/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("budgeting_app.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// AutoMigrate the models
	err = DB.AutoMigrate(&models.BudgetItem{}, &models.Account{}, &models.Transaction{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
}
