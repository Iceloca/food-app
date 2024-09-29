package app

import (
	"github.com/r1nb0/food-app/product-svc/internal/app/app"
	"github.com/r1nb0/food-app/product-svc/internal/config"
	"github.com/r1nb0/food-app/product-svc/internal/lib/mongo"
	productRepo "github.com/r1nb0/food-app/product-svc/internal/repository/mongo"
	"github.com/r1nb0/food-app/product-svc/internal/service/product"
	"log"
)

type App struct {
	GRPCServer *app.App
}

func New(cfg *config.Config) *App {
	client, err := mongo.InitClient(cfg)
	if err != nil {
		log.Fatalf("error of init client: %v", err)
	}
	db := client.Database(cfg.MongoDB.DBName)

	productRepository := productRepo.NewProductRepository(db, cfg.MongoDB.Collection)
	productService := product.NewProductService(productRepository)
	grpcApp := app.New(productService, cfg.GRPC.Port)

	return &App{
		GRPCServer: grpcApp,
	}
}
