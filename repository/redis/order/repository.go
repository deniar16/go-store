package order

import (
	"fmt"
	"github.com/deniarianto1606/go-store/order/domain"
	"github.com/deniarianto1606/go-store/order/ports"
	"github.com/deniarianto1606/go-store/order/service"
	"github.com/go-redis/redis"
	"github.com/pkg/errors"
	"strconv"
)

type RedisRepository struct {
	client *redis.Client
}

func newRedisClient(redisUrl string) (*redis.Client, error){
	opts, err := redis.ParseURL(redisUrl)
	if err != nil {
		return nil, err
	}
	client := redis.NewClient(opts)
	_, err = client.Ping().Result()
	return client, err
}

func NewRedisRepository(redisUrl string) (ports.OrderRepository, error)  {
	repo := &RedisRepository{}
	client, err := newRedisClient(redisUrl)
	if err != nil {
		return nil, errors.Wrap(err, "repository.NewRedisRepository")
	}
	repo.client = client
	return repo, nil
}

func (r *RedisRepository) generateKey(code string) string  {
	return fmt.Sprintf("order:%s", code)
}
func (r *RedisRepository) FindByCode(code string) (*domain.Order, error) {
	order := &domain.Order{}
	key := r.generateKey(code)
	data, err := r.client.HGetAll(key).Result()
	if err != nil {
		return nil, errors.Wrap(err, "repository.Redirect.Find")
	}
	if len(data) == 0 {
		return nil, errors.Wrap(service.ErrProductNotFound, "repository.Redirect.FindNotFound")
	}
	createdAt, err := strconv.ParseInt(data["created_at"], 10, 64)
	if err != nil {
		return nil, errors.Wrap(service.ErrOrderNotFound, "repository.Redirect.FindErrorParse")
	}
	order.ProductCode = data["product_code"]
	order.Code = data["code"]
	order.PriceTotal, _ = strconv.ParseInt(data["price_total"], 10, 64)
	order.CreatedAt = createdAt
	return order, nil
}

func (r *RedisRepository) Save(order *domain.Order) error  {
	key := r.generateKey(order.Code)
	data := map[string]interface{}{
		"product_code":       order.ProductCode,
		"code":        order.Code,
		"price_total":        order.PriceTotal,
		"created_at": order.CreatedAt,
	}
	_, err := r.client.HMSet(key, data).Result()
	if err != nil {
		return errors.Wrap(err, "repository.Redirect.Store")
	}
	return nil
}
