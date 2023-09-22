package service

import (
	"crypto/sha256"
	"fmt"

	"github.com/zarasfara/pet-adoption-platoform/internal/entities"
	"github.com/zarasfara/pet-adoption-platoform/internal/repository"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (u *UserService) GetAll() ([]entities.User, error) {
	return u.repo.GetAll()
}

func (u *UserService) Create(user entities.User) {
	user.Password = generatePasswordHash(user.Password)
	u.repo.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha256.New().Sum([]byte(password))
	hashString := fmt.Sprintf("%x", hash)
	return hashString
}