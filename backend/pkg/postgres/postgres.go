package postgres

import (
	"fmt"
	"rest-api/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func New(cfg config.Postgres) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf(
		"user=%s password=%s host=%s port=%d dbname=%s sslmode=%s",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Database,
		cfg.SSL,
	))
	if err != nil {
		return nil, err
	}
	return db, nil
}
