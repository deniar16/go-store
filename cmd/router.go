package main

import (
	"github.com/deniarianto1606/go-store/controller/order"
	"github.com/deniarianto1606/go-store/controller/product"
	orderGateway "github.com/deniarianto1606/go-store/gateway/order"
	serviceOrder "github.com/deniarianto1606/go-store/order/service"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func initializeRouter(uc *appUseCase, repo *appRepo) *chi.Mux {
	handler := product.NewHandler(uc.findByCode, uc.save)
	so := serviceOrder.NewOrderService(orderGateway.NewOrderGateway(repo.productMongo, repo.productMongo,
		repo.orderRedis, repo.orderMongo))
	handlerOrder := order.NewHandler(so)

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/{code}", handler.FindByCode)
	r.Post("/", handler.CreateProduct)

	r.Get("/order/{code}", handlerOrder.FindByCode)
	r.Post("/order", handlerOrder.CreateOrder)
	return r
}
