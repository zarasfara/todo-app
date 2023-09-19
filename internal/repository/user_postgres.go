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

	query := fmt.Sprintf("SELECT id, name, username, password FROM %s", usersTable)

	err := u.db.Select(&userList, query)
	if err != nil {
		return nil, err
	}

	return userList, nil
}

func (u *UserPostgres) CreateUser(user entities.User) error {

	query := fmt.Sprintf("INSERT INTO %s (name, username, password) VALUES ($1, $2, $3)", usersTable)

	stmt, err := u.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Name, user.UserName, user.Password)
	if err != nil {
		return err
	}

	return nil
}