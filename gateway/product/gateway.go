package product

import (
	"github.com/deniarianto1606/go-store/product/domain"
	"github.com/deniarianto1606/go-store/product/ports"
)

type Gateway struct {
	redisRepository ports.ProductRepository
	mongoRepository ports.ProductRepository
}

func NewProductGateway(redisRepository ports.ProductRepository,
	mongoRepository ports.ProductRepository) Gateway {
	return Gateway{
		redisRepository: redisRepository,
		mongoRepository: mongoRepository,
	}
}

func (p *Gateway) FindByCode(code string) (*domain.Product, error)  {
	product, _ := p.redisRepository.FindByCode(code)
	if product != nil {
		return product, nil
	}
	productMongo, _ := p.mongoRepository.FindByCode(code)
	if productMongo != nil {
		p.redisRepository.Save(productMongo)
	}
	return productMongo, nil

}

func (p *Gateway) Save(product *domain.Product) error  {
	return p.mongoRepository.Save(product)

}