package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/r1nb0/food-app/product-svc/internal/cfg"
	"log"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	fmt.Println(cfg.GetConfig())
}
