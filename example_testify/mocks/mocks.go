package mocks

import (
	"github.com/salmander/go-testing-tools/example_testify"
	"github.com/stretchr/testify/mock"
)

type MockProductRepository struct {
	mock.Mock
}

func (m *MockProductRepository) FindProductByEan(ean string) (example_testify.Product, error) {
	args := m.Called(ean)
	return args.Get(0).(example_testify.Product), args.Error(1)
}

type MockLogger struct {
	mock.Mock
}

func (m *MockLogger) Log(message string, args ...interface{}) {
	m.Called(message, args)
}
