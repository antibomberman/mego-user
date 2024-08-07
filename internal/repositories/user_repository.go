package repositories

import (
	"github.com/antibomberman/mego-user/internal/models"

	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
	"log"
)

type UserRepository interface {
	GetById(string) (*models.User, error)
	GetByEmail(string) (*models.User, error)
	GetByPhone(string) (*models.User, error)
	Find(startIndex int, size int, sort, search string) ([]models.User, error)
	Create(data *models.CreateUserRequest) (*models.User, error)
	Update(id string, data *models.User) error
	Delete(id string) error
	ForceDelete(id string) error
	Count() (int, error)
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

func (r *userRepository) Find(startIndex int, size int, sort, search string) ([]models.User, error) {
	var users []models.User
	query := "SELECT * FROM users"
	log.Println(search)
	if search != "" {
		log.Printf("Search query: %s", query)
		query += " WHERE first_name LIKE '% " + search + " %'"
	}

	switch sort {
	case "0":
		query += " ORDER BY created_at DESC"
	case "1":
		query += " ORDER BY created_at ASC"
	default:
		query += " ORDER BY created_at DESC"
	}

	err := r.db.Select(&users, query+" OFFSET $1 LIMIT $2", startIndex, size)
	if err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return []models.User{}, nil
	}

	return users, nil
}
func (r *userRepository) GetById(id string) (*models.User, error) {
	var user models.User
	err := r.db.Get(&user, "SELECT * FROM users WHERE id = $1", id)
	if err != nil {
		log.Printf("Error getting user by id: %v", err)
		return nil, err
	}
	return &user, nil
}
func (r *userRepository) GetByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Get(&user, "SELECT * FROM users WHERE email = $1", email)
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
		log.Printf("Error getting user by phone: %v", err)
		return nil, err
	}
	return &user, nil
}
func (r *userRepository) Create(data *models.CreateUserRequest) (*models.User, error) {
	query := `
		INSERT INTO users (first_name, middle_name, last_name, email)
        VALUES ($1, $2, $3, $4)
        RETURNING id
    `
	id := ""
	err := r.db.QueryRowx(query, data.FirstName, data.MiddleName, data.LastName, data.Email).Scan(&id)

	if err != nil {
		return nil, err
	}
	return r.GetById(id)

}
func (r *userRepository) Update(id string, user *models.User) error {
	_, err := r.db.Exec("UPDATE users SET first_name = $2, middle_name = $3, last_name = $4, email = $5, phone = $6, avatar = $7,about = $8,theme = $9,lang = 10 WHERE id = $1",
		id, user.FirstName, user.MiddleName, user.LastName, user.Email, user.Phone, user.Avatar, user.About, user.Theme, user.Lang)
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
func (r *userRepository) ForceDelete(id string) error {
	_, err := r.db.Exec("DELETE FROM users WHERE id = $1", id)
	return err
}
func (r *userRepository) Count() (int, error) {
	var count int
	err := r.db.Get(&count, "SELECT COUNT(*) FROM users")
	if err != nil {
		return 0, err
	}
	return count, nil
}
