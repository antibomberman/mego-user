package repositories

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/antibomberman/mego-user/internal/models"
	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	GetById(string) (*models.User, error)
}
type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return &userRepository{
		db: db,
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
	err := r.db.Get(&user, "SELECT * FROM users WHERE id = $1", id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
func (r *userRepository) GetByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Get(&user, "SELECT * FROM users WHERE email = $1", email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
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
