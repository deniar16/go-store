package main

import (
	"fmt"
	"github.com/deniarianto1606/go-store/controller/order"
	"github.com/deniarianto1606/go-store/controller/product"
	orderGateway "github.com/deniarianto1606/go-store/gateway/order"
	serviceOrder "github.com/deniarianto1606/go-store/order/service"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func initializeRouter(uc *appUseCase, repo *appRepo) {
	productHandler := product.NewHandler(uc.findByCode, uc.save)
	so := serviceOrder.NewOrderService(orderGateway.NewOrderGateway(repo.productMongo, repo.productMongo,
		repo.orderRedis, repo.orderMongo))
	orderHandler := order.NewHandler(so)

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/{code}", productHandler.FindByCode)
	r.Post("/", productHandler.CreateProduct)
	r.Get("/order/{code}", orderHandler.FindByCode)
	r.Post("/order", orderHandler.CreateOrder)

	listenAndServe(r)
}

func listenAndServe(r *chi.Mux) {
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
