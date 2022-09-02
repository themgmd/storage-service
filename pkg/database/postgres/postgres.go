package postgres

import (
	"fmt"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/onemgvv/storage-service/internal/config"
)

func Init(cfg *config.Config) (*sqlx.DB, error) {
	postgres := cfg.Database.Postgres
	var dsn = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		postgres.Host, postgres.Port, postgres.User, postgres.Name, postgres.Password)

	db, err := sqlx.Open("pgx", dsn)
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
