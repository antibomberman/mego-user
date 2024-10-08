package models

import (
	"database/sql"
	"time"
)

type User struct {
	Id         string         `db:"id" json:"id"`
	FirstName  sql.NullString `db:"first_name" json:"first_name"`
	MiddleName sql.NullString `db:"middle_name" json:"middle_name"`
	LastName   sql.NullString `db:"last_name" json:"last_name"`
	Email      sql.NullString `db:"email" json:"email"`
	Password   sql.NullString `db:"password" json:"-"`
	Phone      sql.NullString `db:"phone" json:"phone"`
	Avatar     sql.NullString `db:"avatar" json:"avatar"`
	About      sql.NullString `db:"about" json:"about"`
	Theme      sql.NullString `db:"theme" json:"theme"`
	Lang       sql.NullString `db:"lang" json:"lang"`
	CreatedAt  sql.NullTime   `db:"created_at" json:"created_at"`
	UpdatedAt  sql.NullTime   `db:"updated_at" json:"updated_at"`
	DeletedAt  sql.NullTime   `db:"deleted_at" json:"deleted_at"`
}

type UserDetails struct {
	Id         string    `db:"id" json:"id"`
	FirstName  string    `db:"first_name" json:"first_name"`
	MiddleName string    `db:"middle_name" json:"middle_name"`
	LastName   string    `db:"last_name" json:"last_name"`
	Email      string    `db:"email" json:"email"`
	Phone      string    `db:"phone" json:"phone"`
	Avatar     *Avatar   `db:"avatar" json:"avatar"`
	About      string    `db:"about" json:"about"`
	Theme      string    `db:"theme" json:"theme"`
	Lang       string    `db:"lang" json:"lang"`
	CreatedAt  time.Time `db:"created_at" json:"created_at"`
	UpdatedAt  time.Time `db:"updated_at" json:"updated_at"`
	DeletedAt  time.Time `db:"deleted_at" json:"deleted_at"`
	Roles      []string  `db:"roles" json:"roles"`
}

type CreateUserRequest struct {
	FirstName  string `json:"first_name"`
	MiddleName string `json:"middle_name"`
	LastName   string `json:"last_name"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	About      string `json:"about"`
	Theme      string `db:"theme" json:"theme"`
	Lang       string `db:"lang" json:"lang"`
}

type UpdateUserRequest struct {
	FirstName  string     `json:"first_name"`
	MiddleName string     `json:"middle_name"`
	LastName   string     `json:"last_name"`
	Email      string     `json:"email"`
	Phone      string     `json:"phone"`
	Avatar     *NewAvatar `json:"avatar"`
	About      string     `json:"about"`
	Theme      string     `json:"theme"`
	Lang       string     `json:"lang"`
}
type NewAvatar struct {
	FileName    string `json:"file_name"`
	ContentType string `json:"content_type"`
	Data        []byte `json:"data"`
}
type Avatar struct {
	FileName string `json:"file_name"`
	Url      string `json:"url"`
}
