package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (s *Server) HandleDeleteUser() http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)
        userId, isPresent := vars["id"]
        if !isPresent  {
            w.WriteHeader(http.StatusNotFound)
            return
        }

        id, err := strconv.ParseInt(userId, 10,32)
        if err != nil {
            errorResponse(w, []string{"Wrong id"}, http.StatusBadRequest)
        }
        err = s.UserRepository.Delete(r.Context(), int(id))
        if err != nil {
            errorResponse(w, []string{"Something went wrong"}, http.StatusInternalServerError)
            log.Println(err)
            return
        }

        w.WriteHeader(http.StatusNotFound)

    }
}
