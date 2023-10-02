package user

import (
	"rahmat/to-do-list-app/features/task"
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
	Task        []task.CoreTask
}

type DataUserInterface interface {
	SelectAll() ([]CoreUser, error)
	Insert(insert CoreUser) error
	Update(insert CoreUser, id uint) error
	SelectById(id uint) (CoreUser, error)
	Delete(id uint) error
	Login(email, password string) (dataLogin CoreUser, err error)
}

type ServiceUserInterface interface {
	GetAll() ([]CoreUser, error)
	CreateUser(insert CoreUser) error
	Update(insert CoreUser, id uint) error
	SelectById(id uint) (CoreUser, error)
	Delete(id uint) error
	Login(email, password string) (dataLogin CoreUser, token string, err error)
}
