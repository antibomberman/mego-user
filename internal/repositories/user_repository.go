package repositories

import (
	"database/sql"
	"errors"
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

func (r *userRepository) GetById(id string) (*models.User, error) {
	var user models.User
	err := r.db.Get(&user, "SELECT * FROM users WHERE id = 1")
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
