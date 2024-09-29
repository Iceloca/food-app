package category

import (
	"context"
	"github.com/r1nb0/food-app/product-svc/internal/domain/models"
	"github.com/r1nb0/food-app/product-svc/internal/repository"
	"github.com/r1nb0/food-app/product-svc/internal/service"
)

type categoryService struct {
	categoryRepository repository.CategoryRepository
}

func NewCategoryService(categoryRepository repository.CategoryRepository) service.CategoryService {
	return &categoryService{
		categoryRepository: categoryRepository,
	}
}

func (s *categoryService) Create(ctx context.Context, category models.CategoryCreate) (int64, error) {
	return s.categoryRepository.Create(ctx, category)
}

func (s *categoryService) Update(ctx context.Context, category models.Category) error {
	return s.categoryRepository.Update(ctx, category)
}

func (s *categoryService) Delete(ctx context.Context, id int64) error {
	return s.categoryRepository.Delete(ctx, id)
}

func (s *categoryService) GetAll(ctx context.Context) ([]models.Category, error) {
	return s.categoryRepository.GetAll(ctx)
}

func (s *categoryService) GetByID(ctx context.Context, id int64) (models.Category, error) {
	return s.categoryRepository.GetByID(ctx, id)
}
