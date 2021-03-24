package ports

import (
	"github.com/deniarianto1606/go-store/product/domain"
)

type ProductService interface {
	FindByCode(code string) (*domain.Product, error)
	Save(product *domain.Product) error
}
