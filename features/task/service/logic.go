package service

import (
	"errors"
	"rahmat/to-do-list-app/features/task"
)

type TaskService struct {
	taskData task.DataTaskInterface
}

func NewTaskService(repo task.DataTaskInterface) task.ServiceTaskInterface {
	return &TaskService{
		taskData: repo,
	}
}

// GetAll implements task.ServiceTaskInterface.
func (s *TaskService) GetAll(userId uint) ([]task.CoreTask, error) {
	// panic("unimplemented")

	result, err := s.taskData.SelectAll(userId)
	return result, err
}

// Create implements task.ServiceTaskInterface.
func (s *TaskService) Create(input task.CoreTask, userId uint) error {
	// panic("unimplemesnted")
	// Cek apakah user sudah login atau belum
	if userId == 0 {
		return errors.New("user not logged in")
	}

	// // Pastikan userID dalam token sesuai dengan userID dalam proyek
	if userId != input.UserId {
		return errors.New("user does not have access to this task")
	}

	if input.Status != "Completed" && input.Status != "Not Completed" {
		return errors.New("incorrect input, can only accept complete and incomplete input")
	}
	err := s.taskData.Insert(input, userId)
	return err

}

// Update implements task.ServiceTaskInterface.
func (s *TaskService) Update(id uint, input task.CoreTask, userId uint) error {
	// panic("unimplemented")
	// Cek apakah user sudah login atau belum
	if userId == 0 {
		return errors.New("user not logged in")
	}

	// // Pastikan userID dalam token sesuai dengan userID dalam proyek
	if userId != input.UserId {
		return errors.New("user does not have access to this task")
	}
	err := s.taskData.Update(id, input, userId)
	return err
}

// Status implements task.ServiceTaskInterface.
func (s *TaskService) Status(id uint, input task.CoreTask, userId uint) error {
	// panic("unimplemented")
	// Cek apakah user sudah login atau belum
	if userId == 0 {
		return errors.New("user not logged in")
	}

	if input.Status == "" {
		return errors.New("Status cannot be empty, Status not updated")
	}

	if input.Status != "Completed" && input.Status != "Not Completed" {
		return errors.New("incorrect input, can only accept complete and incomplete input")
	}
	err := s.taskData.Status(id, input, userId)
	return err

}
