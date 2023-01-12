package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username  string
	Email     string
	Password  string
	AuthToken string
	Profile   Profile
}

func (u *User) AfterCreate(tx *gorm.DB) (err error) {
	tx.Create(&Profile{UserID: int(u.ID)})
	return
}
