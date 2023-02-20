package postgres

import (
	"database/sql"
	"fmt"
	"rest-api/config"

	"github.com/jackc/pgx"
	"github.com/jackc/pgx/stdlib"
)

func New(cfg config.Postgres) (*sql.DB, error) {
	driverConfig := stdlib.DriverConfig{
		ConnConfig: pgx.ConnConfig{
			PreferSimpleProtocol: true,
		},
	}
	stdlib.RegisterDriverConfig(&driverConfig)

	db, err := sql.Open("pgx", fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Database,
	))
	if err != nil {
		return nil, err
	}

	return db, nil
}
