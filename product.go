package go_testing_tools

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -o example_counterfeiter/mocks/product.go --fake-name Product . ProductRepository
//go:generate mockgen -destination=./example_gomock/mocks/product.go -source=product.go
type ProductRepository interface {
	FindProductByEan(ean string) (Product, error)
}

type Product struct {
	Ean         string
	Description string
}

type ProductRetrieveError error

type ProductSearch struct {
	ProductRepo ProductRepository
	Logger      CustomLogger
}

func (p ProductSearch) GetProduct(ean string) (Product, error) {
	product, err := p.ProductRepo.FindProductByEan(ean)
	if err != nil {
		p.Logger.Log("An error occurred finding product with ean %s: %s", ean, err)
		return Product{}, ProductRetrieveError(err)
	}
	return product, nil
}

type ElasticSearch struct{}

func (es ElasticSearch) FindProductByEan(ean string) (Product, error) {
	return Product{}, nil
}