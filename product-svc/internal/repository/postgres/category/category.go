package category

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/r1nb0/food-app/product-svc/internal/domain/models"
	"github.com/r1nb0/food-app/product-svc/internal/repository"
	"strings"
)

type categoryRepository struct {
	db *sqlx.DB
}

func NewCategoryRepository(db *sqlx.DB) repository.CategoryRepository {
	return &categoryRepository{
		db: db,
	}
}

func (r *categoryRepository) Create(ctx context.Context, category models.CategoryCreate) (int64, error) {
	stmt, err := r.db.PrepareContext(ctx, "INSERT INTO category (name, image_url) VALUES ($1, $2) RETURNING id")
	if err != nil {
		return 0, err
	}

	var id int64
	if err = stmt.QueryRowContext(ctx, category.Name, category.ImageURL).Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *categoryRepository) Update(ctx context.Context, category models.Category) error {
	var (
		placeholders []string
		args         []interface{}
		argID        = 1
	)

	if category.Name != "" {
		placeholders = append(placeholders, fmt.Sprintf("name=$%d", argID))
		args = append(args, category.Name)
		argID++
	}

	if category.ImageURL != "" {
		placeholders = append(placeholders, fmt.Sprintf("image_url=$%d", argID))
		args = append(args, category.ImageURL)
		argID++
	}

	if len(placeholders) == 0 {
		return repository.ErrUpdate
	}

	args = append(args, category.ID)

	stmt, err := r.db.PrepareContext(ctx, fmt.Sprintf("UPDATE category SET %s WHERE id = $%d", strings.Join(placeholders, ","), argID))
	if err != nil {
		return err
	}

	res, err := stmt.ExecContext(ctx, args...)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return repository.ErrNotFound
	}

	return nil
}

func (r *categoryRepository) Delete(ctx context.Context, id int64) error {
	stmt, err := r.db.PrepareContext(ctx, "DELETE FROM category WHERE id = $1")
	if err != nil {
		return err
	}

	res, err := stmt.ExecContext(ctx, id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return repository.ErrNotFound
	}

	return nil
}

func (r *categoryRepository) GetAll(ctx context.Context) ([]models.Category, error) {
	stmt, err := r.db.PrepareContext(ctx, "SELECT id, name, image_url FROM category")
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	var categories []models.Category
	for rows.Next() {
		var category models.Category
		if err = rows.Scan(&category.ID, &category.Name, &category.ImageURL); err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}

func (r *categoryRepository) GetByID(ctx context.Context, id int64) (models.Category, error) {
	stmt, err := r.db.PrepareContext(ctx, "SELECT id, name, image_url FROM category WHERE id = $1")
	if err != nil {
		return models.Category{}, err
	}

	var category models.Category
	if err = stmt.QueryRowContext(ctx, id).Scan(&category.ID, &category.Name, &category.ImageURL); err != nil {
		return models.Category{}, err
	}

	return category, nil
}
