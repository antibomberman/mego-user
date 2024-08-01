package services

import (
	"fmt"
	"github.com/antibomberman/mego-user/internal/models"
	"github.com/antibomberman/mego-user/internal/repositories"
	"log"
)

type UserService interface {
	Find(pageSize int, pageToken, search string) ([]models.UserDetails, string, error)
	GetById(string) (*models.UserDetails, error)
}

type userService struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userService{userRepository: userRepo}
}

func (s *userService) Find(pageSize int, pageToken, search string) ([]models.UserDetails, string, error) {

	return []models.UserDetails{}, "", nil
}
func (s *userService) GetById(id string) (*models.UserDetails, error) {
	user, err := s.userRepository.GetById(id)
	fmt.Printf("User: %+v", user)
	if err != nil {
		log.Printf("Error getting user by ID: %v", err)
		return nil, err
	}

	userDetails := models.UserDetails{
		Id:         user.Id,
		FirstName:  user.FirstName.String,
		MiddleName: user.MiddleName.String,
		LastName:   user.LastName.String,
		Email:      user.Email.String,
		Phone:      user.Phone.String,
		Avatar:     user.Avatar.String,
		DeletedAt:  user.DeletedAt.Time,
		CreatedAt:  user.DeletedAt.Time,
		UpdatedAt:  user.DeletedAt.Time,
	}

	return &userDetails, nil
}
