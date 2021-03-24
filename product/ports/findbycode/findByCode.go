package findbycode

import "github.com/deniarianto1606/go-store/product/domain"

// UseCase find by code use case
type UseCase interface {
	FindByCode(code string) (*domain.Product, error)
}
