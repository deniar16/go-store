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
	productRepository ProductRepository
}

func NewProductService(productRepository ProductRepository) ProductService {
	return &productService{
		productRepository: productRepository,
	}
}

func (p *productService) FindByCode(code string) (*Product, error)  {
	return p.productRepository.FindByCode(code)
}

func (p *productService) Save(product *Product) error {
	if err := validate.Validate(product); err != nil {
		return errors.Wrap(ErrProductNotFound, "service.Product.NotFound")
	}

	product.CreatedAt = time.Now().UTC().Unix()
	return p.productRepository.Save(product)
}
