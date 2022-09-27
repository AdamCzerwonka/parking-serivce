package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func (s *Server) HandleLogin() http.HandlerFunc {
    type LoginInput struct {
        Email string `json:"email"`
        Password string `json:"password"`
    }

    jwtKey := []byte("adjahjdbnasuhjbda9102u3iuy12g837ghuHUI!H!UIB#U!")

    return func(w http.ResponseWriter, r *http.Request) {
        input := LoginInput{}
        err := json.NewDecoder(r.Body).Decode(&input)
        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            log.Println(err)
            return
        }

        user, err := s.UserRepository.GetByEmail(r.Context(), input.Email)
        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            log.Println(err)
            return
        }

        if user == nil {
            w.WriteHeader(http.StatusBadRequest)
            return
        }

        err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(input.Password))
        if err != nil {
            errorResponse(w, []string{"Wrong credentials"}, http.StatusBadRequest)
            return
        }

        if err = s.UserRepository.CheckIfVerified(r.Context(), user.Id); err != nil {
            jsonResponse(w, "Account not enabled. Please verify your account", http.StatusBadRequest)
            return
        }

        expirationTime:= time.Now().Add(24 * time.Hour)

        claims := claims {
            Id: user.Id,
            StandardClaims: jwt.StandardClaims{
                ExpiresAt: expirationTime.Unix(),
            },
        }

        token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
        tokenString, err := token.SignedString(jwtKey)
        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            log.Println(err)
            return
        }

        err = s.UserRepository.UpdateLogin(r.Context(), user.Id)
        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            log.Println(err)
            return
        }
        jsonResponse(w, tokenString, http.StatusOK)

    }
}

type claims struct {
    Id int `json:"id"`
    jwt.StandardClaims
}
