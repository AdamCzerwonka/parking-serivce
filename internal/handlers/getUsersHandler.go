package handlers

import (
	"log"
	"net/http"
	"strconv"
	"time"
)

func (s *Server) HandleGetUsers() http.HandlerFunc {
	type output struct {
		Id        int        `json:"id"`
		FirstName string     `json:"firstName"`
		LastName  string     `json:"LastName"`
		Email     string     `json:"email"`
		Role      string     `json:"role"`
		LastLogin *time.Time `json:"lastLogin"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		queries := r.URL.Query()
		pageSize := 10
		page := 1
		pageStr, isPresent := queries["page"]
		if isPresent {
			pageCon, err := strconv.ParseInt(pageStr[0], 10, 32)
			if err != nil {
				errorResponse(w, []string{"Wrong page format"}, http.StatusBadRequest)
				log.Println(err)
				return
			}
			page = int(pageCon)
		}

		users, err := s.UserRepository.Get(r.Context(), page, pageSize)
		if err != nil {
			log.Println(err)
			errorResponse(w, []string{"Something went wrong"}, http.StatusBadRequest)
			return
		}

		if users == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		resultUsers := []output{}

		for _, user := range users {
			var lastLogin *time.Time
			if user.LastLogin != nil {
				lastLogin = &user.LastLogin.Time
			} else {
				lastLogin = nil
			}

			outputUser := output{
				Id:        user.Id,
				FirstName: user.FirstName,
				LastName:  user.LastName,
				Email:     user.Email,
				Role:      user.Role,
				LastLogin: lastLogin,
			}

			resultUsers = append(resultUsers, outputUser)
		}

		jsonResponse(w, resultUsers, http.StatusOK)

	}
}
