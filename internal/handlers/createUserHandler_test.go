package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"parking-service/internal/repositories"
	"strings"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/matryer/is"
)

func TestCreateUserHandler(t *testing.T) {
	type input struct {
		Email     string `json:"email"`
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		Password  string `json:"password"`
		Password2 string `json:"password2"`
	}

	tests := []struct {
		Name   string
		Data   *input
		result int
	}{
		{"CorrectData", &input{Email: "johnd@example.com", FirstName: "John", LastName: "Doe", Password: "Test123", Password2: "Test123"}, 201},
		{"AllFieldsMissingFields", &input{}, 400},
		{"EmptyData", nil, 400},
		{"PasswordDoesNotMatch", &input{Email: "johnd@example.com", FirstName: "John", LastName: "Doe", Password: "Test123", Password2: "Test"}, 400},
		{"WrongEmailFormat", &input{Email: "johnd", FirstName: "John", LastName: "Doe", Password: "Test123", Password2: "Test123"}, 400},
		{"UserAlreadyExists", &input{Email: "johnd@example.com", FirstName: "John", LastName: "Doe", Password: "Test123", Password2: "Test123"}, 400},
	}

	is := is.New(t)
	server := Server{
		Validate:       validator.New(),
		UserRepository: repositories.NewInMemoryUserRepository(),
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			iis := is.New(t)

			json, err := json.Marshal(test.Data)
			if err != nil {
				iis.Fail()
			}
			req, err := http.NewRequest("POST", "/api/v1/user", strings.NewReader(string(json)))
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(server.HandleCreateUser())

			handler.ServeHTTP(rr, req)

			iis.Equal(test.result, rr.Code)

		})
	}

}
