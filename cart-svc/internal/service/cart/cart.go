package cart

import (
	"context"
	"errors"
	"github.com/r1nb0/food-app/cart-svc/internal/domain/models"
	"github.com/r1nb0/food-app/cart-svc/internal/repository"
	"github.com/r1nb0/food-app/cart-svc/internal/service"
	"github.com/r1nb0/food-app/pkg/database"
	"log/slog"
)

type cartService struct {
	cartRepository repository.CartRepository
	logger         *slog.Logger
}

func NewCartService(cartRepository repository.CartRepository, logger *slog.Logger) service.CartService {
	return &cartService{
		cartRepository: cartRepository,
		logger:         logger,
	}
}

func (s *cartService) GetByID(ctx context.Context, id string) (models.Cart, error) {
	const op = "cartService.GetByID"

	log := s.logger.With(
		slog.String("operation", op),
		slog.String("data", id),
	)

	log.Info("attempting to get cart by id")

	cart, err := s.cartRepository.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, database.ErrNotFound) {
			log.Error(
				"cart not found",
				slog.String("error", err.Error()),
			)
		} else {
			log.Error(
				"failed to get cart by id",
				slog.String("error", err.Error()),
			)
		}
		return models.Cart{}, err
	}

	return cart, nil
}

func (s *cartService) GetAll(ctx context.Context) ([]models.Cart, error) {
	const op = "cartService.GetAll"

	log := s.logger.With(slog.String("operation", op))

	log.Info("attempting to get all carts")

	carts, err := s.cartRepository.GetAll(ctx)
	if err != nil {
		if errors.Is(err, database.ErrNotFound) {
			log.Error(
				"not found carts",
				slog.String("error", err.Error()),
			)
		} else {
			log.Error(
				"failed to get carts",
				slog.String("error", err.Error()),
			)
		}
		return nil, err
	}

	return carts, nil
}

func (s *cartService) Create(ctx context.Context, cart models.CartCreate) (string, error) {
	const op = "cartService.Create"

	log := s.logger.With(
		slog.String("operation", op),
		slog.Any("data", cart),
	)

	log.Info("attempting to create cart")

	id, err := s.cartRepository.Create(ctx, cart)
	if err != nil {
		log.Error(
			"failed to create cart",
			slog.String("error", err.Error()),
		)
	}

	return id, err
}

func (s *cartService) Delete(ctx context.Context, id string) error {
	const op = "cartService.Delete"

	log := s.logger.With(
		slog.String("operation", op),
		slog.String("data", id),
	)

	log.Info("attempting to delete cart")

	err := s.cartRepository.Delete(ctx, id)
	if err != nil {
		if errors.Is(err, database.ErrNotFound) {
			log.Error(
				"cart not found",
				slog.String("error", err.Error()),
			)
		} else {
			log.Error(
				"failed to delete cart",
				slog.String("error", err.Error()),
			)
		}
		return err
	}

	return nil
}

func (s *cartService) AddItem(ctx context.Context, basketID string, item models.Item) error {
	const op = "cartService.AddItem"

	log := s.logger.With(
		slog.String("operation", op),
		slog.Any("data", item),
	)

	log.Info("attempting to add item")

	err := s.cartRepository.AddItem(ctx, basketID, item)

	if err != nil {
		log.Error(
			"failed to add item",
			slog.String("error", err.Error()),
		)
		return err
	}

	return nil
}

func (s *cartService) DeleteItem(ctx context.Context, basketID string, itemID int64) error {
	const op = "cartService.DeleteItem"

	log := s.logger.With(
		"operation", op,
		slog.String("basket_id", basketID),
		slog.Any("item_id", itemID),
	)

	log.Info("attempting to delete item")

	err := s.cartRepository.DeleteItem(ctx, basketID, itemID)
	if err != nil {
		if errors.Is(err, database.ErrNotFound) {
			log.Error(
				"item not found",
				slog.String("error", err.Error()),
			)
		} else {
			log.Error(
				"failed to delete item",
				slog.String("error", err.Error()),
			)
		}
		return err
	}

	return nil
}

func (s *cartService) UpdateItem(ctx context.Context, basketID string, item models.Item) error {
	return nil
}
