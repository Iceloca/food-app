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
	return &productService{
		productRepository: productRepository,
	}
}

func (s *productService) Create(ctx context.Context, product models.ProductCreate) (int64, error) {
	return s.productRepository.Create(ctx, product)
}

func (s *productService) Update(ctx context.Context, product models.Product) error {
	return s.productRepository.Update(ctx, product)
}

func (s *productService) Delete(ctx context.Context, id int64) error {
	return s.productRepository.Delete(ctx, id)
}

func (s *productService) GetAll(ctx context.Context) ([]models.Product, error) {
	return s.productRepository.GetAll(ctx)
}

func (s *productService) GetByID(ctx context.Context, id int64) (models.Product, error) {
	return s.productRepository.GetByID(ctx, id)
}

func (s *productService) GetByCategory(ctx context.Context, categoryID int64) ([]models.Product, error) {
	return s.productRepository.GetByCategory(ctx, categoryID)
}

func (s *productService) GetDailyRecs(ctx context.Context) ([]models.Product, error) {
	return s.productRepository.GetDailyRecs(ctx)
}
