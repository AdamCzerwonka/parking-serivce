package entities

import "database/sql"

type User struct {
	Id           int
	FirstName    string
	LastName     string
	Email        string
	PasswordHash string
	CreatedAt    *sql.NullTime
	UpdatedAt    *sql.NullTime
	DeletedAt    *sql.NullTime
}
