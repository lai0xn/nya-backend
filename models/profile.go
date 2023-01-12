package models

import (
	"gorm.io/gorm"
)

type Profile struct {
	gorm.Model
	ProfilePic string  `gorm:"default:'profiles/default.jpg'"`
	Watchlist  []Anime ///we store a slice of mal_IDS

	Bio    string
	UserID int
}
