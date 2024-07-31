package models

type User struct {
	Id         int    `db:"id" json:"id"`
	FirstName  string `db:"first_name" json:"first_name"`
	MiddleName string `db:"middle_name" json:"middle_name"`
	Email      string `db:"email" json:"email"`
	Password   string `db:"password" json:"-"`
	Avatar     string `db:"avatar" json:"avatar"`
}
