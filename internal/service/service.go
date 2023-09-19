package service

import (
	"github.com/zarasfara/pet-adoption-platoform/internal/entities"
	"github.com/zarasfara/pet-adoption-platoform/internal/repository"
)

type User interface {
	GetAll() ([]entities.User, error)
	CreateUser(user entities.User)
}

type Service struct {
	User 
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		User: NewUserService(repos.User),
	}
}