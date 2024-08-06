package services

import (
	"context"
	"fmt"
	pb "github.com/antibomberman/mego-protos/gen/go/auth"
	"github.com/antibomberman/mego-user/internal/clients"
	"github.com/antibomberman/mego-user/internal/dto"
	"github.com/antibomberman/mego-user/internal/models"
	"github.com/antibomberman/mego-user/internal/repositories"
	"github.com/antibomberman/mego-user/pkg/utils"
	"log"
)

type UserService interface {
	Find(pageSize int, pageToken, sort, search string) ([]*models.UserDetails, string, error)
	GetById(string) (*models.UserDetails, error)
	GetByToken(token string) (*models.UserDetails, error)
	GetByEmail(email string) (*models.UserDetails, error)
	GetByPhone(phone string) (*models.UserDetails, error)
}

type userService struct {
	userRepository repositories.UserRepository
	authClient     *clients.AuthClient
}

func NewUserService(userRepo repositories.UserRepository, client *clients.AuthClient) UserService {
	return &userService{userRepository: userRepo, authClient: client}
}

func (s *userService) Find(pageSize int, pageToken, sort, search string) ([]*models.UserDetails, string, error) {
	if pageSize < 1 {
		pageSize = 10
	}
	startIndex := 0

	if pageToken != "" {
		var err error
		startIndex, err = utils.DecodePageToken(pageToken)
		if err != nil {
			log.Printf("Error decoding page token: %v", err)
		}
	}

	users, err := s.userRepository.Find(startIndex, pageSize+1, sort, search)
	if err != nil {
		log.Printf("Error getting users: %v", err)
		return nil, "", err
	}
	var nextPageToken string
	if len(users) > pageSize {
		nextPageToken = utils.EncodePageToken(startIndex + pageSize)
		users = users[:pageSize]
	}
	var userDetails []*models.UserDetails
	for _, user := range users {
		userDetails = append(userDetails, dto.ToUserDetail(&user))
	}
	return userDetails, nextPageToken, nil
}
func (s *userService) GetById(id string) (*models.UserDetails, error) {
	user, err := s.userRepository.GetById(id)
	fmt.Printf("User: %+v", user)
	if err != nil {
		log.Printf("Error getting user by ID: %v", err)
		return nil, err
	}

	return &models.UserDetails{
		Id:         user.Id,
		FirstName:  user.FirstName.String,
		MiddleName: user.MiddleName.String,
		LastName:   user.LastName.String,
		Email:      user.Email.String,
		Phone:      user.Phone.String,
		Avatar:     user.Avatar.String,
		CreatedAt:  user.CreatedAt.Time,
		UpdatedAt:  user.UpdatedAt.Time,
		DeletedAt:  user.DeletedAt.Time,
	}, nil
}
func (s *userService) GetByToken(token string) (*models.UserDetails, error) {
	response, err := s.authClient.Parse(context.Background(), &pb.ParseRequest{Token: token})
	if err != nil {
		return nil, fmt.Errorf("error parsing token: %v", err)
	}
	if response.Success != true {
		return nil, fmt.Errorf("invalid token")
	}
	return s.GetById(response.UserId)
}
func (s *userService) GetByEmail(email string) (*models.UserDetails, error) {
	user, err := s.userRepository.GetByEmail(email)
	if err != nil {
		return nil, fmt.Errorf("invalid email")
	}
	return s.GetById(user.Id)
}
func (s *userService) GetByPhone(phone string) (*models.UserDetails, error) {
	user, err := s.userRepository.GetByPhone(phone)
	if err != nil {
		return nil, fmt.Errorf("invalid phone")
	}
	return s.GetById(user.Id)
}
