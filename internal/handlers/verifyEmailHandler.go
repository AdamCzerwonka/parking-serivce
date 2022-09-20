package handlers

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func (s *Server) HandleVerifyEmail() http.HandlerFunc {
    type TokenInput struct {
        UserId int `json:"userId"`
        Token string `json:"token"`
    }

    return func(w http.ResponseWriter, r *http.Request) {
        vars := r.URL.Query()
        log.Println(vars)
        token,isPresent := vars["token"]
        if !isPresent {
            w.WriteHeader(http.StatusBadRequest)
        }

        decodedToken, err := base64.StdEncoding.DecodeString(token[0])
        if err != nil {
            log.Println(err)
            errorResponse(w, []string{"Something went wrong while processing your request"}, http.StatusInternalServerError)
            return
        }
        input := TokenInput{}

        err = json.Unmarshal(decodedToken, &input)
        if err != nil {
            log.Println(err)
            errorResponse(w, []string{"Something went wrong while processing your request"}, http.StatusInternalServerError)
            return
        }

        dbToken, err := s.EmailTokenRepository.Get(r.Context(), input.UserId)

        if dbToken == nil {
            log.Println("There was no token for given user id")
            errorResponse(w, []string{"There was a problem with checking your token"}, http.StatusBadRequest) 
            return
        }
        
        if err != nil {
            log.Println("Error during obteining token from the db: ", err)
            errorResponse(w, []string{"There was a problem with checking your token"}, http.StatusBadRequest) 
            return
        }

        if dbToken.Token != input.Token {
            response := "Incorrect token passed for email verification"
            log.Println(response)
            errorResponse(w, []string{response}, http.StatusBadRequest)
            return
        }

        if time.Now().After(dbToken.ValidTo) {
            errorResponse(w, []string{"Token expired"},http.StatusBadRequest)
            return
        }

        err = s.UserRepository.VerifyUser(r.Context(), dbToken.UserId)
        if err != nil {
            log.Println(err)
            errorResponse(w, []string{"Something went wrong while processing your request"}, http.StatusInternalServerError)
            return
        }
    }
}
