package repositories

import (
	"context"
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
	firstName, lastName, email, passwordHash string) error {
	sql := `INSERT INTO users(first_name, last_name, email, password_hash)
        VALUES ($1,$2,$2,$4);`

	_, err := r.db.ExecContext(ctx, sql, firstName, lastName, email, passwordHash)
	return err
}

func (r *DbUserRepository) GetByEmail(ctx context.Context, email string) (*entities.User, error) {
    user := &entities.User{}

    sql := `SELECT id, first_name, last_name, email, password_hash, created_at, updated_at, deleted_at FROM users WHERE email=$1`
    err := r.db.QueryRowxContext(ctx,sql,email).StructScan(user)
    if err != nil {
        return nil, err
    }

    return user,nil
}
