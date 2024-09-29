package product

import (
	"context"
	"github.com/r1nb0/food-app/product-svc/internal/domain/models"
	"github.com/r1nb0/food-app/product-svc/internal/repository"
	"github.com/r1nb0/food-app/product-svc/internal/service"
)

type productService struct {
	productRepository repository.ProductRepository
}

func NewProductService(productRepository repository.ProductRepository) service.ProductService {
	return &productService{productRepository: productRepository}
}

func (s *productService) CreateProduct(ctx context.Context, dto models.CreateProductDTO) (string, error) {
	return s.productRepository.CreateProduct(ctx, dto)
}

func (s *productService) UpdateProduct(ctx context.Context, dto models.UpdateProductDTO) (bool, error) {
	return s.productRepository.UpdateProduct(ctx, dto)
}

func (s *productService) DeleteProduct(ctx context.Context, id string) (bool, error) {
	return s.productRepository.DeleteProduct(ctx, id)
}

func (s *productService) GetProductByID(ctx context.Context, id string) (models.Product, error) {
	return s.productRepository.GetProductByID(ctx, id)
}

func (s *productService) GetAllProducts(ctx context.Context) ([]models.Product, error) {
	return s.productRepository.GetAllProducts(ctx)
}

func (s *productService) GetProductsByCategory(ctx context.Context, categoryName string) ([]models.Product, error) {
	return s.productRepository.GetProductsByCategory(ctx, categoryName)
}

func (s *productService) GetAllCategories(ctx context.Context) ([]models.CategoryProduct, error) {
	return s.productRepository.GetAllCategories(ctx)
}

func (s *productService) GetBestProducts(ctx context.Context) ([]models.Product, error) {
	return s.productRepository.GetBestProducts(ctx)
}
