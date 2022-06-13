package database

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/onemgvv/storage-service.git/internal/config"
)

func Init(cfg *config.Config) (*sqlx.DB, error) {
	var dsn = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		cfg.Database.Host, cfg.Database.Port, cfg.Database.Name, cfg.Database.User, cfg.Database.Password)

	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func Close(db *sqlx.DB) error {
	if err := db.Close(); err != nil {
		return err
	}

	return nil
}
