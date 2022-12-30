package auth

import (
	"errors"

	"github.com/jnxvi/nyalist/database"
)

type Validator struct{}

func (Validator) ValidateEmail(email string) error {
	var user User

	database.DB.First(&user, "email = ?", email)

	if user.Email != "" {
		return errors.New("Email already in use")
	}
	return nil
}

func (Validator) ValidateUsername(username string) error {
	var user User

	database.DB.First(&user, "username = ?", username)

	if user.Email != "" {
		return errors.New("Username already in use")
	}
	return nil
}
