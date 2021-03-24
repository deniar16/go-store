package order

import (
	productDomain "github.com/deniarianto1606/go-store/product/domain"
	productPorts "github.com/deniarianto1606/go-store/product/ports"
	orderDomain "github.com/deniarianto1606/go-store/order/domain"
	orderPorts "github.com/deniarianto1606/go-store/order/ports"
)

type Gateway struct {
	redisProductRepository productPorts.ProductRepository
	mongoProductRepository productPorts.ProductRepository
	redisOrderRepository orderPorts.OrderRepository
	mongoOrderRepository orderPorts.OrderRepository

}

func NewOrderGateway(
	redisProductRepository productPorts.ProductRepository,
	mongoProductRepository productPorts.ProductRepository,
	redisOrderRepository orderPorts.OrderRepository,
	mongoOrderRepository orderPorts.OrderRepository,
	) Gateway {
	return Gateway{
		redisProductRepository: redisProductRepository,
		mongoProductRepository: mongoProductRepository,
		redisOrderRepository: redisOrderRepository,
		mongoOrderRepository: mongoOrderRepository,
	}
}

func (p *Gateway) FindProductByCode(code string) (*productDomain.Product, error)  {
	product, _ := p.redisProductRepository.FindByCode(code)
	if product != nil {
		return product, nil
	}
	productMongo, _ := p.mongoProductRepository.FindByCode(code)
	if productMongo != nil {
		p.redisProductRepository.Save(productMongo)
	}
	return productMongo, nil
}

func (p *Gateway) FindOrderByCode(code string) (*orderDomain.Order, error)  {
	order, _ := p.redisOrderRepository.FindByCode(code)
	if order != nil {
		return order, nil
	}
	orderMongo, _ := p.mongoOrderRepository.FindByCode(code)
	if orderMongo != nil {
		p.redisOrderRepository.Save(orderMongo)
	}
	return orderMongo, nil
}

func (p *Gateway) Save(product *orderDomain.Order) error  {
	return p.mongoOrderRepository.Save(product)

}
