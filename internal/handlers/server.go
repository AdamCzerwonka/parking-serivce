package handlers

import (
	"parking-service/internal/repositories"

	"github.com/go-playground/validator/v10"
)

type Server struct {
	Validate             *validator.Validate
	UserRepository       repositories.UserRepository
	EmailTokenRepository repositories.EmailTokenRepository
}
