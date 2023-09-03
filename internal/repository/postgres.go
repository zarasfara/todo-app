package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/zarasfara/pet-adoption-platoform/internal/config"
)

const (
	users = "users"
)

func NewPostgresDB(cfg config.Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.DBConnection.Host,
		cfg.DBConnection.Port,
		cfg.DBConnection.Username,
		cfg.DBConnection.Database,
		cfg.DBConnection.Password,
		"disable",
	))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
