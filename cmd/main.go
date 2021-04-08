package main

import (
	"fmt"
	productGateway "github.com/deniarianto1606/go-store/gateway/product"
	po "github.com/deniarianto1606/go-store/order/ports"
	pp "github.com/deniarianto1606/go-store/product/ports"
	"github.com/deniarianto1606/go-store/product/ports/findbycode"
	"github.com/deniarianto1606/go-store/product/ports/save"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type appContext struct {
	useCase *appUseCase
	gateway *appGateway
	repo    *appRepo
	db      *appDB
}

type appDB struct {
	mongo *mongo.Database
}

type appUseCase struct {
	findByCode findbycode.UseCase
	save       save.UseCase
}

type appGateway struct {
	product productGateway.Gateway
}

type appRepo struct {
	productRedis pp.ProductRepository
	productMongo pp.ProductRepository
	orderMongo   po.OrderRepository
	orderRedis   po.OrderRepository
}

func NewContext() *appContext {
	return &appContext{
		useCase: &appUseCase{},
		gateway: &appGateway{},
		repo:    &appRepo{},
		db:      &appDB{},
	}
}

func main() {
	ctx := NewContext()
	initializeDB(ctx.db)
	initializeRepo(ctx.db, ctx.repo)
	initializeGateway(ctx.repo, ctx.gateway)
	initializeUseCase(ctx.gateway, ctx.useCase)
	r := initializeRouter(ctx.useCase, ctx.repo)

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
