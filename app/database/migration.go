package database

import (
	_taskdata "rahmat/to-do-list-app/features/task/data"
	_userdata "rahmat/to-do-list-app/features/user/data"

	"gorm.io/gorm"
)

func InitMigration(db *gorm.DB) {
	db.AutoMigrate(&_userdata.User{})
	db.AutoMigrate(&_taskdata.Task{})
}
