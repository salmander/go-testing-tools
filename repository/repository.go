package repository

import "github.com/salmander/go-testing-tools/entity"

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -o ../counterfeiter_mocks/repository.go --fake-name Repository . ProductRepository
//go:generate mockgen -package=gomock_mocks -destination=./../gomock_mocks/repository.go -source=repository.go
type ProductRepository interface {
	Get(ean string) (entity.Product, error)
}
