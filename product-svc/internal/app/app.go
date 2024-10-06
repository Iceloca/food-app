package app

import (
	"fmt"
	"github.com/r1nb0/food-app/pkg/database/postgres"
	"github.com/r1nb0/food-app/product-svc/internal/app/grpc"
	"github.com/r1nb0/food-app/product-svc/internal/config"
	categoryRepo "github.com/r1nb0/food-app/product-svc/internal/repository/postgres/category"
	productRepo "github.com/r1nb0/food-app/product-svc/internal/repository/postgres/product"
	categoryServ "github.com/r1nb0/food-app/product-svc/internal/service/category"
	productServ "github.com/r1nb0/food-app/product-svc/internal/service/product"
	sloglogrus "github.com/samber/slog-logrus/v2"
	"github.com/sirupsen/logrus"
	"log"
	"log/slog"
)

type App struct {
	GRPCServer *grpc.App
}

func New(cfg *config.Config) *App {
	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=%s password=%s",
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.Username,
		cfg.Postgres.DBName,
		cfg.Postgres.SSLMode,
		cfg.Postgres.Password,
	)

	logger := slog.New(
		sloglogrus.Option{
			Level:  slog.LevelDebug,
			Logger: logrus.New()}.NewLogrusHandler(),
	)

	db, err := postgres.InitDB(dataSourceName)
	if err != nil {
		log.Fatalf("error connecting to database: %v", err)
	}

	categoryRepository := categoryRepo.NewCategoryRepository(db)
	categoryService := categoryServ.NewCategoryService(categoryRepository, logger)
	productRepository := productRepo.NewProductRepository(db)
	productService := productServ.NewProductService(productRepository, categoryRepository, logger)
	gRPCServer := grpc.New(categoryService, productService, cfg.GRPC.Port)

	return &App{
		GRPCServer: gRPCServer,
	}
}
