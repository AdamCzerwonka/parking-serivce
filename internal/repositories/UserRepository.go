package repositories

import (
	"context"
	"parking-service/internal/entities"
)

type UserRepository interface {
	Create(ctx context.Context, firstName, lastName, email, passwordHash string) (int, error)
	GetByEmail(ctx context.Context, email string) (*entities.User, error)
	VerifyUser(ctx context.Context, userId int) error
	GetById(ctx context.Context, userId int) (*entities.User, error)
}
