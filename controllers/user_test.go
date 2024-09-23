package controllers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/elishambadi/sharebite/controllers"
	"github.com/elishambadi/sharebite/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) GetUsers() ([]models.User, error) {
	args := m.Called()
	return args.Get(0).([]models.User), args.Error(1)
}

func TestGetUsers(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := new(MockUserService)

	r := gin.Default()

	r.GET("/users", controllers.GetUsersHandler(mockService))

	t.Run("Success test", func(t *testing.T) {
		fakeUsers := models.FakeUsers()

		mockService.On("GetUsers").Return(fakeUsers, nil)

		req, _ := http.NewRequest(http.MethodGet, "/users", nil)
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), "Request successful")
		assert.Contains(t, w.Body.String(), "Alice")
		assert.Contains(t, w.Body.String(), "Bob")
		mockService.AssertExpectations(t)
	})

	t.Run("Error test", func(t *testing.T) {
		mockService.On("GetUsers").Return([]models.User{}, assert.AnError)

		req, _ := http.NewRequest(http.MethodGet, "/users", nil)
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		// Assertions
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), "Error getting users")
		mockService.AssertExpectations(t)
	})
}
