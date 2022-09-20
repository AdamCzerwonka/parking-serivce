package repositories

import (
	"context"
	"parking-service/internal/entities"
	"time"
)

type EmailTokenRepository interface {
	Create(ctx context.Context, user_id int, token string, valid_for time.Duration) error
	Get(ctx context.Context, user_id int) (*entities.EmailToken, error)
}
