package repositories

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/antibomberman/mego-user/internal/models"

	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

type UserRepository interface {
	GetById(string) (*models.User, error)
	GetByEmail(string) (*models.User, error)
	GetByPhone(string) (*models.User, error)
	Find() ([]models.User, error)
	Update(user *models.User) error
	Delete(id string) error
	SetEmailCode(id string, code string) error
	GetEmailCode(id string) (string, error)
}
type userRepository struct {
	db    *sqlx.DB
	redis *redis.Client
}

func NewUserRepository(db *sqlx.DB, rdb *redis.Client) UserRepository {
	return &userRepository{
		db:    db,
		redis: rdb,
	}
}

func (r *userRepository) Find() ([]models.User, error) {
	var users []models.User
	err := r.db.Select(&users, "SELECT * FROM users")
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return users, nil
}
func (r *userRepository) GetById(id string) (*models.User, error) {
	var user models.User
	err := r.db.Get(&user, "SELECT * FROM users WHERE id = ?", id)
	if err != nil {
		log.Printf("Error getting user by id: %v", err)
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
func (r *userRepository) GetByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Get(&user, "SELECT * FROM users WHERE email = ?", email)
	if err != nil {
		log.Printf("Error getting user by email: %v", err)
		return nil, err
	}
	return &user, nil
}
func (r *userRepository) GetByPhone(phone string) (*models.User, error) {
	var user models.User
	err := r.db.Get(&user, "SELECT * FROM users WHERE phone = $1", phone)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
func (r *userRepository) Create(user *models.User) (models.User, error) {
	result, err := r.db.Exec("INSERT INTO users (id, first_name, middle_name, last_name, email, phone, password, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7, NOW())",
		user.Id, user.FirstName, user.MiddleName, user.LastName, user.Email, user.Phone, user.Password)
	if err != nil {
		return *user, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return *user, err
	}
	user.Id = fmt.Sprintf("%d", id)

	return *user, nil

}
func (r *userRepository) Update(user *models.User) error {
	_, err := r.db.Exec("UPDATE users SET first_name = $2, middle_name = $3, last_name = $4, email = $5, phone = $6 WHERE id = $1",
		user.Id, user.FirstName, user.MiddleName, user.LastName, user.Email, user.Phone)
	if err != nil {
		return err
	}
	return nil
}
func (r *userRepository) Delete(id string) error {
	_, err := r.db.Exec("UPDATE users SET deleted_at = NOW() WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
func (r *userRepository) Count() (int, error) {
	var count int
	err := r.db.Get(&count, "SELECT COUNT(*) FROM users")
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *userRepository) SetEmailCode(email, code string) error {
	err := r.redis.Set(context.Background(), email, code, time.Minute*5).Err()
	return err
}
func (r *userRepository) GetEmailCode(email string) (string, error) {
	return r.redis.Get(context.Background(), email).Result()
}
