package example_counterfeiter_test

import (
	"errors"
	"testing"

	go_testing_tools "github.com/salmander/go-testing-tools"
	"github.com/salmander/go-testing-tools/example_counterfeiter/mocks"
	"github.com/stretchr/testify/assert"
)

func TestProductSearch_GetProductReturnsErrorIfNoProductIsFound(t *testing.T) {
	//Arrange
	ean := "1234"

	mockProductRepository := &mocks.Product{}
	mockLogger := &mocks.Log{}

	productRepositoryError := errors.New("some error")

	mockProductRepository.FindProductByEanReturnsOnCall(0, go_testing_tools.Product{}, productRepositoryError)

	productSearch := go_testing_tools.ProductSearch{
		ProductRepo: mockProductRepository,
		Logger:      mockLogger,
	}

	// Act
	actual, err := productSearch.GetProduct(ean)

	// Assert
	assert.Equal(t, 1, mockProductRepository.FindProductByEanCallCount())
	callingEan := mockProductRepository.FindProductByEanArgsForCall(0)
	assert.Equal(t, ean, callingEan)

	assert.Equal(t, 1, mockLogger.LogCallCount())
	_, callingArgs := mockLogger.LogArgsForCall(0)
	assert.Equal(t, ean, callingArgs[0])
	assert.Equal(t, productRepositoryError, callingArgs[1])

	assert.IsType(t, err, go_testing_tools.ProductRetrieveError(errors.New("err")))
	assert.Equal(t, actual, go_testing_tools.Product{})
}
