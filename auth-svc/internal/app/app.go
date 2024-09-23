package app

import (
	"github.com/r1nb0/food-app/auth-svc/internal/app/grpc"
	"github.com/r1nb0/food-app/auth-svc/internal/config"
	"github.com/r1nb0/food-app/auth-svc/internal/lib/postgres"
	userRepo "github.com/r1nb0/food-app/auth-svc/internal/repository/postgres"
	"github.com/r1nb0/food-app/auth-svc/internal/service/auth"
	"log"
)

type App struct {
	GRPCServer *grpc.App
}

func New(cfg *config.Config) *App {
	db, err := postgres.InitDB(cfg)

	if err != nil {
		log.Fatalf("error connecting to the database: %s", err.Error())
	}

	userRepository := userRepo.NewUserRepository(db)

	authService := auth.NewAuthService(userRepository, cfg.TokenTTL)

	grpcApp := grpc.New(authService, cfg.GRPC.Port)

	return &App{
		GRPCServer: grpcApp,
	}
}
