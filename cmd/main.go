package main

import (
	"fmt"
	c "github.com/deniarianto1606/go-store/controller/product"
	"github.com/deniarianto1606/go-store/product"
	m "github.com/deniarianto1606/go-store/repository/mongo/product"
	"github.com/deniarianto1606/go-store/repository/redis"
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
	service := product.NewProductService(product.NewProductGateway(repoRedis, repoMongo))
	handler := c.NewHandler(service)

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/{code}", handler.FindByCode)
	r.Post("/", handler.CreateProduct)

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


func repoMongo() product.ProductRepository {

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

func repoRedis() product.ProductRepository {
	log.Printf("get mongodb")
	redisUrl := "redis://localhost:6379" //os.Getenv("MONGO_URL")
	repo, err := redis.NewRedisRepository(redisUrl)
	if err != nil {
		log.Fatal(err)
	}
	return repo
}