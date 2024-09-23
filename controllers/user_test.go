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

func (m *MockUserService) CreateUser(ctx *gin.Context) error {
	args := m.Called()
	return args.Error(1)
}

func (m *MockUserService) GetUserById(id string) (models.User, error) {
	args := m.Called()
	return models.User{}, args.Error(1)
}

func (m *MockUserService) DeleteUserById(id string) error {
	args := m.Called()
	return args.Error(1)
}

func (m *MockUserService) AuthenticateUser(ctx *gin.Context) (token string, error error) {
	args := m.Called()
	return args.Get(0).(string), args.Error(1)
}

func (m *MockUserService) ResetUserPassword(ctx *gin.Context) error {
	args := m.Called()
	return args.Error(1)
}

func (m *MockUserService) GetUserFromRequest(c *gin.Context) (models.User, error) {
	args := m.Called()
	return args.Get(0).(models.User), args.Error(1)
}

func TestGetUsers(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("Success test", func(t *testing.T) {
		fakeUsers := models.FakeUsers()

		mockService := new(MockUserService)
		r := gin.Default()
		r.GET("/users", controllers.GetUsersHandler(mockService))

		mockService.On("GetUsers").Return(fakeUsers, nil)

		req, _ := http.NewRequest(http.MethodGet, "/users", nil)
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), "request successful")
		assert.Contains(t, w.Body.String(), "Alice")
		assert.Contains(t, w.Body.String(), "Bob")
		mockService.AssertExpectations(t)
	})

	t.Run("Error test", func(t *testing.T) {
		mockService := new(MockUserService)
		r := gin.Default()
		r.GET("/users", controllers.GetUsersHandler(mockService))
		mockService.On("GetUsers").Return([]models.User{}, assert.AnError)

		req, _ := http.NewRequest(http.MethodGet, "/users", nil)
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		// Assertions
		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Contains(t, w.Body.String(), "error getting users")
		mockService.AssertExpectations(t)
	})

	t.Run("no users found", func(t *testing.T) {
		mockService := new(MockUserService)
		r := gin.Default()
		r.GET("/users", controllers.GetUsersHandler(mockService))
		mockService.On("GetUsers").Return([]models.User{}, nil)

		req, _ := http.NewRequest(http.MethodGet, "/users", nil)
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		// Assertions
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), "[]")
		mockService.AssertExpectations(t)
	})
}
