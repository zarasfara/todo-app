package service

import (
	"github.com/zarasfara/pet-adoption-platoform/internal/entities"
	"github.com/zarasfara/pet-adoption-platoform/internal/repository"
)

type TodoService struct {
	repo repository.Todo
}

func NewTodoService(repo repository.Todo) *TodoService {
	return &TodoService{
		repo: repo,
	}
}

func (t *TodoService) GetAll() ([]entities.Todo, error) {
	return t.repo.GetAll()
}

func (t *TodoService) CreateTodo(todo entities.Todo) error {
	return t.repo.CreateTodo(todo)
}

func (t *TodoService) GetTodoByID(id int) (entities.Todo, error) {
	return t.repo.GetTodoByID(id)
}