package repository

import (
	"basket-svc/internal/domain/models"
	"context"
)

type BasketRepository interface {
	GetByID(ctx context.Context, id string) (models.Item, error)
	GetAll(ctx context.Context) ([]models.Item, error)
	Create(ctx context.Context, userID int64) (string, error)
	Delete(ctx context.Context, id string) (bool, error)
	AddItem(ctx context.Context, basketID string, item models.Item)
	DeleteItem(ctx context.Context, basketID string, itemID int64) (bool, error)
	UpdateItem(ctx context.Context, basketID string, item models.Item) (bool, error)
}
