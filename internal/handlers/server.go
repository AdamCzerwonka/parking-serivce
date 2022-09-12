package handlers

import "github.com/go-playground/validator/v10"

type Server struct {
    Validate *validator.Validate
}
