package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/r1nb0/food-app/api-gateway/internal/auth"
	"github.com/r1nb0/food-app/api-gateway/internal/config"
	"log"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	cfg := config.GetConfig()

	r := gin.Default()

	auth.RegisterRoutes(r, cfg)

	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatal(err)
	}
}
