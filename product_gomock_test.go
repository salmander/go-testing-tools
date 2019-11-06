package go_testing_tools_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	go_testing_tools "github.com/salmander/go-testing-tools"
	mock_go_testing_tools "github.com/salmander/go-testing-tools/mocks"
	"github.com/stretchr/testify/assert"
)


func TestProductSearch_GetProductReturnsErrorIfNoProductIsFound(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockProductRepository := mock_go_testing_tools.NewMockProductRepository(ctrl)

	productSearch := go_testing_tools.ProductSearch{
		mockProductRepository,
	}
	expectedProduct := go_testing_tools.Product{
		Ean: "1234",
		Description: "gopher",
	}
	actual, err := productSearch.GetProduct("1234")

	assert.NoError(t, err)
	assert.Equal(t, expectedProduct, actual)

}