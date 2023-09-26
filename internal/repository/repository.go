package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/zarasfara/pet-adoption-platoform/internal/entities"
)

type Todo interface {
	GetAll() ([]entities.Todo, error)
	CreateTodo(todo entities.Todo) error
	GetTodoByID(id int) (entities.Todo, error)
}

type Repository struct {
	Todo
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Todo: NewTodoPostgres(db),
	}
}
