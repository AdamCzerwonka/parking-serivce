package router

import (
	"net/http"
	"parking-service/internal/handlers"
	"parking-service/internal/repositories"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func New(db *sqlx.DB) *mux.Router {
	r := mux.NewRouter()

	srv := handlers.Server{
		Validate:             validator.New(),
		UserRepository:       repositories.NewDbUserRepository(db),
		EmailTokenRepository: repositories.NewDbEmailTokenRepository(db),
	}

	s := r.PathPrefix("/api/v1").Subrouter()

	s.HandleFunc("/user", srv.HandleCreateUser()).Methods("POST")
    s.HandleFunc("/verifyEmail", srv.HandleVerifyEmail()).Methods(http.MethodGet)

	return r
}
