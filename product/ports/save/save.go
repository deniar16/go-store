package save

import "github.com/deniarianto1606/go-store/product/domain"

// UseCase save use case
type UseCase interface {
	// Save save product
	Save(product *domain.Product) error
}
