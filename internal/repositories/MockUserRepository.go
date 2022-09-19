package repositories

import (
	"context"
	"database/sql"
	"parking-service/internal/entities"
	"time"
)

type InMemoryUserRepository struct {
	users []entities.User
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{users: []entities.User{}}
}

func (r *InMemoryUserRepository) Create(_ context.Context, firstName, lastName, email, passwordHash string) error {
	user := entities.User{
		Id:           len(r.users),
		FirstName:    firstName,
		LastName:     lastName,
		Email:        email,
		PasswordHash: passwordHash,
		CreatedAt:    &sql.NullTime{Time: time.Now()},
		UpdatedAt:    &sql.NullTime{Time: time.Now()},
		DeletedAt:    nil,
	}

	r.users = append(r.users, user)
	return nil
}

func (r *InMemoryUserRepository) GetByEmail(_ context.Context, email string) (*entities.User, error) {
	for _, user := range r.users {
		if user.Email == email {
			return &user, nil
		}
	}

	return nil, nil
}
