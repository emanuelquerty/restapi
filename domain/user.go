package domain

import (
	"context"
)

type User struct {
	ID           int    `json:"id,omitempty"`
	FirstName    string `json:"first_name,omitempty"`
	LastName     string `json:"last_name,omitempty"`
	Email        string `json:"email,omitempty"`
	Password     string `json:"password,omitempty"`
	PasswordHash string `json:"password_hash,omitempty"`
}

type UserService interface {
	Create(ctx context.Context, user *User) error
	Update(ctx context.Context, id int, updates map[string]any) (User, error)
	FindByID(ctx context.Context, id int) (User, error)
	Delete(ctx context.Context, id int) error
	FindAll(ctx context.Context) ([]User, error)
}
