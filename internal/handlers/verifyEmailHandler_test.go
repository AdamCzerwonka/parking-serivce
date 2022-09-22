package handlers

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"parking-service/internal/repositories"
	"testing"
	"time"

	"github.com/matryer/is"
)

func TestHandleVerifyEmail(t *testing.T) {
	// Preapare mock users accounts
	srv := Server{
		UserRepository:       repositories.NewInMemoryUserRepository(),
		EmailTokenRepository: repositories.NewInMemoryEmailTokenRepository(),
	}

	ctx := context.Background()

	tests := []struct {
		name   string
		token  string
		result int
	}{
		{
			name:   "EmptyToken",
			token:  "",
			result: http.StatusInternalServerError,
		},
		{
			name:   "IncorrectTokenFormat",
			token:  "asdasdquiweh1y312831",
			result: http.StatusInternalServerError,
		},
		{
			name:   "UserDoesNotExists",
			token:  "eyJpZCI6MTUsInRva2VuIjoiYXNkYXNkc2Fkc2EyMTkzMjEzODcyMTkwMzIxIn0=",
			result: http.StatusBadRequest,
		},
		{
			name:   "UserExistsButInvalidToken",
			token:  "eyJpZCI6MCwidG9rZW4iOiJhc2Rhc2RzYWRzYTIxOTMyMTM4NzIxOTAzMjEifQ==",
			result: http.StatusBadRequest,
		},
		{
			name: "CorrectToken",
			token: func() string {
				userId, _ := srv.UserRepository.Create(ctx, "Test", "Test", "Test@test.com", "hash")

				token := genereteToken()
				verToken, _ := createEmailConfirmationToken(userId, token)
				srv.EmailTokenRepository.Create(ctx, userId, token, time.Hour*24)

				return verToken
			}(),
			result: http.StatusOK,
		},
		{
			name: "TokenTooOld",
			token: func() string {
				userId, _ := srv.UserRepository.Create(ctx, "Test", "Test", "Test@test.com", "hash")

				token := genereteToken()
				verToken, _ := createEmailConfirmationToken(userId, token)
				srv.EmailTokenRepository.Create(ctx, userId, token, time.Millisecond)

				time.Sleep(time.Millisecond * 10)

				return verToken
			}(),
			result: http.StatusBadRequest,
		},
	}

	is := is.New(t)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			iss := is.New(t)
			req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/verifyEmail?token=%s", test.token), nil)
			if err != nil {
				iss.Fail()
			}

			rr := httptest.NewRecorder()
			handler := srv.HandleVerifyEmail()

			handler.ServeHTTP(rr, req)

			iss.Equal(test.result, rr.Code)
		})
	}

}
