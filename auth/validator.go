package auth

import (
	"errors"

	"github.com/jnxvi/nyalist/database"
	"github.com/jnxvi/nyalist/models"
)

type Validator struct{}

func (Validator) ValidateEmail(email string) error {
	var user models.User

	database.DB.First(&user, "email = ?", email)

	if user.Email != "" {
		return errors.New("Email already in use")
	}
	return nil
}

func (Validator) ValidateUsername(username string) error {
	var user models.User

	database.DB.First(&user, "username = ?", username)

	if user.Email != "" {
		return errors.New("Username already in use")
	}
	return nil
}
