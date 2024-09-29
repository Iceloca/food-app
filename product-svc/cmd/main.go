package main

import (
	"github.com/joho/godotenv"
	"github.com/r1nb0/food-app/product-svc/internal/app"
	"github.com/r1nb0/food-app/product-svc/internal/config"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg := config.GetConfig()

	application := app.New(cfg)

	go func() {
		application.GRPCServer.Run()
	}()

	stop := make(chan os.Signal, 1)

	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<-stop

	application.GRPCServer.Stop()
}
