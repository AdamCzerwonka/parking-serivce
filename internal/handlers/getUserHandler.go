package handlers

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (s *Server) HandleGetUser() http.HandlerFunc {
	type output struct {
		Id        int           `json:"id"`
		FirstName string        `json:"firstName"`
		LastName  string        `json:"lastName"`
		Email     string        `json:"email"`
		LastLogin *sql.NullTime `json:"lastlogin"`
		Role      string        `json:"role"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		userId, err := strconv.Atoi(vars["id"])
		if err != nil {
			log.Println(err)
			errorResponse(w, []string{"Invalid user id passed as a paramter"}, http.StatusBadRequest)
			return
		}

		user, err := s.UserRepository.GetById(r.Context(), userId)
		if err != nil {
			log.Println(err)
			errorResponse(w, []string{err.Error()}, http.StatusInternalServerError)
			return
		}

		if user == nil {
            w.WriteHeader(http.StatusNotFound)
			return
		}

		result := output{
			Id:        user.Id,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
			LastLogin: user.LastLogin,
			Role:      user.Role,
		}

		jsonResponse(w, result, http.StatusOK)
	}
}
