package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/zarasfara/pet-adoption-platoform/internal/entities"
)

type TodoPostgres struct {
	db *sqlx.DB
}

// Create todo postgres repo.
func NewTodoPostgres(db *sqlx.DB) *TodoPostgres {
	return &TodoPostgres{
		db: db,
	}
}

func (t *TodoPostgres) GetAll() ([]entities.Todo, error) {
	var todos []entities.Todo

	query := fmt.Sprintf("SELECT id, title, description, created_at FROM %s", todoTable)

	err := t.db.Select(&todos, query)
	if err != nil {
		return nil, err
	}

	return todos, nil
}

func (t *TodoPostgres) CreateTodo(todo entities.Todo) error {
	query := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2)", todoTable)

	stmt, err := t.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(todo.Title, todo.Description)
	if err != nil {
		return err
	}

	return nil
}

func (t *TodoPostgres) GetTodoByID(id int) (entities.Todo, error) {
	var todo entities.Todo

	query := fmt.Sprintf("SELECT id, title, description, created_at FROM %s WHERE id = $1", todoTable)

	err := t.db.Get(&todo, query, id)
	if err != nil {
		return todo, err
	}

	return todo, nil
}
