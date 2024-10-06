package category

import (
	"context"
	"errors"
	"github.com/r1nb0/food-app/pkg/database"
	"github.com/r1nb0/food-app/product-svc/internal/domain/models"
	"github.com/r1nb0/food-app/product-svc/internal/repository"
	"github.com/r1nb0/food-app/product-svc/internal/service"
	"log/slog"
)

type categoryService struct {
	categoryRepository repository.CategoryRepository
	logger             *slog.Logger
}

func NewCategoryService(
	categoryRepository repository.CategoryRepository,
	logger *slog.Logger,
) service.CategoryService {
	return &categoryService{
		categoryRepository: categoryRepository,
		logger:             logger,
	}
}

func (s *categoryService) Create(
	ctx context.Context,
	category models.CategoryCreate,
) (int64, error) {
	const op = "categoryService.Create"

	log := s.logger.With(
		slog.String("operation", op),
		slog.Any("data", category),
	)

	log.Info("attempting to create category")

	id, err := s.categoryRepository.Create(ctx, category)
	if err != nil {
		log.Error(
			"failed to create category",
			slog.String("error", err.Error()),
		)
		return 0, err
	}

	return id, nil
}

func (s *categoryService) Update(
	ctx context.Context,
	category models.Category,
) error {
	const op = "categoryService.Update"

	log := s.logger.With(
		slog.String("operation", op),
		slog.Any("data", category),
	)

	log.Info("attempting to update category")

	if err := s.categoryRepository.Update(ctx, category); err != nil {
		log.Error(
			"failed to update category",
			slog.String("error", err.Error()),
		)
		return err
	}

	return nil
}

func (s *categoryService) Delete(ctx context.Context, id int64) error {
	const op = "categoryService.Delete"

	log := s.logger.With(
		slog.String("operation", op),
		slog.Any("data", id),
	)

	log.Info("attempting to delete category")

	if err := s.categoryRepository.Delete(ctx, id); err != nil {
		log.Error(
			"failed to delete category",
			slog.String("error", err.Error()),
		)
		return err
	}

	return nil
}

func (s *categoryService) GetAll(ctx context.Context) ([]models.Category, error) {
	const op = "categoryService.GetAll"

	log := s.logger.With(slog.String("operation", op))

	log.Info("attempting to get all categories")

	categories, err := s.categoryRepository.GetAll(ctx)
	if err != nil {
		if errors.Is(err, database.ErrNotFound) {
			log.Error(
				"categories not found",
				slog.String("error", err.Error()),
			)
		} else {
			log.Error(
				"failed to get all categories",
				slog.String("error", err.Error()),
			)
		}
		return nil, err
	}

	return categories, nil
}

func (s *categoryService) GetByID(ctx context.Context, id int64) (models.Category, error) {
	const op = "categoryService.GetByID"

	log := s.logger.With(
		slog.String("operation", op),
		slog.Any("data", id),
	)

	log.Info("attempting to get category")

	category, err := s.categoryRepository.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, database.ErrNotFound) {
			log.Error(
				"category not found",
				slog.String("error", err.Error()),
			)
		} else {
			log.Error(
				"failed to get category",
				slog.String("error", err.Error()),
			)
		}
		return models.Category{}, err
	}

	return category, nil
}
