package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/r1nb0/food-app/api-gateway/internal/auth/client"
	"github.com/r1nb0/food-app/api-gateway/internal/config"
)

func RegisterAuth(r *gin.RouterGroup, cfg *config.Config) {
	authClient := client.NewAuthClient(cfg)
	auth := r.Group("/auth")
	{
		auth.POST("/register", authClient.Register)
		auth.POST("/login", authClient.Login)
	}
}
