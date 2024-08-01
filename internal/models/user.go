package models

type User struct {
	Id         string `db:"id" json:"id"`
	FirstName  string `db:"first_name" json:"first_name"`
	MiddleName string `db:"middle_name" json:"middle_name"`
	LastName   string `db:"last_name" json:"last_name"`
	Email      string `db:"email" json:"email"`
	Password   string `db:"password" json:"-"`
	Phone      string `db:"phone" json:"phone"`
	Avatar     string `db:"avatar" json:"avatar"`
	CreatedAt  string `db:"created_at" json:"created_at"`
	CreatedBy  string `db:"created_by" json:"created_by"`
	UpdatedAt  string `db:"updated_at" json:"updated_at"`
	UpdatedBy  string `db:"updated_by" json:"updated_by"`
	DeletedAt  string `db:"deleted_at" json:"deleted_at"`
	DeletedBy  string `db:"deleted_by" json:"deleted_by"`
}

type UserDetails struct {
	Id         string   `db:"id" json:"id"`
	FirstName  string   `db:"first_name" json:"first_name"`
	MiddleName string   `db:"middle_name" json:"middle_name"`
	LastName   string   `db:"last_name" json:"last_name"`
	Email      string   `db:"email" json:"email"`
	Phone      string   `db:"phone" json:"phone"`
	Avatar     string   `db:"avatar" json:"avatar"`
	CreatedAt  string   `db:"created_at" json:"created_at"`
	UpdatedAt  string   `db:"updated_at" json:"updated_at"`
	DeletedAt  string   `db:"deleted_at" json:"deleted_at"`
	Roles      []string `db:"roles" json:"roles"`
}

type CreateUserRequest struct {
	FirstName  string `json:"first_name"`
	MiddleName string `json:"middle_name"`
	LastName   string `json:"last_name"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Avatar     string `json:"avatar"`
	Password   string `json:"password"`
}

type UpdateUserRequest struct {
	FirstName  string `json:"first_name"`
	MiddleName string `json:"middle_name"`
	LastName   string `json:"last_name"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Avatar     string `json:"avatar"`
	Password   string `json:"password"`
}
