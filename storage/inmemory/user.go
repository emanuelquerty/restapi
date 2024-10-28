package inmemory

import (
	"context"
	"errors"
	"restapi/domain"
	"sync"
)

type UserStore struct {
	users []domain.User
	mu    sync.Mutex
}

func NewUserStore(Users []domain.User) *UserStore {
	return &UserStore{
		users: Users,
	}
}

func (u *UserStore) Create(ctx context.Context, user *domain.User) error {
	u.mu.Lock()
	defer u.mu.Unlock()

	id := len(u.users) + 1
	user.ID = id

	u.users = append(u.users, *user)
	return nil
}

func (u *UserStore) FindByID(ctx context.Context, id int) (domain.User, error) {
	u.mu.Lock()
	defer u.mu.Unlock()

	for _, currUser := range u.users {
		if currUser.ID == id {
			return currUser, nil
		}
	}

	err := errors.New("error finding user: invalid id")
	return domain.User{}, err
}

func (u *UserStore) Update(ctx context.Context, id int, updates map[string]any) (domain.User, error) {
	u.mu.Lock()
	defer u.mu.Unlock()

	isValidID := false
	for _, currUser := range u.users {
		if currUser.ID == id {
			isValidID = true
			break
		}
	}

	if !isValidID {
		return domain.User{}, errors.New("error updating user: invalid id")
	}

	for key, val := range updates {
		switch key {
		case "first_name":
			u.users[id].FirstName = val.(string)
		case "last_name":
			u.users[id].LastName = val.(string)
		case "email":
			u.users[id].Email = val.(string)
		case "password":
			u.users[id].Password = val.(string)
		case "password_hash":
			u.users[id].PasswordHash = val.(string)
		}
	}

	return u.users[id], nil
}

func (u *UserStore) Delete(ctx context.Context, id int) error {
	for _, currUser := range u.users {
		if currUser.ID == id {
			// Mock deletion by updating user at given id with its zero value
			u.users[id] = domain.User{}
			return nil
		}
	}

	err := errors.New("error deleting user: invalid id")
	return err
}

func (u *UserStore) FindAll(ctx context.Context) ([]domain.User, error) {
	return u.users, nil
}
