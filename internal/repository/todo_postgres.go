package repository

import (
	"fmt"
	"strings"

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

func (r *TodoPostgres) UpdateTodo(id int, input entities.UpdateTodoInput) error {
	queryParts := []string{fmt.Sprintf("UPDATE %s SET completed = $1", todoTable)}
	var args []interface{}
	args = append(args, input.Completed)

	if input.Title != nil {
		queryParts = append(queryParts, "title = $2")
		args = append(args, input.Title)
	}

	if input.Description != nil {
		queryParts = append(queryParts, "description = $3")
		args = append(args, input.Description)
	}

	queryParts = append(queryParts, "WHERE id = $4")
	args = append(args, id)

	query := strings.Join(queryParts, ", ")
	_, err := r.db.Exec(query, args...)
	if err != nil {
		return err
	}
	return nil
}

func (r *TodoPostgres) Delete(id int) error {
	query := fmt.Sprintf("DELETE from %s WHERE id = $1", todoTable)

	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *TodoPostgres) ChangeStatus(id int) error {
	query := fmt.Sprintf("UPDATE %s SET completed = NOT completed WHERE id = $1", todoTable)

	if _, err := r.db.Exec(query, id); err != nil {
		return err
	}

	return nil
}