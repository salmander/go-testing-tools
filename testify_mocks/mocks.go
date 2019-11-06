package testify_mocks

import (
	"github.com/salmander/go-testing-tools/entity"
	"github.com/stretchr/testify/mock"
)

type MockProductRepository struct {
	mock.Mock
}

func (m *MockProductRepository) Get(ean string) (entity.Product, error) {
	args := m.Called(ean)
	return args.Get(0).(entity.Product), args.Error(1)
}

type MockLogger struct {
	mock.Mock
}

func (m *MockLogger) Log(message string, args ...interface{}) {
	m.Called(message, args)
}
