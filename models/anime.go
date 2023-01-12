package models

import "gorm.io/gorm"

type Anime struct {
	gorm.Model
	MalID     int
	Title     string
	Image     string
	Status    string
	ProfileID int
}
