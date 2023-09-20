package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/zarasfara/pet-adoption-platoform/internal/entities"
)

type User interface {
	GetAll() ([]entities.User, error)
	CreateUser(user entities.User) error
}

type Todo interface {
	GetAll() ([]entities.Todo, error)
	CreateTodo(todo entities.Todo) error
	GetTodoByID(id int) (entities.Todo, error)
}

type Repository struct {
	User
	Todo
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		User: NewUserPostgres(db),
		Todo: NewTodoPostgres(db),
	}
}
