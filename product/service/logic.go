package service

import (
	"github.com/deniarianto1606/go-store/gateway/product"
	"github.com/deniarianto1606/go-store/product/domain"
	"github.com/deniarianto1606/go-store/product/ports"
	"github.com/pkg/errors"
	"gopkg.in/dealancer/validate.v2"
	"time"
)

var (
	ErrProductNotFound = errors.New("Product Not Found")
	ErrProductInvalid = errors.New("Product Invalid")
)

type productService struct {
	productGateway product.Gateway
}

func NewProductService(productGateway product.Gateway) ports.ProductService {
	return &productService{
		productGateway: productGateway,
	}
}

func (p *productService) FindByCode(code string) (*domain.Product, error)  {
	return p.productGateway.FindByCode(code)
}

func (p *productService) Save(product *domain.Product) error {
	if err := validate.Validate(product); err != nil {
		return errors.Wrap(ErrProductNotFound, "service.Product.NotFound")
	}

	product.CreatedAt = time.Now().UTC().Unix()
	return p.productGateway.Save(product)
}
