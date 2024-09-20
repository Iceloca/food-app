package main

import (
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/r1nb0/food-app/auth-svc/internal/app"
	"github.com/r1nb0/food-app/auth-svc/internal/config"
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

	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop

	application.GRPCServer.Stop()
}
