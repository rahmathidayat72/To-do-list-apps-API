package data

import (
	"errors"
	"log"
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

// Insert implements task.DataTaskInterface.
func (r *QueryTask) Insert(input task.CoreTask, userId uint) error {
	// panic("unimplemented")
	taskGorm := CoreToModel(input)

	tx := r.db.Create(&taskGorm)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("failed to insert, row affected is 0")
	}
	return nil
}

// Update implements task.DataTaskInterface.
func (r *QueryTask) Update(id uint, input task.CoreTask, userId uint) error {
	// panic("unimplemented")

	var updateTask Task
	tx := r.db.Where("id = ? AND user_id = ?", id, userId).First(&updateTask)
	if tx.Error != nil {
		return tx.Error
	}
	//mengecek ada data yang terupdate atau tidak
	if tx.RowsAffected == 0 {
		return errors.New("task id not found")
	}
	updateTask.Name = input.Name
	updateTask.UserId = input.UserId
	updateTask.Description = input.Description

	tx = r.db.Model(&Task{}).Where("id=?", id).Updates(updateTask)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// Status implements task.DataTaskInterface.
func (r *QueryTask) Status(id uint, input task.CoreTask, userId uint) error {
	// panic("unimplemented")

	var updateStatus Task
	tx := r.db.Where("id = ? AND user_id = ?", id, userId).First(&updateStatus, id, userId)
	if tx.Error != nil {
		return tx.Error
	}
	//mengecek ada data yang terupdate atau tidak
	if tx.RowsAffected == 0 {
		return errors.New("Task id not found")
	}

	updateStatus.Status = input.Status

	tx = r.db.Model(&Task{}).Where("id=?", id).Updates(updateStatus)

	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// Delete implements task.DataTaskInterface.
func (r *QueryTask) Delete(Id uint, userId uint) error {
	// panic("unimplemented")
	var deleteTask Task
	// Validasi Id dan userId
	if Id == 0 || userId == 0 {
		return errors.New("Invalid Id or userId")
	}

	tx := r.db.Where("id = ? AND user_id = ?", Id, userId).Delete(&deleteTask, Id)
	if tx.Error != nil {
		return errors.New("failed to delete task" + tx.Error.Error())
	}
	//mengecek ada data yang terupdate atau tidak
	if tx.RowsAffected == 0 {
		return errors.New("Task id not found")
	}
	// Log penghapusan tugas
	log.Printf("Task with Id %d deleted by user %d", Id, userId)

	return nil
}
