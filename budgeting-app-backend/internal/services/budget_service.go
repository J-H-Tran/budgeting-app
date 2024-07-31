package services

import (
	"net/http"

	"github.com/J-H-Tran/budgeting-app-backend/config"
	"github.com/J-H-Tran/budgeting-app-backend/internal/models"
	"github.com/gin-gonic/gin"
)

// @Summary Create a new budget item
// @Description Create a new budget item in the database
// @Tags Budget Items
// @Accept json
// @Produce json
// @Param item body models.BudgetItem true "Budget Item"
// @Success 201 {object} models.BudgetItem
// @Router /api/budget-items [post]
func CreateBudgetItem(c *gin.Context) {
	var item models.BudgetItem
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	config.DB.Create(&item)
	c.JSON(http.StatusCreated, item)
}

// @Summary Get all budget items
// @Description Retrieve all budget items from the database
// @Tags Budget Items
// @Accept json
// @Produce json
// @Success 200 {array} models.BudgetItem
// @Router /api/budget-items [get]
func GetBudgetItems(c *gin.Context) {
	var items []models.BudgetItem
	config.DB.Find(&items)
	c.JSON(http.StatusOK, items)
}

// @Summary Get a specific budget item
// @Description Retrieve a specific budget item by ID
// @Tags Budget Items
// @Accept json
// @Produce json
// @Param id path int true "Budget Item ID"
// @Success 200 {object} models.BudgetItem
// @Router /api/budget-items/{id} [get]
func GetBudgetItem(c *gin.Context) {
	id := c.Param("id")
	var item models.BudgetItem
	if result := config.DB.First(&item, id); result.Error != nil {
		c.JSON(http.StatusNotFound, result.Error.Error())
		return
	}
	c.JSON(http.StatusOK, item)
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
func UpdateBudgetItem(c *gin.Context) {
	id := c.Param("id")
	var item models.BudgetItem
	if result := config.DB.First(&item, id); result.Error != nil {
		c.JSON(http.StatusNotFound, result.Error.Error())
		return
	}
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	config.DB.Save(&item)
	c.JSON(http.StatusOK, item)
}

// @Summary Delete a budget item
// @Description Delete a budget item from the database
// @Tags Budget Items
// @Accept json
// @Produce json
// @Param id path int true "Budget Item ID"
// @Success 200 {string} string "Item deleted"
// @Router /api/budget-items/{id} [delete]
func DeleteBudgetItem(c *gin.Context) {
	id := c.Param("id")
	var item models.BudgetItem
	if result := config.DB.Delete(&item, id); result.Error != nil {
		c.JSON(http.StatusNotFound, result.Error.Error())
		return
	}
	c.JSON(http.StatusOK, "Item deleted")
}
