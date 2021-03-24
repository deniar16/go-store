package findbycode

import (
	"github.com/deniarianto1606/go-store/gateway/product"
	"github.com/deniarianto1606/go-store/product/domain"
	"github.com/deniarianto1606/go-store/product/ports/findbycode"
)

type useCase struct {
	productGateway product.Gateway
}

// NewUseCase initialize
func NewUseCase(productGateway product.Gateway) findbycode.UseCase {
	return &useCase{
		productGateway: productGateway,
	}
}

// FindByCode find by code use case
func (p *useCase) FindByCode(code string) (*domain.Product, error) {
	return p.productGateway.FindByCode(code)
}
