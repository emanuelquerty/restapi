package app

import "golang.org/x/crypto/bcrypt"

func generateHash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	passwordHash := string(bytes)
	return passwordHash, nil
}
