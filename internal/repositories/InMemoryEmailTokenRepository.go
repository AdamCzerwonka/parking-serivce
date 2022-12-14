package repositories

import (
	"context"
	"errors"
	"parking-service/internal/entities"
	"time"
)

type InMemoryEmailTokenRepository struct {
	tokens []entities.EmailToken
}

func NewInMemoryEmailTokenRepository() *InMemoryEmailTokenRepository {
	return &InMemoryEmailTokenRepository{}
}

func (r *InMemoryEmailTokenRepository) Create(_ context.Context, user_id int, token string, valid_for time.Duration) error {
	eToken := entities.EmailToken{
		UserId:    user_id,
		Token:     token,
		ValidFrom: time.Now(),
		ValidTo:   time.Now().Add(valid_for),
	}

	r.tokens = append(r.tokens, eToken)
	return nil
}

func (r *InMemoryEmailTokenRepository) Get(_ context.Context, user_id int) (*entities.EmailToken, error) {
	for _, token := range r.tokens {
		if token.UserId == user_id {
			return &token, nil
		}
	}

	return nil, errors.New("Token not found")
}
