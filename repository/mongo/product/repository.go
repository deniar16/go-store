package product

import (
	"context"
	"github.com/deniarianto1606/go-store/product/domain"
	"github.com/deniarianto1606/go-store/product/ports"
	"github.com/deniarianto1606/go-store/product/service"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

type MongoRepository struct {
	client   *mongo.Client
	database string
	timeout  time.Duration
	DB       *mongo.Database
}

func newMongoClient(mongoUrl string, mongoTimeout int) (*mongo.Client, error) {
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

func NewMongoRepository(mongoUrl, mongoDb string, mongoTimeout int, mongoDB *mongo.Database) (ports.ProductRepository, error) {
	repo := &MongoRepository{
		timeout:  time.Duration(mongoTimeout) * time.Second,
		database: mongoDb,
		DB:       mongoDB,
	}
	client, err := newMongoClient(mongoUrl, mongoTimeout)
	if err != nil {
		return nil, errors.Wrap(err, "repository.NewMongoRepository")
	}
	repo.client = client
	return repo, nil
}

func (r *MongoRepository) FindByCode(code string) (*domain.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	redirect := &domain.Product{}
	//collection := r.client.Database(r.database).Collection("product")
	collection := r.DB.Collection("product")
	filter := bson.M{"code": code}
	err := collection.FindOne(ctx, filter).Decode(&redirect)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.Wrap(service.ErrProductInvalid, "notFound.ErrNoDocument")
		}
		return nil, errors.Wrap(err, "repository.Product.FindError")
	}
	return redirect, nil
}

func (r *MongoRepository) Save(product *domain.Product) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection("product")
	_, err := collection.InsertOne(
		ctx,
		bson.M{
			"code":       product.Code,
			"name":       product.Name,
			"desc":       product.Desc,
			"created_at": product.CreatedAt,
		},
	)
	if err != nil {
		return errors.Wrap(err, "repository.Product.Store Error")
	}
	return nil
}
