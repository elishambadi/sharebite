package mocks

import (
	"github.com/elishambadi/sharebite/models"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

// MockDB is a mock of gorm.DB
type MockDB struct {
	mock.Mock
}

// Find is a mock method that simulates the Find operation
func (m *MockDB) Find(out interface{}, where ...interface{}) *gorm.DB {
	args := m.Called(out)
	if args.Error(1) != nil {
		return &gorm.DB{Error: args.Error(1)}
	}
	*(out.(*[]models.User)) = args.Get(0).([]models.User)
	return &gorm.DB{}
}
