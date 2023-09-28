package database

import (
	_userdata "rahmat/to-do-list-app/features/user/data"

	"gorm.io/gorm"
)

func InitMigration(db *gorm.DB) {
	db.AutoMigrate(&_userdata.User{})
}
