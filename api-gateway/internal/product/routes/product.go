package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/r1nb0/food-app/api-gateway/internal/config"
	"github.com/r1nb0/food-app/api-gateway/internal/product/client"
)

func RegisterProduct(router *gin.RouterGroup, cfg *config.Config) {
	productClient := client.NewProductClient(cfg)

	product := router.Group("/product")
	{
		product.GET("/", productClient.GetAll)
		product.GET("/:id", productClient.GetByID)
		product.GET("/category/:id", productClient.GetByCategory)
		product.GET("/recs", productClient.GetDailyRecs)
		product.POST("/", productClient.CreateProduct)
		product.PUT("/", productClient.UpdateProduct)
		product.DELETE("/:id", productClient.DeleteProduct)
	}
}
