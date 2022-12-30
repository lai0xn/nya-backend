package auth

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username  string
	Email     string
	Password  string
	AuthToken string
}
