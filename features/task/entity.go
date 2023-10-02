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
}

type ServiceTaskInterface interface {
	GetAll(userId uint) ([]CoreTask, error)
}
