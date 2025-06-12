package service

import (
	"github.com/grazierShahid/microserive-blogspot/user-service/models"
	"github.com/grazierShahid/microserive-blogspot/user-service/repository"
)

type UserService interface {
	Register(user *models.User) error
}

type userService struct {
	repo repository.User
}

func NewUserService(r repository.User) UserService {
	return &userService{repo: r}
}

func (s *userService) Register(user *models.User) error {
	return s.repo.Register(user)
}
