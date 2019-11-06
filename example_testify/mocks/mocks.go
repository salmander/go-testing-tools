package mocks

import (
	go_testing_tools "github.com/salmander/go-testing-tools"
	"github.com/stretchr/testify/mock"
)

type MockProductRepository struct {
	mock.Mock
}

func (m *MockProductRepository) FindProductByEan(ean string) (go_testing_tools.Product, error) {
	args := m.Called(ean)
	return args.Get(0).(go_testing_tools.Product), args.Error(1)
}

type MockLogger struct {
	mock.Mock
}

func (m *MockLogger) Log(message string, args ...interface{}) {
	m.Called(message, args)
}
