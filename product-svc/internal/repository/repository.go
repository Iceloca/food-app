package repository

import (
	"context"
	"errors"
	"github.com/r1nb0/food-app/product-svc/internal/domain/models"
)

var (
	ErrNotFound      = errors.New("product not found")
	ErrAlreadyExists = errors.New("product already exists")
)

type ProductRepository interface {
	CreateProduct(ctx context.Context, dto models.CreateProductDTO) (string, error)
	UpdateProduct(ctx context.Context, dto models.UpdateProductDTO) (bool, error)
	DeleteProduct(ctx context.Context, id string) (bool, error)
	GetProductByID(ctx context.Context, id string) (models.Product, error)
	GetAllProducts(ctx context.Context) ([]models.Product, error)
	GetProductsByCategory(ctx context.Context, categoryName string) ([]models.Product, error)
}
