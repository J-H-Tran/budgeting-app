package test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/J-H-Tran/budgeting-app-backend/internal/models"
	"github.com/J-H-Tran/budgeting-app-backend/internal/services"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockService is a mock implementation of the service layer
type MockService struct {
	mock.Mock
}

func (m *MockService) GetBudgetItem(id int) (*models.BudgetItem, error) {
	args := m.Called(id)
	return args.Get(0).(*models.BudgetItem), args.Error(1)
}

func (m *MockService) GetBudgetItems() ([]*models.BudgetItem, error) {
	args := m.Called()
	return args.Get(0).([]*models.BudgetItem), args.Error(1)
}

func TestGetBudgetItem(t *testing.T) {
	mockService := new(MockService)

	// Mock the GetBudgetItem method
	mockBudgetItem := &models.BudgetItem{ID: 1, Name: "Test Item"}
	mockService.On("GetBudgetItem", 1).Return(mockBudgetItem, nil)

	req, err := http.NewRequest("GET", "/api/budget-items/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	handler := services.NewBudgetItemHandler(mockService) // Inject the mock service
	router.HandleFunc("/api/budget-items/{id}", handler.GetBudgetItemHandler)
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	var budgetItem models.BudgetItem
	err = json.NewDecoder(rr.Body).Decode(&budgetItem)
	assert.NoError(t, err)
	assert.Equal(t, 1, budgetItem.ID)

	// Assert that the expectations were met
	mockService.AssertExpectations(t)
}

func TestGetBudgetItems(t *testing.T) {
	mockService := new(MockService)

	// Mock the GetBudgetItems method
	mockBudgetItems := []*models.BudgetItem{
		{ID: 1, Name: "Test Item 1"},
		{ID: 2, Name: "Test Item 2"},
	}
	mockService.On("GetBudgetItems").Return(mockBudgetItems, nil)

	req, err := http.NewRequest("GET", "/api/budget-items", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	handler := services.NewBudgetItemHandler(mockService) // Inject the mock service
	router.HandleFunc("/api/budget-items", handler.GetBudgetItemsHandler)
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	var budgetItems []*models.BudgetItem
	err = json.NewDecoder(rr.Body).Decode(&budgetItems)
	assert.NoError(t, err)
	assert.Len(t, budgetItems, 2)
	assert.Equal(t, 1, budgetItems[0].ID)
	assert.Equal(t, 2, budgetItems[1].ID)

	// Assert that the expectations were met
	mockService.AssertExpectations(t)
}
