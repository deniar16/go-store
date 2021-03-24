package ports

import "github.com/deniarianto1606/go-store/order/domain"

type OrderService interface {
	FindByCode(code string) (*domain.Order, error)
	Save(product *domain.Order) error
}
