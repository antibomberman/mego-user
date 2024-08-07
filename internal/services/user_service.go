package services

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	pb "github.com/antibomberman/mego-protos/gen/go/auth"
	"github.com/antibomberman/mego-protos/gen/go/storage"
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
	Create(user *models.CreateUserRequest) (*models.UserDetails, error)
	Update(id string, user *models.UpdateUserRequest) (*models.UserDetails, error)
	Delete(id string) error
	ForceDelete(id string) error
	UpdateProfile(id, firstName, middleName, lastName, about string, avatar *models.NewAvatar) (*models.UserDetails, error)
	UpdateLang(id, lang string) error
	UpdateTheme(id, theme string) error
	UpdateEmailSendCode(id, email string) error
	UpdateEmail(id, code string) error
}

type userService struct {
	userRepository repositories.UserRepository
	authClient     *clients.AuthClient
	storageClient  *clients.StorageClient
}

func NewUserService(userRepo repositories.UserRepository, client *clients.AuthClient, storageClient *clients.StorageClient) UserService {
	return &userService{userRepository: userRepo, authClient: client, storageClient: storageClient}
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
func (s *userService) Create(data *models.CreateUserRequest) (*models.UserDetails, error) {
	newUser, err := s.userRepository.Create(data)
	if err != nil {
		return nil, fmt.Errorf("error creating user: %v", err)
	}
	userDetail := dto.ToUserDetail(newUser)
	return userDetail, nil
}
func (s *userService) Update(id string, data *models.UpdateUserRequest) (*models.UserDetails, error) {

	user := &models.User{
		FirstName: sql.NullString{String: data.FirstName},
		LastName:  sql.NullString{String: data.LastName},
	}
	if data.Avatar != nil {
		oldUser, err := s.userRepository.GetById(id)
		if err != nil {
			return nil, fmt.Errorf("user not found")
		}
		if oldUser.Avatar.String != "" {
			_, err := s.storageClient.DeleteObject(context.Background(), &storage.DeleteObjectRequest{FileName: oldUser.Avatar.String})
			if err != nil {
				return nil, err
			}
		}
		object, err := s.storageClient.PutObject(context.Background(), &storage.PutObjectRequest{
			FileName: data.Avatar.FileName,
			Data:     data.Avatar.Data,
		})
		if err != nil {
			return nil, err
		}
		user.Avatar = sql.NullString{String: object.FileName, Valid: true}

	}
	err := s.userRepository.Update(id, user)

	if err != nil {
		return nil, fmt.Errorf("error updating user: %v", err)
	}
	return s.GetById(id)
}
func (s *userService) Delete(id string) error {
	err := s.userRepository.Delete(id)
	if err != nil {
		return fmt.Errorf("error deleting user: %v", err)
	}
	return nil
}
func (s *userService) ForceDelete(id string) error {
	err := s.userRepository.Delete(id)
	if err != nil {
		return fmt.Errorf("error deleting user: %v", err)
	}
	return nil
}

func (s *userService) UpdateProfile(id, firsName, middleName, lastName, about string, avatar *models.NewAvatar) (*models.UserDetails, error) {
	user := &models.User{
		FirstName:  sql.NullString{String: firsName, Valid: true},
		MiddleName: sql.NullString{String: middleName, Valid: true},
		LastName:   sql.NullString{String: lastName, Valid: true},
		About:      sql.NullString{String: about, Valid: true},
	}

	if avatar != nil {
		oldUser, err := s.userRepository.GetById(id)
		if err != nil {
			return nil, fmt.Errorf("user not found")
		}
		if oldUser.Avatar.String != "" {
			_, err := s.storageClient.DeleteObject(context.Background(), &storage.DeleteObjectRequest{FileName: oldUser.Avatar.String})
			if err != nil {
				return nil, err
			}
		}
		object, err := s.storageClient.PutObject(context.Background(), &storage.PutObjectRequest{
			FileName: avatar.FileName,
			Data:     avatar.Data,
		})
		if err != nil {
			return nil, err
		}
		user.Avatar = sql.NullString{String: object.FileName, Valid: true}
	}
	err := s.userRepository.Update(id, user)
	if err != nil {
		return nil, fmt.Errorf("error updating user: %v", err)
	}
	return s.GetById(id)
}
func (s *userService) UpdateLang(id, lang string) error {
	user := &models.User{
		Lang: sql.NullString{String: lang, Valid: true},
	}

	err := s.userRepository.Update(id, user)
	if err != nil {
		return fmt.Errorf("error updating user: %v", err)
	}
	return nil
}
func (s *userService) UpdateTheme(id, theme string) error {
	user := &models.User{
		Theme: sql.NullString{String: theme, Valid: true},
	}

	err := s.userRepository.Update(id, user)
	if err != nil {
		return fmt.Errorf("error updating user: %v", err)
	}
	return nil
}

func (s *userService) UpdateEmailSendCode(id, email string) error {
	exists := s.userRepository.ExistsEmail(email)
	if exists {
		return errors.New("email already exists")
	}
	code := generateRandomCode(1000, 9999)
	return s.userRepository.EmailUpdateSaveCode(id, email, code)

}
func (s *userService) UpdateEmail(id, code string) error {
	email, err := s.userRepository.EmailUpdateCheckCode(id, code)
	if err != nil {
		return errors.New("invalid code")
	}
	err = s.userRepository.Update(id, &models.User{Email: sql.NullString{String: email, Valid: true}})
	if err != nil {
		return err
	}
	return nil
}
func generateRandomCode(max, min int) string {
	//return strconv.Itoa(rand.IntN(max-min) + min)
	return "1234"
}
