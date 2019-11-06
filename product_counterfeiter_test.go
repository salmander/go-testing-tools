package go_testing_tools_test

import (
	"errors"
	"testing"

	go_testing_tools "github.com/salmander/go-testing-tools"
	"github.com/salmander/go-testing-tools/counterfeiter_mocks"
	"github.com/salmander/go-testing-tools/entity"
	"github.com/stretchr/testify/assert"
)

func TestProductSearch_GetProductReturnsErrorIfNoProductIsFound_Counterfeiter(t *testing.T) {
	//Arrange
	ean := "1234"

	mockProductRepository := &counterfeiter_mocks.Repository{}
	mockLogger := &counterfeiter_mocks.Log{}

	productRepositoryError := errors.New("some error")

	mockProductRepository.GetReturnsOnCall(0, entity.Product{}, productRepositoryError)

	productSearch := go_testing_tools.ProductService{
		ProductRepo: mockProductRepository,
		Logger:      mockLogger,
	}

	// Act
	actual, err := productSearch.GetProduct(ean)

	// Assert
	assert.Equal(t, 1, mockProductRepository.GetCallCount())
	callingEan := mockProductRepository.GetArgsForCall(0)
	assert.Equal(t, ean, callingEan)

	assert.Equal(t, 1, mockLogger.LogCallCount())
	_, callingArgs := mockLogger.LogArgsForCall(0)
	assert.Equal(t, ean, callingArgs[0])
	assert.Equal(t, productRepositoryError, callingArgs[1])

	assert.Equal(t, err, go_testing_tools.ProductRetrieveError(productRepositoryError))
	assert.Equal(t, actual, entity.Product{})
}
