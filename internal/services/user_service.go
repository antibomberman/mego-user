package services

import (
	"github.com/antibomberman/mego-user/internal/models"
	"github.com/antibomberman/mego-user/internal/repositories"
)

type UserService interface {
	Find(pageSize int, pageToken, search string) ([]models.User, string, error)
}

type userService struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userService{userRepository: userRepo}
}

func (s *userService) Find(pageSize int, pageToken, search string) ([]models.User, string, error) {

	return []models.User{}, "", nil
}
