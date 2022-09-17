package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

func (s *Server) HandleCreateUser() http.HandlerFunc {
	type createUserInput struct {
		Email     string `json:"email" validate:"required,email"`
		FirstName string `json:"firstName" validate:"required"`
		LastName  string `json:"lastName" validate:"required"`
		Password  string `json:"password" validate:"required"`
		Password2 string `json:"password2" validate:"required"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		input := &createUserInput{}
		if err := json.NewDecoder(r.Body).Decode(input); err != nil {
			log.Printf("Could not decode request body. Err: %s", err)
			errorResponse(w, []string{"Ther was an error during processing the request"}, http.StatusInternalServerError)
			return
		}

		err := s.Validate.Struct(input)
		if err != nil {
			errors := []string{}

			for _, err := range err.(validator.ValidationErrors) {
				if err.Tag() == "required" {
					errors = append(errors, fmt.Sprintf("Field: %s is required", err.Field()))
				} else if err.Tag() == "email" {
					errors = append(errors, "Incorrect email format")
				}
			}

			errorResponse(w, errors, http.StatusBadRequest)
			return
		}

		user, err := s.UserRepository.GetByEmail(r.Context(),input.Email)
		if user != nil {
			errorResponse(w, []string{"User with given email already exists"}, http.StatusBadRequest)
			return
		}
		if err != nil {
			log.Println(err)
			errorResponse(w, []string{"Something went wrong while processing your request"}, http.StatusInternalServerError)
            return
		}
		if input.Password != input.Password2 {
			errorResponse(w, []string{"Passwords does not match"}, http.StatusBadRequest)
			return
		}

		hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), 12)
		if err != nil {
			errorResponse(w, []string{"Something went wrong while processing your request"}, http.StatusInternalServerError)
			return
		}

		err = s.UserRepository.Create(r.Context(),input.FirstName, input.LastName, input.Email, string(hash))
		if err != nil {
			errorResponse(w, []string{"Something went wrong while processing your request"}, http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)

	}
}
