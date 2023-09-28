package user

import (
	"time"
)

type CoreUser struct {
	ID          uint
	Name        string
	Email       string
	Password    string
	Address     string
	PhoneNumber string
	CreatedAt   time.Time
	UpdateAt    time.Time
}

type DataUserInterface interface {
	SelectAll() ([]CoreUser, error)
	Insert(insert CoreUser) error
	Update(insert CoreUser, id uint) error
}

type ServiceUserInterface interface {
	GetAll() ([]CoreUser, error)
	CreateUser(insert CoreUser) error
	Update(insert CoreUser, id uint) error
}
