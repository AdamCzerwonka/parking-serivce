package repositories

import (
	"context"
	"database/sql"
	"errors"
	"parking-service/internal/entities"

	"github.com/jmoiron/sqlx"
)

type DbUserRepository struct {
	db *sqlx.DB
}

func NewDbUserRepository(db *sqlx.DB) *DbUserRepository {
	return &DbUserRepository{
		db: db,
	}
}

func (r *DbUserRepository) Create(ctx context.Context,
	firstName, lastName, email, passwordHash string) (int,error) {
	sql := `INSERT INTO users(first_name, last_name, email, password_hash, role)
        VALUES ($1,$2,$3,$4,$5) RETURNING id;`

    var id int

	 err := r.db.GetContext(ctx, &id, sql, firstName, lastName, email, passwordHash, "user")
	return id,err
}

func (r *DbUserRepository) GetByEmail(ctx context.Context, email string) (*entities.User, error) {
	user := &entities.User{}

	sqlQuery := `SELECT id, first_name, last_name, email, password_hash, created_at, updated_at, deleted_at FROM users WHERE email=$1`
	err := r.db.QueryRowxContext(ctx, sqlQuery, email).StructScan(user)

	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return user, nil
}
