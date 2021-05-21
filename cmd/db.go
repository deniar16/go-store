package main

import (
	po "github.com/deniarianto1606/go-store/order/ports"
	pp "github.com/deniarianto1606/go-store/product/ports"
	om "github.com/deniarianto1606/go-store/repository/mongo/order"
	m "github.com/deniarianto1606/go-store/repository/mongo/product"
	redisOrder "github.com/deniarianto1606/go-store/repository/redis/order"
	redisProduct "github.com/deniarianto1606/go-store/repository/redis/product"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func initializeDB(db *appDB) {
	mongoDB, _ := newMongoDatabase()
	db.mongo = mongoDB
}

func initializeRepo(db *appDB, repo *appRepo) {
	repoMongo := repoMongo(db.mongo)
	repoRedis := repoRedis()
	orderMongo := orderMongo()
	orderRedis := orderRedis()
	repo.productMongo = repoMongo
	repo.productRedis = repoRedis
	repo.orderMongo = orderMongo
	repo.orderRedis = orderRedis
}

func repoMongo(mongoDB *mongo.Database) pp.ProductRepository {
	log.Printf("get mongodb")
	mongoUrl := "mongodb://localhost/" //os.Getenv("MONGO_URL")
	mongodb := "go-store"              //os.Getenv("MONGO_DB")
	mongoTimeout := 30                 //strconv.Atoi(os.Getenv("MONGO_TIMEOUT"))
	repo, err := m.NewMongoRepository(mongoUrl, mongodb, mongoTimeout, mongoDB)
	if err != nil {
		log.Fatal(err)
	}
	return repo
}

func repoRedis() pp.ProductRepository {
	log.Printf("get mongodb")
	redisUrl := "redis://localhost:6379" //os.Getenv("MONGO_URL")
	repo, err := redisProduct.NewRedisRepository(redisUrl)
	if err != nil {
		log.Fatal(err)
	}
	return repo
}

func orderMongo() po.OrderRepository {
	log.Printf("get mongodb")
	mongoUrl := "mongodb://localhost/" //os.Getenv("MONGO_URL")
	mongodb := "go-store"              //os.Getenv("MONGO_DB")
	mongoTimeout := 30                 //strconv.Atoi(os.Getenv("MONGO_TIMEOUT"))
	repo, err := om.NewMongoRepository(mongoUrl, mongodb, mongoTimeout)
	if err != nil {
		log.Fatal(err)
	}
	return repo
}

func orderRedis() po.OrderRepository {
	log.Printf("get mongodb")
	redisUrl := "redis://localhost:6379" //os.Getenv("MONGO_URL")
	repo, err := redisOrder.NewRedisRepository(redisUrl)
	if err != nil {
		log.Fatal(err)
	}
	return repo
}
