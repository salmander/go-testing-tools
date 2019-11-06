package go_testing_tools_test

import (
	"errors"
	"testing"

	go_testing_tools "github.com/salmander/go-testing-tools"
	"github.com/salmander/go-testing-tools/counterfeiter_mocks"
	"github.com/stretchr/testify/assert"
)

func TestProductSearch_GetProductReturnsErrorIfNoProductIsFound_Counterfeiter(t *testing.T) {
	//Arrange
	ean := "1234"

	mockProductRepository := &counterfeiter_mocks.Product{}
	mockLogger := &counterfeiter_mocks.Log{}

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

	assert.Equal(t, err, go_testing_tools.ProductRetrieveError(productRepositoryError))
	assert.Equal(t, actual, go_testing_tools.Product{})
}
