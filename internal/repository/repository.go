package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/zarasfara/pet-adoption-platoform/internal/entities"
)

type User interface {
	GetAll() ([]entities.User, error)
	CreateUser(user entities.User) error
}

type Repository struct {
	User
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		User: NewUserPostgres(db),
	}
}
