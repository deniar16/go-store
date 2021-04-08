package save

import (
	"github.com/deniarianto1606/go-store/gateway/product"
	"github.com/deniarianto1606/go-store/order/service"
	"github.com/deniarianto1606/go-store/product/domain"
	"github.com/deniarianto1606/go-store/product/ports/save"
	"github.com/pkg/errors"
	"gopkg.in/dealancer/validate.v2"
	"time"
)

type useCase struct {
	productGateway product.Gateway
}

// NewUseCase initialize
func NewUseCase(productGateway product.Gateway) save.UseCase {
	return &useCase{
		productGateway: productGateway,
	}
}

func (p *useCase) Save(product *domain.Product) error {
	if err := validate.Validate(product); err != nil {
		return errors.Wrap(service.ErrProductNotFound, "service.Product.NotFound")
	}

	product.CreatedAt = time.Now().UTC().Unix()
	return p.productGateway.Save(product)
}
