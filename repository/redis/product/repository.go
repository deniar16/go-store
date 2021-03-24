package product

import (
	"fmt"
	"github.com/deniarianto1606/go-store/product/domain"
	"github.com/deniarianto1606/go-store/product/ports"
	"github.com/deniarianto1606/go-store/product/service"
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

func NewRedisRepository(redisUrl string) (ports.ProductRepository, error)  {
	repo := &RedisRepository{}
	client, err := newRedisClient(redisUrl)
	if err != nil {
		return nil, errors.Wrap(err, "repository.NewRedisRepository")
	}
	repo.client = client
	return repo, nil
}

func (r *RedisRepository) generateKey(code string) string  {
	return fmt.Sprintf("redirect:%s", code)
}
func (r *RedisRepository) FindByCode(code string) (*domain.Product, error) {
	redirect := &domain.Product{}
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
		return nil, errors.Wrap(service.ErrProductInvalid, "repository.Redirect.FindErrorParse")
	}
	redirect.Code = data["code"]
	redirect.Name = data["name"]
	redirect.Desc = data["desc"]
	redirect.CreatedAt = createdAt
	return redirect, nil
}

func (r *RedisRepository) Save(redirect *domain.Product) error  {
	key := r.generateKey(redirect.Code)
	data := map[string]interface{}{
		"code":       redirect.Code,
		"name":        redirect.Name,
		"desc":        redirect.Desc,
		"created_at": redirect.CreatedAt,
	}
	_, err := r.client.HMSet(key, data).Result()
	if err != nil {
		return errors.Wrap(err, "repository.Redirect.Store")
	}
	return nil
}