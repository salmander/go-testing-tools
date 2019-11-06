package example_gomock_test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/salmander/go-testing-tools/example_gomock"
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
			Return(example_gomock.Product{}, productRepositoryError),
		mockLogger.EXPECT().
			Log(gomock.Any(), ean, productRepositoryError).Times(1),
	)

	productSearch := example_gomock.ProductSearch{
		ProductRepo: mockProductRepository,
		Logger:      mockLogger,
	}

	// Act
	actual, err := productSearch.GetProduct(ean)

	// Assert
	assert.IsType(t, err, example_gomock.ProductRetrieveError(errors.New("err")))
	assert.Equal(t, actual, example_gomock.Product{})
}
