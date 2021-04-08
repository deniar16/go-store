package main

import (
	productGateway "github.com/deniarianto1606/go-store/gateway/product"
	po "github.com/deniarianto1606/go-store/order/ports"
	pp "github.com/deniarianto1606/go-store/product/ports"
	"github.com/deniarianto1606/go-store/product/ports/findbycode"
	"github.com/deniarianto1606/go-store/product/ports/save"
	"go.mongodb.org/mongo-driver/mongo"
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
	initializeRouter(ctx.useCase, ctx.repo)
}
