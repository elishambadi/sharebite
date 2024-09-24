package services_test

import (
	"testing"

	"github.com/elishambadi/sharebite/models"
	"github.com/elishambadi/sharebite/services"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) FindAll() ([]models.User, error) {
	args := m.Called()
	return args.Get(0).([]models.User), args.Error(1)
}
func (m *MockUserRepository) Create(newUser models.User) error {
	args := m.Called(newUser)
	return args.Error(0)
}
func (m *MockUserRepository) UpdateAPIToken(user models.User, ApiToken string) error {
	args := m.Called(user, ApiToken)
	return args.Error(0)
}

func (m *MockUserRepository) GetUserById(id string) (models.User, error) {
	args := m.Called()
	return args.Get(0).(models.User), args.Error(1)
}
func (m *MockUserRepository) GetUserByEmail(email string) (models.User, error) {
	args := m.Called()
	return args.Get(0).(models.User), args.Error(1)
}
func (m *MockUserRepository) DeleteUserById(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestGetUsers(t *testing.T) {
	// Get users from the service
	t.Run("get users from service", func(t *testing.T) {
		mockRepo := new(MockUserRepository)
		logger, _ := zap.NewDevelopment()
		userService := services.NewUserService(mockRepo, logger)

		fakeUsers := models.FakeUsers()

		mockRepo.On("FindAll").Return(fakeUsers, nil)
		users, err := userService.GetUsers()

		assert.Equal(t, err, nil)
		assert.Equal(t, len(users), 2)
		assert.NoError(t, err)
		assert.Equal(t, fakeUsers, users)
		mockRepo.AssertExpectations(t)
	})
}
