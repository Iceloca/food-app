package app

import (
	"github.com/r1nb0/food-app/product-svc/internal/app/grpc"
	"github.com/r1nb0/food-app/product-svc/internal/config"
	"github.com/r1nb0/food-app/product-svc/internal/lib/postgres"
	categoryRepo "github.com/r1nb0/food-app/product-svc/internal/repository/postgres/category"
	productRepo "github.com/r1nb0/food-app/product-svc/internal/repository/postgres/product"
	categoryServ "github.com/r1nb0/food-app/product-svc/internal/service/category"
	productServ "github.com/r1nb0/food-app/product-svc/internal/service/product"
	"log"
)

type App struct {
	GRPCServer *grpc.App
}

func New(cfg *config.Config) *App {
	db, err := postgres.InitDB(cfg)
	if err != nil {
		log.Fatalf("error connecting to database: %v", err)
	}

	categoryRepository := categoryRepo.NewCategoryRepository(db)
	categoryService := categoryServ.NewCategoryService(categoryRepository)
	productRepository := productRepo.NewProductRepository(db)
	productService := productServ.NewProductService(productRepository)
	gRPCServer := grpc.New(categoryService, productService, cfg.GRPC.Port)

	return &App{
		GRPCServer: gRPCServer,
	}
}
