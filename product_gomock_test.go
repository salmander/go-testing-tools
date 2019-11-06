package go_testing_tools_test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	go_testing_tools "github.com/salmander/go-testing-tools"
	"github.com/salmander/go-testing-tools/gomock_mocks"
	"github.com/stretchr/testify/assert"
)

func TestProductSearch_GetProductReturnsErrorIfNoProductIsFound_GoMock(t *testing.T) {
	//Arrange
	ean := "1234"

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockProductRepository := gomock_mocks.NewMockProductRepository(ctrl)
	mockLogger := gomock_mocks.NewMockCustomLogger(ctrl)

	productRepositoryError := errors.New("some error")

	gomock.InOrder(
		mockProductRepository.EXPECT().
			FindProductByEan(ean).
			Times(1).
			Return(go_testing_tools.Product{}, productRepositoryError),
		mockLogger.EXPECT().
			Log(gomock.Any(), ean, productRepositoryError).Times(1),
	)

	productSearch := go_testing_tools.ProductSearch{
		ProductRepo: mockProductRepository,
		Logger:      mockLogger,
	}

	// Act
	actual, err := productSearch.GetProduct(ean)

	// Assert
	assert.Equal(t, err, go_testing_tools.ProductRetrieveError(productRepositoryError))
	assert.Equal(t, actual, go_testing_tools.Product{})
}
