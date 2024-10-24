package domain

import "golang.org/x/crypto/bcrypt"

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