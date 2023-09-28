package data

import (
	"rahmat/to-do-list-app/features/user"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	// ID          uint `gorm:"primaryKey"`
	// CreatedAt   time.Time
	// UpdateAt   time.Time
	// DeletedAt  gorm.DeletedAt `gorm:"index"`
	Name        string
	Email       string `gorm:"unique"`
	Password    string
	Address     string
	PhoneNumber string
	// Task        []Task `json:"task"`
}

func CoreToModel(dataCore user.CoreUser) User {
	return User{
		Name:        dataCore.Name,
		Email:       dataCore.Email,
		Password:    dataCore.Password,
		Address:     dataCore.Address,
		PhoneNumber: dataCore.PhoneNumber,
	}
}

func ModelToCore(dataModel User) user.CoreUser {
	return user.CoreUser{
		ID:          dataModel.ID,
		Name:        dataModel.Name,
		Email:       dataModel.Email,
		Password:    dataModel.Password,
		Address:     dataModel.Address,
		PhoneNumber: dataModel.PhoneNumber,
		CreatedAt:   dataModel.CreatedAt,
		UpdateAt:    dataModel.UpdatedAt,
	}
}
