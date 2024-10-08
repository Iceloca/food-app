package main

import (
	"github.com/joho/godotenv"
	"github.com/r1nb0/food-app/cart-svc/internal/app"
	"github.com/r1nb0/food-app/cart-svc/internal/config"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file : %s", err.Error())
	}

	cfg := config.GetConfig()

	application := app.New(cfg)

	go func() {
		application.GRPCServer.Run()
	}()

	stop := make(chan os.Signal, 1)

	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop

	application.GRPCServer.Stop()
}
