package service

import (
	"github.com/zarasfara/pet-adoption-platoform/internal/entities"
	"github.com/zarasfara/pet-adoption-platoform/internal/repository"
)

type Todo interface {
	GetAll() ([]entities.Todo, error)
	CreateTodo(todo entities.Todo) error
	GetTodoByID(id int) (entities.Todo, error)
	UpdateTodo(id int, input entities.UpdateTodoInput) error
	Delete(id int) error
	ChangeStatus(id int) error
}

type Service struct {
	Todo
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Todo: NewTodoService(repos.Todo),
	}
}
