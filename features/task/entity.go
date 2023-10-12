package task

import "time"

type CoreTask struct {
	ID          uint
	Name        string
	UserId      uint
	Description string
	Status      string
	CreatedAt   time.Time
	UpdateAt    time.Time
}

type DataTaskInterface interface {
	SelectAll(userId uint) ([]CoreTask, error)
	Insert(input CoreTask, userId uint) error
	Update(id uint, input CoreTask, userId uint) error
	Status(id uint, input CoreTask, userId uint) error
}

type ServiceTaskInterface interface {
	GetAll(userId uint) ([]CoreTask, error)
	Create(input CoreTask, userId uint) error
	Update(id uint, input CoreTask, userId uint) error
	Status(id uint, input CoreTask, userId uint) error
}
