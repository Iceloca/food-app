package service

import (
	"context"
	"github.com/r1nb0/food-app/cart-svc/internal/domain/models"
)

type CartService interface {
	GetByID(ctx context.Context, id string) (models.Cart, error)
	GetAll(ctx context.Context) ([]models.Cart, error)
	Create(ctx context.Context, cart models.CartCreate) (string, error)
	Delete(ctx context.Context, id string) error
	AddItem(ctx context.Context, basketID string, item models.Item) error
	DeleteItem(ctx context.Context, basketID string, itemID int64) error
	UpdateItem(ctx context.Context, basketID string, item models.Item) error
}
