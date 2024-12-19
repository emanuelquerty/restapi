package domain

import (
	"context"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID           int    `json:"id,omitempty"`
	FirstName    string `json:"first_name,omitempty"`
	LastName     string `json:"last_name,omitempty"`
	Email        string `json:"email,omitempty"`
	Password     string `json:"password,omitempty"`
	PasswordHash string `json:"password_hash,omitempty"`
}

func (u *User) HashPassword() error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.PasswordHash = string(bytes)
	return nil
}

type UserService interface {
	Create(ctx context.Context, user *User) error
	Update(ctx context.Context, id int, updates map[string]any) (User, error)
	FindByID(ctx context.Context, id int) (User, error)
	Delete(ctx context.Context, id int) error
	FindAll(ctx context.Context) ([]User, error)
}

// Public user hides secret fields such as email and password
type PublicUser struct {
	ID        int    `json:"id,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
}
