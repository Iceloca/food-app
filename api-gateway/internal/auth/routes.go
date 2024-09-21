package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/r1nb0/food-app/api-gateway/internal/auth/routes"
	"github.com/r1nb0/food-app/api-gateway/internal/config"
)

func RegisterRoutes(r *gin.Engine, cfg *config.Config) {
	svc := &ServiceClient{
		Client: InitServiceClient(cfg),
	}

	auth := r.Group("/auth")
	auth.POST("/register", svc.Register)
	auth.POST("/login", svc.Login)
}

func (svc *ServiceClient) Register(c *gin.Context) {
	routes.Register(c, svc.Client)
}

func (svc *ServiceClient) Login(c *gin.Context) {
	routes.Login(c, svc.Client)
}
