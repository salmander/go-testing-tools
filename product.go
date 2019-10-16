package go_testing_tools

type Product struct {
	Ean string
	Description string
}

type ProductSearch struct {
	ProductRepo ProductRepository
}

func (p ProductSearch) GetProduct(ean string) (Product, error) {
	return Product{}, nil
}

type ProductRepository interface {
	GetProductByEan(ean string) (Product, error)
}
