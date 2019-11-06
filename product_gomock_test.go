package go_testing_tools_test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	go_testing_tools "github.com/salmander/go-testing-tools"
	mock_go_testing_tools "github.com/salmander/go-testing-tools/mocks"
	"github.com/stretchr/testify/assert"
)

func TestProductSearch_GetProductReturnsErrorIfNoProductIsFound(t *testing.T) {
	//Arrange
	ean := "1234"

	expectedErr := errors.New("error finding product")

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockProductRepository := mock_go_testing_tools.NewMockProductRepository(ctrl)

	gomock.InOrder(
		mockProductRepository.EXPECT().
			FindProductByEan(ean).
			Times(1).
			Return(go_testing_tools.Product{}, expectedErr),
	)

	productSearch := go_testing_tools.ProductSearch{
		ProductRepo: mockProductRepository,
	}

	// Act
	actual, err := productSearch.GetProduct(ean)

	// Assert
	assert.Equal(t, err, expectedErr)
	assert.Equal(t, actual, go_testing_tools.Product{})
}
