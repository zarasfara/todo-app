package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/zarasfara/pet-adoption-platoform/internal/entities"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{
		db: db,
	}
}

func (u *UserPostgres) GetAll() ([]entities.User, error) {
	var userList []entities.User

	query := fmt.Sprintf("SELECT id, name, username FROM %s", users)

	err := u.db.Select(&userList, query)
	if err != nil {
		return nil, err
	}

	return userList, nil
}
