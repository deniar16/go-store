package product

import (
	"fmt"
	"github.com/deniarianto1606/go-store/product/domain"
	"github.com/deniarianto1606/go-store/product/ports"
)

// Gateway product
type Gateway interface {
	FindByCode(code string) (*domain.Product, error)
	Save(product *domain.Product) error
}

type gateway struct {
	redisRepository ports.ProductRepository
	mongoRepository ports.ProductRepository
}

func NewProductGateway(redisRepository ports.ProductRepository,
	mongoRepository ports.ProductRepository) Gateway {
	return &gateway{
		redisRepository: redisRepository,
		mongoRepository: mongoRepository,
	}
}

func (p *gateway) FindByCode(code string) (*domain.Product, error) {
	product, _ := p.redisRepository.FindByCode(code)
	if product != nil {
		return product, nil
	}
	productMongo, _ := p.mongoRepository.FindByCode(code)
	if productMongo != nil {
		err := p.redisRepository.Save(productMongo)
		if err != nil {
			fmt.Println("failed save to redis: ", err)
		}
	}
	return productMongo, nil

}

func (p *gateway) Save(product *domain.Product) error {
	return p.mongoRepository.Save(product)

}
