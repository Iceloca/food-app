package app

import (
	"fmt"
	"github.com/r1nb0/food-app/cart-svc/internal/app/grpc"
	"github.com/r1nb0/food-app/cart-svc/internal/config"
	cartRepo "github.com/r1nb0/food-app/cart-svc/internal/repository/mongodb"
	"github.com/r1nb0/food-app/cart-svc/internal/service/cart"
	"github.com/r1nb0/food-app/pkg/database/mongodb"
	sloglogrus "github.com/samber/slog-logrus/v2"
	"github.com/sirupsen/logrus"
	"log"
	"log/slog"
)

type App struct {
	GRPCServer *grpc.App
}

func New(cfg *config.Config) *App {
	mongoURI := fmt.Sprintf("mongodb://%s:%s/", cfg.MongoDB.Host, cfg.MongoDB.Port)

	logger := slog.New(
		sloglogrus.Option{
			Level:  slog.LevelDebug,
			Logger: logrus.New()}.NewLogrusHandler(),
	)

	client, err := mongodb.InitClient(mongoURI)

	if err != nil {
		log.Fatalf("failed to connet to the mongo client: %s", err.Error())
	}

	collection := client.Database(cfg.MongoDB.DBName).Collection(cfg.MongoDB.Collection)

	cartRepository := cartRepo.NewCartRepository(collection)
	cartService := cart.NewCartService(cartRepository, logger)
	grpcApp := grpc.New(cartService, cfg.GRPC.Port)

	return &App{
		GRPCServer: grpcApp,
	}
}
