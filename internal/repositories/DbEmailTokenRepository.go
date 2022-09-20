package repositories

import (
	"context"
	"database/sql"
	"errors"
	"parking-service/internal/entities"
	"time"

	"github.com/jmoiron/sqlx"
)

type DbEmailTokenRepository struct {
	db *sqlx.DB
}

func NewDbEmailTokenRepository(db *sqlx.DB) *DbEmailTokenRepository {
	return &DbEmailTokenRepository{
		db: db}
}

func (r *DbEmailTokenRepository) Create(ctx context.Context, user_id int, token string, valid_for time.Duration) error {
	sqlQuery := `INSERT INTO email_tokens (user_id, token, valid_from, valid_to) VALUES ($1,$2,$3,$4);`
	valid_from := time.Now()
	valid_to := time.Now().Add(valid_for)

	_, err := r.db.ExecContext(ctx, sqlQuery, user_id, token, valid_from, valid_to)
	if err != nil {
		return err
	}

	return nil
}

func (r *DbEmailTokenRepository) Get(ctx context.Context, user_id int) (*entities.EmailToken, error) {
    sqlQuery := `SELECT user_id, token, valid_from,valid_to FROM email_tokens WHERE user_id=$1`
    token := entities.EmailToken{} 

    err := r.db.QueryRowxContext(ctx, sqlQuery,user_id).StructScan(&token)
    if err !=nil && errors.Is(err, sql.ErrNoRows) {
        return nil,nil
    }

    if err != nil {
        return nil, err
    }

    return &token, nil
    
}
