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
