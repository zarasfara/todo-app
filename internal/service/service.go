package service

import (
	"github.com/zarasfara/pet-adoption-platoform/internal/entities"
	"github.com/zarasfara/pet-adoption-platoform/internal/repository"
)

type User interface {
	GetAll() ([]entities.User, error)
	Create(user entities.User)
}

type Todo interface {
	GetAll() ([]entities.Todo, error)
	CreateTodo(todo entities.Todo) error
	GetTodoByID(id int) (entities.Todo, error)
	UpdateTodo(id int, input entities.UpdateTodoInput) error
	Delete(id int) error
	ChangeStatus(id int) error
}

type Service struct {
	User
	Todo
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		User: NewUserService(repos.User),
		Todo: NewTodoService(repos.Todo),
	}
}
