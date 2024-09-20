package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/r1nb0/food-app/auth-svc/internal/config"
)

func InitDB(cfg *config.Config) (*sqlx.DB, error) {

	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=%s password=%s",
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.Username,
		cfg.Postgres.DBName,
		cfg.Postgres.SSLMode,
		cfg.Postgres.Password,
	)

	db, err := sqlx.Connect("postgres", dataSourceName)

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, err
}
