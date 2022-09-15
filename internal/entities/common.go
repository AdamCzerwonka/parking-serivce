package entities

import "database/sql"

type common struct {
	CreatedAt *sql.NullTime
	UpdatedAt *sql.NullTime
	DeletedAt *sql.NullTime
}
