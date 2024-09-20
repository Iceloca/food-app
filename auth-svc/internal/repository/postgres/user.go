package postgres

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/r1nb0/food-app/auth-svc/internal/domain/models"
	"github.com/r1nb0/food-app/auth-svc/internal/repository"
)

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) repository.UserRepository {
	return &userRepository{
		db: db,
	}
}

// TODO add logger
func (r *userRepository) SaveUser(ctx context.Context, email string, passHash []byte) (int64, error) {
	stmt, err := r.db.PrepareContext(ctx, "INSERT INTO users (email, pass_hash) VALUES ($1, $2) RETURNING id")
	if err != nil {
		return 0, err
	}

	var id int64
	if err = stmt.QueryRowContext(ctx, email, passHash).Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

// TODO add logger
func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (models.User, error) {
	stmt, err := r.db.PrepareContext(ctx, "SELECT id, email, pass_hash FROM users WHERE email = $1")

	if err != nil {
		return models.User{}, err
	}

	var user models.User
	row := stmt.QueryRowContext(ctx, email)
	if err = row.Scan(&user.ID, &user.Email, &user.PassHash); err != nil {
		return models.User{}, nil
	}

	return user, nil
}
