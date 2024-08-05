package services

import (
	"fmt"
	"github.com/antibomberman/mego-user/internal/dto"
	"github.com/antibomberman/mego-user/internal/models"
	"github.com/antibomberman/mego-user/internal/repositories"
	"github.com/antibomberman/mego-user/internal/secure"
	"github.com/antibomberman/mego-user/pkg/utils"
	"log"
)

type UserService interface {
	Find(pageSize int, pageToken, sort, search string) ([]*models.UserDetails, string, error)
	GetById(string) (*models.UserDetails, error)
	LoginByEmail(email, code string) (string, error)
	LoginByEmailSendCode(email string) error
	GetByToken(token string) (*models.UserDetails, error)
	GetByEmail(email string) (*models.UserDetails, error)
	GetByPhone(phone string) (*models.UserDetails, error)
}

type userService struct {
	userRepository repositories.UserRepository
	secure         secure.Secure
}

func NewUserService(userRepo repositories.UserRepository, secure secure.Secure) UserService {
	return &userService{userRepository: userRepo, secure: secure}
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
	claims, err := s.secure.Parse(token)
	if err != nil {
		return nil, fmt.Errorf("error parsing token: %v", err)
	}
	return s.GetById(claims.UserId)
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

func (s *userService) LoginByEmail(email, code string) (string, error) {
	user, err := s.userRepository.GetByEmail(email)
	if err != nil {
		log.Printf("Error getting user by email: %v", err)
		return "", fmt.Errorf("invalid email")
	}

	savedCode, err := s.userRepository.GetEmailCode(email)
	if err != nil {
		log.Printf("Error getting email code: %v", err)
		return "", fmt.Errorf("invalid code")
	}
	if savedCode != code {
		log.Printf("Invalid code != saved code")
		return "", fmt.Errorf("invalid code")
	}

	token, err := s.secure.Generate(user.Id)
	if err != nil {
		log.Printf("Error generating token: %v", err)
		return "", fmt.Errorf("error generating token: %v", err)
	}
	return token, nil
}

func (s *userService) LoginByEmailSendCode(email string) error {
	_, err := s.userRepository.GetByEmail(email)
	if err != nil {
		return fmt.Errorf("invalid email")
	}
	code := generateRandomCode(9999, 1000)
	return s.userRepository.SetEmailCode(email, code)
}

func generateRandomCode(max, min int) string {
	//return strconv.Itoa(rand.IntN(max-min) + min)
	return "1234"
}
