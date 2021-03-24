package product

type Gateway struct {
	redisRepository ProductRepository
	mongoRepository ProductRepository
}

func NewProductGateway(redisRepository ProductRepository,
	mongoRepository ProductRepository) Gateway {
	return Gateway{
		redisRepository: redisRepository,
		mongoRepository: mongoRepository,
	}
}

func (p *Gateway) FindByCode(code string) (*Product, error)  {
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

func (p *Gateway) Save(product *Product) error  {
	return p.mongoRepository.Save(product)

}