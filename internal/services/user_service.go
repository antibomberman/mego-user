package services

import (
	"github.com/antibomberman/mego-user/internal/models"
	"github.com/antibomberman/mego-user/internal/repositories"
)

type UserService interface {
	Find(pageSize int, pageToken, search string) ([]models.User, string, error)
	GetById(string) (*models.UserDetails, error)
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
func (s *userService) GetById(id string) (*models.UserDetails, error) {
	user, err := s.userRepository.GetById(id)
	if err != nil {
		return nil, err
	}

	userDetails := models.UserDetails{
		Id:        user.Id,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return &userDetails, nil
}
