package go_testing_tools_test

import (
	gotestingtools "go-testing-tools"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductSearch_GetProductReturnsErrorIfNoProductIsFound(t *testing.T) {
	productSearch := gotestingtools.ProductSearch{}
	expectedProduct := gotestingtools.Product{
		Ean: "1234",
		Description: "gopher",
	}
	actual, err := productSearch.GetProduct("1234")

	assert.NoError(t, err)
	assert.Equal(t, expectedProduct, actual)

}