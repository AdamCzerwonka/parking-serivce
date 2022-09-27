package repositories

import (
	"context"
	"parking-service/internal/entities"
	"time"
)

type InMemoryUserRepository struct {
	users []*entities.User
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{users: []*entities.User{}}
}

func (r *InMemoryUserRepository) Delete(_ context.Context, userId int) error {
    var idx int
    for i,user := range r.users {
        if user.Id == userId {
            idx = i
            break
        }
    }

    r.users[idx] = r.users[len(r.users)-1]
    r.users = r.users[:len(r.users)-1] 

    return nil
}

func (r *InMemoryUserRepository) Get(_ context.Context, page, pageSize int) ([]*entities.User, error) {
	toSkip := pageSize * (page-1)
    if len(r.users) == 0 {
        return nil, nil
    }

    end := toSkip + pageSize
    if len(r.users) < toSkip + pageSize {
        end = len(r.users)
    }
	return r.users[toSkip : end], nil
}
func (r *InMemoryUserRepository) GetById(_ context.Context, userId int) (*entities.User, error) {
	for _, user := range r.users {
		if user.Id == userId {
			return user, nil
		}
	}

	return nil, nil

}

func (r *InMemoryUserRepository) Create(_ context.Context, firstName, lastName, email, passwordHash string) (int, error) {
	user := entities.User{
		Id:           len(r.users),
		FirstName:    firstName,
		LastName:     lastName,
		Email:        email,
		PasswordHash: passwordHash,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		DeletedAt:    nil,
	}

	r.users = append(r.users, &user)
	return user.Id, nil
}

func (r *InMemoryUserRepository) GetByEmail(_ context.Context, email string) (*entities.User, error) {
	for _, user := range r.users {
		if user.Email == email {
			return user, nil
		}
	}

	return nil, nil
}

func (r *InMemoryUserRepository) VerifyUser(_ context.Context, userId int) error {
	for _, user := range r.users {
		if user.Id == userId {
			user.Enabled = true
			break
		}
	}
	return nil
}
