package service

import (
	"github.com/grazierShahid/microserive-blogspot/user-service/models"
	"github.com/grazierShahid/microserive-blogspot/user-service/repository"
)

type UserService interface {
	Register(user *models.User) error
	Login(user *models.LoginRequest) models.LoginResponse
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

func (s *userService) Login(req *models.LoginRequest) models.LoginResponse {
	user, err := s.repo.GetByUsername(req.Username)
	if err != nil {
		return models.LoginResponse{Status: "user not found"}
	}

	if user.Password != req.Password {
		return models.LoginResponse{Status: "invalid password"}
	}

	if user.IsDeleted || !user.Active {
		return models.LoginResponse{Status: "user inactive or deleted"}
	}

	return models.LoginResponse{Status: "success"}
}
