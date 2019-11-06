package repository

import "github.com/salmander/go-testing-tools/entity"

type ElasticSearch struct{}

func (es ElasticSearch) Get(ean string) (entity.Product, error) {
	return entity.Product{}, nil
}
