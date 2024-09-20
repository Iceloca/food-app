package repository

import (
	"context"
	"github.com/r1nb0/food-app/auth-svc/internal/domain/models"
)

type UserRepository interface {
	SaveUser(ctx context.Context, email string, passHash string) (int64, error)
	GetUserByEmail(ctx context.Context, email string) (models.User, error)
}
