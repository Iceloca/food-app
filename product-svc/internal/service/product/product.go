package product

import (
	"context"
	"errors"
	"github.com/r1nb0/food-app/pkg/database"
	"github.com/r1nb0/food-app/product-svc/internal/domain/models"
	"github.com/r1nb0/food-app/product-svc/internal/repository"
	"github.com/r1nb0/food-app/product-svc/internal/service"
	"log/slog"
)

type productService struct {
	productRepository  repository.ProductRepository
	categoryRepository repository.CategoryRepository
	logger             *slog.Logger
}

func NewProductService(
	productRepository repository.ProductRepository,
	categoryRepository repository.CategoryRepository,
	logger *slog.Logger,
) service.ProductService {
	return &productService{
		productRepository:  productRepository,
		categoryRepository: categoryRepository,
		logger:             logger,
	}
}

func (s *productService) Create(
	ctx context.Context,
	product models.ProductCreate,
) (int64, error) {
	const op = "productService.Create"

	log := s.logger.With(
		slog.String("operation", op),
		slog.Any("data", product),
	)

	log.Info("attempting to create product")

	_, err := s.categoryRepository.GetByID(ctx, product.CategoryID)
	if err != nil {
		log.Error(
			"category not found",
			//TODO handle err from GRPC
			slog.String("error", err.Error()),
		)
		return 0, err
	}

	id, err := s.productRepository.Create(ctx, product)
	if err != nil {
		log.Error(
			"failed to create product",
			slog.String("error", err.Error()),
		)
		return 0, err
	}

	return id, nil
}

func (s *productService) Update(
	ctx context.Context,
	product models.Product,
) error {
	const op = "productService.Update"

	log := s.logger.With(
		slog.String("operation", op),
		slog.Any("data", product),
	)

	log.Info("attempting to update product")

	if err := s.productRepository.Update(ctx, product); err != nil {
		log.Error(
			"failed to update product",
			slog.String("error", err.Error()),
		)
		return err
	}

	return nil
}

func (s *productService) Delete(ctx context.Context, id int64) error {
	const op = "productService.Delete"
	log := s.logger.With(
		slog.String("operation", op),
		slog.Any("data", id),
	)

	log.Info("attempting to delete product")

	if err := s.productRepository.Delete(ctx, id); err != nil {
		log.Error(
			"failed to delete product",
			slog.String("error", err.Error()),
		)
		return err
	}

	return s.productRepository.Delete(ctx, id)
}

func (s *productService) GetAll(ctx context.Context) ([]models.Product, error) {
	const op = "productService.GetAll"

	log := s.logger.With(slog.String("operation", op))

	log.Info("attempting to get all products")

	products, err := s.productRepository.GetAll(ctx)
	if err != nil {
		if errors.Is(err, database.ErrNotFound) {
			log.Error("no products found")
		} else {
			log.Error(
				"failed to get all products",
				slog.String("error", err.Error()),
			)
		}
		return nil, err
	}

	return products, nil
}

func (s *productService) GetByID(ctx context.Context, id int64) (models.Product, error) {
	const op = "productService.GetByID"

	log := s.logger.With(
		slog.String("operation", op),
		slog.Any("data", id),
	)

	log.Info("attempting to get product by id")
	product, err := s.productRepository.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, database.ErrNotFound) {
			log.Error(
				"product not found",
				slog.String("error", err.Error()),
			)
		}
		log.Error("failed to get product by id",
			slog.String("error", err.Error()),
		)
		return models.Product{}, err
	}

	return product, nil
}

func (s *productService) GetByCategory(ctx context.Context, categoryID int64) ([]models.Product, error) {
	const op = "productService.GetByCategory"

	log := s.logger.With(
		slog.String("operation", op),
		slog.Any("data", categoryID),
	)

	log.Info("attempting to get products by category")

	_, err := s.categoryRepository.GetByID(ctx, categoryID)
	if err != nil {
		log.Error(
			"category not found",
			slog.String("error", err.Error()),
		)
		return nil, err
	}

	products, err := s.productRepository.GetByCategory(ctx, categoryID)
	if err != nil {
		if errors.Is(err, database.ErrNotFound) {
			log.Error(
				"no products found",
				slog.String("error", err.Error()),
			)
		} else {
			log.Error(
				"failed to get products by category",
				slog.String("error", err.Error()),
			)
		}
		return nil, err
	}

	return products, err
}

func (s *productService) GetDailyRecs(ctx context.Context) ([]models.Product, error) {
	const op = "productService.GetDailyRecs"

	log := s.logger.With(slog.String("operation", op))

	log.Info("attempting to get daily recs")

	products, err := s.productRepository.GetDailyRecs(ctx)
	if err != nil {
		if errors.Is(err, database.ErrNotFound) {
			log.Error(
				"daily recs not found",
				slog.String("error", err.Error()))
		} else {
			log.Error("failed to get daily recs",
				slog.String("error", err.Error()),
			)
		}
		return nil, err
	}
	return products, nil
}
