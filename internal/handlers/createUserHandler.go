package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func (s *Server) HandleCreateUser() http.HandlerFunc {
    type createUserInput struct {
        Email string `json:"email" validate:"required"`
        FirstName string `json:"firstName" validate:"required"`
        LastName string `json:"lastName" validate:"required"`
        Password string `json:"password" validate:"required"` 
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
               errors= append(errors, fmt.Sprintf("Field: %s is required",err.Field()))
            }

            errorResponse(w, errors, http.StatusBadRequest)
            return
        }

        log.Println(input)
    }
}
