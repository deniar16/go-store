package order


import (
	"context"
	"github.com/deniarianto1606/go-store/order/domain"
	"github.com/deniarianto1606/go-store/order/ports"
	"github.com/deniarianto1606/go-store/order/service"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

type MongoRepository struct {
	client *mongo.Client
	database string
	timeout time.Duration
}

func newMongoClient(mongoUrl string, mongoTimeout int) (*mongo.Client, error){
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(mongoTimeout)*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoUrl))
	if err != nil {
		return nil, err
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}
	return client, nil
}

func NewMongoRepository(mongoUrl, mongoDb string, mongoTimeout int) (ports.OrderRepository, error)  {
	repo := &MongoRepository{
		timeout: time.Duration(mongoTimeout) * time.Second,
		database: mongoDb,
	}
	client, err := newMongoClient(mongoUrl, mongoTimeout)
	if err != nil {
		return nil, errors.Wrap(err, "repository.NewMongoRepository")
	}
	repo.client = client
	return repo, nil
}

func (r *MongoRepository) FindByCode(code string) (*domain.Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	redirect := &domain.Order{}
	collection := r.client.Database(r.database).Collection("order")
	filter := bson.M{"code": code}
	err := collection.FindOne(ctx, filter).Decode(&redirect)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.Wrap(service.ErrOrderNotFound, "notFound.ErrNoDocument")
		}
		return nil, errors.Wrap(err, "repository.Order.FindError")
	}
	return redirect, nil
}

func (r *MongoRepository) Save(order *domain.Order) error  {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection("order")
	_, err:= collection.InsertOne(
		ctx,
		bson.M{
			"price_total": order.PriceTotal,
			"code": order.Code,
			"product_code": order.ProductCode,
			"created_at": order.CreatedAt,
		},
	)
	if err != nil {
		return errors.Wrap(err, "repository.Product.Store Error")
	}
	return nil
}
