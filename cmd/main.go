package main

import (
	"fmt"
	order "github.com/deniarianto1606/go-store/controller/order"
	product "github.com/deniarianto1606/go-store/controller/product"
	orderGateway "github.com/deniarianto1606/go-store/gateway/order"
	productGateway "github.com/deniarianto1606/go-store/gateway/product"
	po "github.com/deniarianto1606/go-store/order/ports"
	serviceOrder "github.com/deniarianto1606/go-store/order/service"
	pp "github.com/deniarianto1606/go-store/product/ports"
	serviceProduct "github.com/deniarianto1606/go-store/product/service"
	om "github.com/deniarianto1606/go-store/repository/mongo/order"
	m "github.com/deniarianto1606/go-store/repository/mongo/product"
	redisProduct "github.com/deniarianto1606/go-store/repository/redis/product"
	redisOrder "github.com/deniarianto1606/go-store/repository/redis/order"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	repoMongo := repoMongo()
	repoRedis := repoRedis()

	orderMongo := orderMongo()
	orderRedis := orderRedis()

	service := serviceProduct.NewProductService(productGateway.NewProductGateway(repoRedis, repoMongo))
	handler := product.NewHandler(service)

	serviceOrder := serviceOrder.NewOrderService(orderGateway.NewOrderGateway(repoMongo, repoMongo, orderRedis, orderMongo))
	handlerOrder := order.NewHandler(serviceOrder)



	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/{code}", handler.FindByCode)
	r.Post("/", handler.CreateProduct)

	r.Get("/order/{code}", handlerOrder.FindByCode)
	r.Post("/order", handlerOrder.CreateOrder)

	errs := make(chan error, 2)
	go func() {
		fmt.Println("Listening on port :50001")
		errs <- http.ListenAndServe(httpPort(), r)
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()
	fmt.Printf("Terminated %s", <-errs)
}

func httpPort() string {
	port := "50001"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	return fmt.Sprintf(":%s", port)
}


func repoMongo() pp.ProductRepository {

	log.Printf("get mongodb")
	mongoUrl := "mongodb://localhost/" //os.Getenv("MONGO_URL")
	mongodb := "go-store" //os.Getenv("MONGO_DB")
	mongoTimeout := 30 //strconv.Atoi(os.Getenv("MONGO_TIMEOUT"))
	repo, err := m.NewMongoRepository(mongoUrl, mongodb, mongoTimeout)
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
	mongodb := "go-store" //os.Getenv("MONGO_DB")
	mongoTimeout := 30 //strconv.Atoi(os.Getenv("MONGO_TIMEOUT"))
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