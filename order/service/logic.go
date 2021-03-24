package service

import (
	"github.com/deniarianto1606/go-store/gateway/order"
	orderPorts "github.com/deniarianto1606/go-store/order/ports"
	orderDomain "github.com/deniarianto1606/go-store/order/domain"
	"github.com/pkg/errors"
	"time"
)

var (
	ErrProductNotFound = errors.New("Product Not Found")
	ErrOrderNotFound = errors.New("Order Not Found")
	ErrOrderInvalid = errors.New("Order Invalid")
)

type orderService struct {
	orderGateway order.Gateway
}

func NewOrderService(orderGateway order.Gateway) orderPorts.OrderService {
	return &orderService{
		orderGateway: orderGateway,
	}
}

func (p *orderService) FindByCode(code string) (*orderDomain.Order, error)  {
	return p.orderGateway.FindOrderByCode(code)
}

func (p *orderService) Save(order *orderDomain.Order) error {
	product, _ := p.orderGateway.FindProductByCode(order.ProductCode)
	if product == nil {
		return errors.Wrap(ErrProductNotFound, "service.order.Product.NotFound")
	}
	order.PriceTotal = product.CreatedAt
	order.CreatedAt = time.Now().UTC().Unix()
	return p.orderGateway.Save(order)
}