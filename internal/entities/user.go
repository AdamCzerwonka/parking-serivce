package entities

import "database/sql"

type User struct {
	Id           int           `db:"id"`
	FirstName    string        `db:"first_name"`
	LastName     string        `db:"last_name"`
	Email        string        `db:"email"`
	PasswordHash string        `db:"password_hash"`
	CreatedAt    *sql.NullTime `db:"created_at"`
	UpdatedAt    *sql.NullTime `db:"updated_at"`
	DeletedAt    *sql.NullTime `db:"deleted_at"`
}
