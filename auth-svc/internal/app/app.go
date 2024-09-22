package app

import (
	"github.com/r1nb0/food-app/auth-svc/internal/app/grpc"
	"github.com/r1nb0/food-app/auth-svc/internal/config"
	"github.com/r1nb0/food-app/auth-svc/internal/lib/postgres"
	userRepo "github.com/r1nb0/food-app/auth-svc/internal/repository/postgres"
	"github.com/r1nb0/food-app/auth-svc/internal/service"
	"log"
)

type App struct {
	GRPCServer *grpc.App
}

func New(cfg *config.Config) *App {
	db, err := postgres.InitDB(cfg)

	if err != nil {
		log.Fatal(err)
	}

	userRepository := userRepo.NewUserRepository(db)

	authService := service.NewAuthService(userRepository, cfg.TokenTTL)

	grpcApp := grpc.New(authService, cfg.GRPC.Port)

	return &App{
		GRPCServer: grpcApp,
	}
}
