package service

import "github.com/zarasfara/pet-adoption-platoform/internal/repository"

type Service struct {

}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}