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

func TestHandleGetUsers(t *testing.T) {
    is := is.New(t)


    tests := []struct {
        name string
        page int
        result int
        repo *repositories.InMemoryUserRepository
    } {
        {
            name: "NoUsersShould404",
            page: 1,
            result: 404,
            repo: func() *repositories.InMemoryUserRepository {
                tmp := repositories.NewInMemoryUserRepository()
                return tmp
            }(),
        },
        {
            name: "SomeUsersShould200",
            page: 1,
            result: 200,
            repo: func() *repositories.InMemoryUserRepository {
                tmp := repositories.NewInMemoryUserRepository()
                tmp.Create(context.Background(), "test","test","test","test")
                return tmp
            }(),
        },
    }

    for _,test := range tests {
        t.Run(test.name, func(t *testing.T) {
            srv := Server {
                UserRepository: test.repo,
            }
            is := is.New(t)
            req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/user?page=%d",test.page), nil)
            is.NoErr(err)

            req = mux.SetURLVars(req, map[string]string{"page":fmt.Sprint(test.page)})
            
            handler := srv.HandleGetUsers()

            rr := httptest.NewRecorder()

            handler.ServeHTTP(rr,req)

            is.Equal(test.result, rr.Code)

        })
    }
}
