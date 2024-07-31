package routes

import (
	"github.com/J-H-Tran/budgeting-app-backend/internal/services"
	"github.com/gin-gonic/gin"
)

func BudgetRoutes(r *gin.Engine) {
	// Grouping routes under /api/budget-items
	budgetRoutes := r.Group("/api/budget-items")
	{
		budgetRoutes.POST("", services.CreateBudgetItem)
		budgetRoutes.GET("", services.GetBudgetItems)
		budgetRoutes.GET("/:id", services.GetBudgetItem)
		budgetRoutes.PUT("/:id", services.UpdateBudgetItem)
		budgetRoutes.DELETE("/:id", services.DeleteBudgetItem)
	}
}
