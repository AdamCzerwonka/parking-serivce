package entities

import (
	"database/sql"
	"time"
)

type User struct {
	Id           int           `db:"id"`
	FirstName    string        `db:"first_name"`
	LastName     string        `db:"last_name"`
	Email        string        `db:"email"`
	PasswordHash string        `db:"password_hash"`
	CreatedAt    time.Time     `db:"created_at"`
	UpdatedAt    time.Time     `db:"updated_at"`
	DeletedAt    *sql.NullTime `db:"deleted_at"`
	LastLogin    *sql.NullTime `db:"last_login"`
	Role         string        `db:"role"`
	Enabled      bool          `db:"enabled"`
}
