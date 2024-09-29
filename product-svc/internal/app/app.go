package app

import (
	"github.com/r1nb0/food-app/product-svc/internal/app/grpc"
	"github.com/r1nb0/food-app/product-svc/internal/config"
	"github.com/r1nb0/food-app/product-svc/internal/lib/postgres"
	categoryRepo "github.com/r1nb0/food-app/product-svc/internal/repository/postgres/category"
	categoryServ "github.com/r1nb0/food-app/product-svc/internal/service/category"
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
	gRPCServer := grpc.New(categoryService, cfg.GRPC.Port)

	return &App{
		GRPCServer: gRPCServer,
	}
}
