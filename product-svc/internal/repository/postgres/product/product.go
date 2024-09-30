package product

import (
	"context"
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/r1nb0/food-app/product-svc/internal/domain/models"
	"github.com/r1nb0/food-app/product-svc/internal/repository"
)

type productRepository struct {
	db *sqlx.DB
}

func NewProductRepository(db *sqlx.DB) repository.ProductRepository {
	return &productRepository{
		db: db,
	}
}

func (r *productRepository) Create(ctx context.Context, product models.ProductCreate) (int64, error) {
	stmt, err := r.db.PrepareContext(
		ctx,
		`INSERT INTO product (name, description,image_url,price,is_daily_rec,category_id) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`,
	)
	if err != nil {
		return 0, err
	}

	var id int64
	if err := stmt.QueryRowContext(
		ctx,
		product.Name,
		product.Description,
		product.ImageURL,
		product.Price,
		product.IsDailyRec,
		product.CategoryID,
	).Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *productRepository) Update(ctx context.Context, product models.Product) error {
	//TODO implement
	return nil
}

func (r *productRepository) Delete(ctx context.Context, id int64) error {
	stmt, err := r.db.PrepareContext(ctx, "DELETE FROM product WHERE id = $1")
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

func (r *productRepository) GetAll(ctx context.Context) ([]models.Product, error) {
	stmt, err := r.db.PrepareContext(
		ctx,
		`SELECT id, name, description, image_url, price,is_daily_rec, category_id FROM product`,
	)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		return nil, err
	}

	var products []models.Product
	for rows.Next() {
		product, err := scanProduct(rows)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

func (r *productRepository) GetByID(ctx context.Context, id int64) (models.Product, error) {
	stmt, err := r.db.PrepareContext(
		ctx,
		`SELECT id, name, description, image_url, price, is_daily_rec, category_id FROM product WHERE id = $1`,
	)
	if err != nil {
		return models.Product{}, err
	}

	var product models.Product
	if err := stmt.QueryRowContext(ctx, id).Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		&product.ImageURL,
		&product.Price,
		&product.IsDailyRec,
		&product.CategoryID,
	); err != nil {
		return models.Product{}, err
	}

	return product, nil
}

func (r *productRepository) GetDailyRecs(ctx context.Context) ([]models.Product, error) {
	stmt, err := r.db.PrepareContext(
		ctx,
		`SELECT id, name, description, image_url, price, is_daily_rec, category_idFROM product WHERE is_daily_rec = true`,
	)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		return nil, err
	}

	var products []models.Product
	for rows.Next() {
		product, err := scanProduct(rows)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

func (r *productRepository) GetByCategory(ctx context.Context, categoryID int64) ([]models.Product, error) {
	stmt, err := r.db.PrepareContext(
		ctx,
		`SELECT id, name, description, image_url, price, is_daily_rec, category_id FROM product WHERE category_id = $1`,
	)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.QueryContext(ctx, categoryID)
	if err != nil {
		return nil, err
	}

	var products []models.Product
	for rows.Next() {
		product, err := scanProduct(rows)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

func scanProduct(rows *sql.Rows) (models.Product, error) {
	var product models.Product
	if err := rows.Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		&product.ImageURL,
		&product.Price,
		&product.IsDailyRec,
		&product.CategoryID,
	); err != nil {
		return models.Product{}, err
	}

	return product, nil
}
