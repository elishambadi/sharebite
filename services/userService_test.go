package services_test

import (
	"testing"

	"github.com/elishambadi/sharebite/models"
	"github.com/elishambadi/sharebite/services"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) FindAll() ([]models.User, error) {
	args := m.Called()
	return args.Get(0).([]models.User), args.Error(1)
}

func TestGetUsers(t *testing.T) {
	// Get users from the service
	t.Run("get users from service", func(t *testing.T) {
		mockRepo := new(MockUserRepository)
		userService := services.NewUserService(mockRepo)

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
