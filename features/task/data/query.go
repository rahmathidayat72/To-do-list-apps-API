package data

import (
	"rahmat/to-do-list-app/features/task"

	"gorm.io/gorm"
)

type QueryTask struct {
	db *gorm.DB
}

func New(db *gorm.DB) task.DataTaskInterface {
	return &QueryTask{
		db: db,
	}
}

// SelectAll implements task.DataTaskInterface.
func (r *QueryTask) SelectAll(userId uint) ([]task.CoreTask, error) {
	// panic("unimplemented")
	var taskData []Task
	tx := r.db.Where("user_id", userId).Find(&taskData)
	if tx.Error != nil {
		return nil, tx.Error
	}

	var taskCore []task.CoreTask
	for _, v := range taskData {
		task := ModelToCore(v)
		taskCore = append(taskCore, task)
	}
	return taskCore, nil
}
