package profiles

import (
	"github.com/jnxvi/nyalist/auth"
	"gorm.io/gorm"
)

type Profile struct {
	gorm.Model
	ProfilePic string `gorm:"default:'profiles/default.jpg'"`
	Watchlist  string
	Favorties  string
	Watching   string
	Bio        string
	UserID     int
	User       auth.User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
