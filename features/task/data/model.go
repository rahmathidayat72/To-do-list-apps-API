package data

import (
	"rahmat/to-do-list-app/features/task"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Name        string `json:"name" gorm:"not null"`
	UserId      uint   `json:"user_id" gorm:"not null"`
	Description string `json:"description" gorm:"not null"`
	Status      string `json:"status" gorm:"enum('Not Completed','Completed')"`
}

func CoreToModel(dataCore task.CoreTask) Task {
	return Task{

		Name:        dataCore.Name,
		UserId:      dataCore.UserId,
		Description: dataCore.Description,
		Status:      dataCore.Status,
	}
}

func ModelToCore(dataMode Task) task.CoreTask {
	return task.CoreTask{
		ID:          dataMode.ID,
		Name:        dataMode.Name,
		UserId:      dataMode.UserId,
		Description: dataMode.Description,
		Status:      dataMode.Status,
		CreatedAt:   dataMode.CreatedAt,
		UpdateAt:    dataMode.UpdatedAt,
	}
}
