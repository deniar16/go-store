package product

import (
	"github.com/pkg/errors"
	"gopkg.in/dealancer/validate.v2"
	"time"
)

var (
	ErrProductNotFound = errors.New("Product Not Found")
	ErrProductInvalid = errors.New("Product Invalid")
)

type productService struct {
	productGateway Gateway
}

func NewProductService(productGateway Gateway) ProductService {
	return &productService{
		productGateway: productGateway,
	}
}

func (p *productService) FindByCode(code string) (*Product, error)  {
	return p.productGateway.FindByCode(code)
}

func (p *productService) Save(product *Product) error {
	if err := validate.Validate(product); err != nil {
		return errors.Wrap(ErrProductNotFound, "service.Product.NotFound")
	}

	product.CreatedAt = time.Now().UTC().Unix()
	return p.productGateway.Save(product)
}
