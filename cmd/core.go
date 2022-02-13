package main

import (
	productGateway "github.com/deniarianto1606/go-store/gateway/product"
	"github.com/deniarianto1606/go-store/product/service/findbycode"
	"github.com/deniarianto1606/go-store/product/service/save"
)

func initializeGateway(repo *appRepo, gw *appGateway) {
	gw.product = productGateway.NewProductGateway(repo.productRedis, repo.productMongo)
}

func initializeUseCase(gw *appGateway, uc *appUseCase) {
	productFindByCodeUseCase := findbycode.NewUseCase(gw.product)
	productSaveUseCase := save.NewUseCase(gw.product)
	uc.findByCode = productFindByCodeUseCase
	uc.save = productSaveUseCase
}
