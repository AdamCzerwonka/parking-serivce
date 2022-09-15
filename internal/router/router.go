package router

import (
	"parking-service/internal/handlers"
	"parking-service/internal/repositories"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

func New() *mux.Router {
	r := mux.NewRouter()

	srv := handlers.Server{
		Validate:       validator.New(),
		UserRepository: repositories.NewInMemoryUserRepository(),
	}

	s := r.PathPrefix("/api/v1").Subrouter()

	s.HandleFunc("/user", srv.HandleCreateUser()).Methods("POST")

	return r
}
