package repositories

import (
	"context"
	"parking-service/internal/entities"
)

type UserRepository interface {
	Create(ctx context.Context, firstName, lastName, email, passwordHash string) error
	GetByEmail(ctx context.Context, email string) (*entities.User, error)
}
