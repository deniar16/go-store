package main

import (
	productGateway "github.com/deniarianto1606/go-store/gateway/product"
	serviceProduct "github.com/deniarianto1606/go-store/product/service"
	"github.com/deniarianto1606/go-store/product/service/findbycode"
)

func initializeGateway(repo *appRepo, gw *appGateway) {
	productGw := productGateway.NewProductGateway(repo.productRedis, repo.productMongo)
	gw.product = productGw
}

func initializeUseCase(gw *appGateway, uc *appUseCase) {
	service := serviceProduct.NewProductService(gw.product)
	productFindByCodeUseCase := findbycode.NewUseCase(gw.product)
	uc.findByCode = productFindByCodeUseCase
	uc.service = service
}
