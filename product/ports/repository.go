package ports

import (
	"github.com/deniarianto1606/go-store/product/domain"
)

type ProductRepository interface {
	FindByCode(code string) (*domain.Product, error)
	Save(product *domain.Product) error
}
