package handlers

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"parking-service/internal/repositories"
	"testing"

	"github.com/gorilla/mux"
	"github.com/matryer/is"
)

func TestHandleGetUser(t *testing.T) {

	srv := Server{
		UserRepository: repositories.NewInMemoryUserRepository(),
	}

	tests := []struct {
		name   string
		id     string
		result int
	}{
		{
			name:   "InvalidIdFormat",
			id:     "asdnjasbdas",
			result: http.StatusBadRequest,
		},
		{
			name: "ExistingId",
			id: func() string {
				id, _ := srv.UserRepository.Create(context.Background(), "Test", "test", "test", "test")
				return fmt.Sprint(id)
			}(),
			result: http.StatusOK,
		},
	}

	is := is.New(t)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			iss := is.New(t)
			req, _ := http.NewRequest(http.MethodGet, "/api/v1/user/"+test.id, nil)

			vars := map[string]string{
				"id": test.id,
			}

			req = mux.SetURLVars(req, vars)

			handler := srv.HandleGetUser()

			rr := httptest.NewRecorder()

			handler.ServeHTTP(rr, req)

			iss.Equal(test.result, rr.Code)
		})
	}

}
