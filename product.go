package go_testing_tools

import (
	"github.com/salmander/go-testing-tools/entity"
	"github.com/salmander/go-testing-tools/repository"
)

type ProductRetrieveError error

type ProductService struct {
	ProductRepo repository.ProductRepository
	Logger      CustomLogger
}

func (p ProductService) GetProduct(ean string) (entity.Product, error) {
	product, err := p.ProductRepo.Get(ean)
	if err != nil {
		p.Logger.Log("An error occurred finding product with ean %s: %s", ean, err)
		return entity.Product{}, ProductRetrieveError(err)
	}
	return product, nil
}
