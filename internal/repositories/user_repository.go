package repositories

import "github.com/jmoiron/sqlx"

type UserRepository interface {
}
type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}
