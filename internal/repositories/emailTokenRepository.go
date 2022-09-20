package repositories

import (
	"context"
	"time"
)

type EmailTokenRepository interface {
	Create(ctx context.Context, user_id int, token string, valid_for time.Duration) error
	Get(ctx context.Context, user_id int) (string, error)
}
