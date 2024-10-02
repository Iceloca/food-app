package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	authRoutes "github.com/r1nb0/food-app/api-gateway/internal/auth/routes"
	"github.com/r1nb0/food-app/api-gateway/internal/config"
	productRoutes "github.com/r1nb0/food-app/api-gateway/internal/product/routes"
	"log"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	cfg := config.GetConfig()

	r := gin.Default()

	api := r.Group("/api")
	v1 := api.Group("/v1")

	authRoutes.RegisterAuth(v1, cfg)
	productRoutes.RegisterProduct(v1, cfg)
	productRoutes.RegisterCategory(v1, cfg)

	if err := r.Run("0.0.0.0:" + cfg.Port); err != nil {
		log.Fatal(err)
	}
}
