package example_testify_test

import (
	"errors"
	"testing"

	go_testing_tools "github.com/salmander/go-testing-tools"
	"github.com/salmander/go-testing-tools/example_testify/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestProductSearch_GetProductReturnsErrorIfNoProductIsFound(t *testing.T) {
	//Arrange
	ean := "1234"

	mockProductRepository := new(mocks.MockProductRepository)
	mockLogger := new(mocks.MockLogger)

	productRepositoryError := errors.New("some error")

	mockProductRepository.
		On("FindProductByEan", ean).
		Times(1).
		Return(go_testing_tools.Product{}, productRepositoryError)

	mockLogger.
		On("Log", mock.Anything, []interface{}{ean, productRepositoryError}).
		Times(1)

	productSearch := go_testing_tools.ProductSearch{
		ProductRepo: mockProductRepository,
		Logger:      mockLogger,
	}

	// Act
	actual, err := productSearch.GetProduct(ean)

	// Assert
	mockProductRepository.AssertExpectations(t)
	mockLogger.AssertExpectations(t)

	assert.IsType(t, err, go_testing_tools.ProductRetrieveError(errors.New("err")))
	assert.Equal(t, actual, go_testing_tools.Product{})
}
