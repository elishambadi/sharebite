package controllers_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
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

func (m *MockUserService) CreateUser(newUser models.User) error {
	args := m.Called(newUser)
	return args.Error(0)
}

func (m *MockUserService) GetUserById(id string) (models.User, error) {
	args := m.Called(id)

	user := args.Get(0).(models.User)
	return user, args.Error(1)
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

// --------------------------- TESTS RUN HERE -------------------------

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

func TestGetUserById(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("Success get user", func(t *testing.T) {
		gin.SetMode(gin.TestMode)

		r := gin.Default()

		fakeUsers := models.FakeUsers()
		user := fakeUsers[0]

		mockService := new(MockUserService)
		id := "1"
		r.GET("/users/:id", controllers.GetUserByIdHandler(mockService))
		mockService.On("GetUserById", id).Return(user, nil)

		req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/users/%d", user.ID), nil)
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		// Assertions
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), user.Email)
	})

	t.Run("No user found", func(t *testing.T) {
		gin.SetMode(gin.TestMode)

		r := gin.Default()
		mockService := new(MockUserService)
		id := "9999999"
		mockService.On("GetUserById", id).Return(models.User{}, errors.New("record not found"))
		// Add handler
		r.GET("/users/:id", controllers.GetUserByIdHandler(mockService))

		// Create requests
		req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/users/%s", id), nil)
		writer := httptest.NewRecorder()

		r.ServeHTTP(writer, req)

		// Assertions
		assert.Equal(t, http.StatusNotFound, writer.Code)
		assert.Contains(t, writer.Body.String(), "no user found")
		assert.Contains(t, writer.Body.String(), "[]")
	})
}

// Test user creation process

func TestCreateUser(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("create successful", func(t *testing.T) {
		r := gin.Default()
		mockService := new(MockUserService)
		fakeUser := models.User{
			Name:     "Mark",
			Email:    "mark4334543@gmail.com",
			Password: "12345678",
			Type:     "DONOR",
		}

		mockService.On("CreateUser", fakeUser).Return(nil)

		r.POST("/signup", controllers.CreateUserHandler(mockService))

		// Run the requests
		body, _ := json.Marshal(fakeUser)
		req, _ := http.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(body))
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		// Assert tests!
		assert.Equal(t, http.StatusCreated, w.Code)
		assert.Contains(t, w.Body.String(), "user created successfully")
	})

	t.Run("create failed", func(t *testing.T) {
		r := gin.Default()
		mockService := new(MockUserService)
		fakeUser := models.User{
			Name:     "Mark",
			Email:    "mark4334543@gmail.com",
			Password: "12345678",
			Type:     "RESP",
		}

		mockService.On("CreateUser", fakeUser).Return(errors.New("error creating user"))

		r.POST("/signup", controllers.CreateUserHandler(mockService))

		// Run the requests
		body, _ := json.Marshal(fakeUser)
		req, _ := http.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(body))
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		// Assert tests!
		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Contains(t, w.Body.String(), "error creating new user")
	})
}
