package repositories

import "parking-service/internal/entities"

type UserRepository interface {
	Create(firstName, lastName, email, passwordHash string) error
	GetByEmail(email string) (*entities.User, error)
}
