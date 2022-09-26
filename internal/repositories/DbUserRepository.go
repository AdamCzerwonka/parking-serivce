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

func (r *DbUserRepository) Get(ctx context.Context, page, pageSize int) ([]*entities.User, error) {
	sqlQuery := `SELECT id, first_name, last_name, email, password_hash, created_at, updated_at, deleted_at, role, last_login FROM users LIMIT $1 OFFSET $2;`
	toSkip := pageSize * (page - 1)

	rows, err := r.db.QueryxContext(ctx, sqlQuery, pageSize, toSkip)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	result := []*entities.User{}

	for rows.Next() {
		p := &entities.User{}
		err = rows.StructScan(p)
		if err != nil {
			return nil, err
		}
		result = append(result, p)
	}

	return result, nil
}

func (r *DbUserRepository) GetById(ctx context.Context, userId int) (*entities.User, error) {
	sqlQuery := `SELECT id, first_name, last_name, email, password_hash, created_at, updated_at, deleted_at,role,last_login FROM users WHERE id=$1`
	user := &entities.User{}

	err := r.db.QueryRowxContext(ctx, sqlQuery, userId).StructScan(user)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return user, nil

}

func (r *DbUserRepository) Create(ctx context.Context,
	firstName, lastName, email, passwordHash string) (int, error) {
	sql := `INSERT INTO users(first_name, last_name, email, password_hash, role)
        VALUES ($1,$2,$3,$4,$5) RETURNING id;`

	var id int

	err := r.db.GetContext(ctx, &id, sql, firstName, lastName, email, passwordHash, "user")
	return id, err
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

func (r *DbUserRepository) VerifyUser(ctx context.Context, userId int) error {
	sqlQuery := `UPDATE users SET enabled=true WHERE id=$1;`

	_, err := r.db.ExecContext(ctx, sqlQuery, userId)
	return err
}
