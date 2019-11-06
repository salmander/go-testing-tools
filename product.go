package go_testing_tools

//go:generate mockgen -destination=./mocks/product.go -source=product.go
type ProductRepository interface {
	FindProductByEan(ean string) (Product, error)
}

type Product struct {
	Ean         string
	Description string
}

type ProductSearch struct {
	ProductRepo ProductRepository
}

func (p ProductSearch) GetProduct(ean string) (Product, error) {
	return Product{}, nil
}
