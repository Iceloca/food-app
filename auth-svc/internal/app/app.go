package app

import (
	"fmt"
	"github.com/r1nb0/food-app/auth-svc/internal/app/grpc"
	"github.com/r1nb0/food-app/auth-svc/internal/config"
	userRepo "github.com/r1nb0/food-app/auth-svc/internal/repository/postgres"
	"github.com/r1nb0/food-app/auth-svc/internal/service/auth"
	"github.com/r1nb0/food-app/pkg/database/postgres"
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
		log.Fatalf("error connecting to the database: %s", err.Error())
	}

	userRepository := userRepo.NewUserRepository(db)
	authService := auth.NewAuthService(userRepository, logger, cfg.TokenTTL)
	grpcApp := grpc.New(authService, cfg.GRPC.Port)

	return &App{
		GRPCServer: grpcApp,
	}
}
