package profiles

import "gorm.io/gorm"

type ProfilesController struct {
	DB *gorm.DB
}
