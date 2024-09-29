package service

import (
	"context"
	"github.com/r1nb0/food-app/product-svc/internal/domain/models"
)

type ProductService interface {
	CreateProduct(ctx context.Context, dto models.CreateProductDTO) (string, error)
	UpdateProduct(ctx context.Context, dto models.UpdateProductDTO) (bool, error)
	DeleteProduct(ctx context.Context, id string) (bool, error)
	GetProductByID(ctx context.Context, id string) (models.Product, error)
	GetAllProducts(ctx context.Context) ([]models.Product, error)
	GetProductsByCategory(ctx context.Context, categoryName string) ([]models.Product, error)
	GetAllCategories(ctx context.Context) ([]models.CategoryProduct, error)
	GetBestProducts(ctx context.Context) ([]models.Product, error)
}
