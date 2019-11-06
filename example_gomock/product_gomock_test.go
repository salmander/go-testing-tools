package example_gomock_test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	go_testing_tools "github.com/salmander/go-testing-tools"
	mock_go_testing_tools "github.com/salmander/go-testing-tools/example_gomock/mocks"
	"github.com/stretchr/testify/assert"
)

func TestProductSearch_GetProductReturnsErrorIfNoProductIsFound(t *testing.T) {
	//Arrange
	ean := "1234"

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockProductRepository := mock_go_testing_tools.NewMockProductRepository(ctrl)
	mockLogger := mock_go_testing_tools.NewMockCustomLogger(ctrl)

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
	assert.IsType(t, err, go_testing_tools.ProductRetrieveError(errors.New("err")))
	assert.Equal(t, actual, go_testing_tools.Product{})
}
