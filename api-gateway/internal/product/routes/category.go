package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/r1nb0/food-app/api-gateway/internal/config"
	"github.com/r1nb0/food-app/api-gateway/internal/product/client"
)

func RegisterCategory(router *gin.RouterGroup, cfg *config.Config) {
	categoryClient := client.NewCategoryClient(cfg)

	category := router.Group("/category")
	{
		category.GET("/", categoryClient.GetAll)
		category.GET("/:id", categoryClient.GetByID)
		category.POST("/", categoryClient.CreateCategory)
		category.DELETE("/:id", categoryClient.DeleteCategory)
		category.PUT("/", categoryClient.UpdateCategory)
	}
}
