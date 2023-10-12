package service

import (
	"errors"
	"log"
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
		return errors.New("error: Status cannot be empty; Status not updated")
	}

	if input.Status != "Completed" && input.Status != "Not Completed" {
		return errors.New("error: Incorrect input; Only 'Completed' and 'Not Completed' are accepted")
	}
	err := s.taskData.Status(id, input, userId)
	if err != nil {
		return errors.New("Error: Failed to update task status - " + err.Error())
	}

	return nil

}

// Delete implements task.ServiceTaskInterface.
func (s *TaskService) Delete(Id uint, userId uint) error {
	// panic("unimplemented")
	// Cek apakah user sudah login atau belum
	if userId == 0 {
		return errors.New("user not logged in")
	}
	if Id == 0 {
		return errors.New("validation error. invalid id")
	}
	err := s.taskData.Delete(Id, userId)
	if err != nil {
		// Tambahkan pesan kesalahan yang lebih informatif jika penghapusan gagal
		return errors.New("failed to delete task: " + err.Error())
	}
	// Logging penghapusan tugas yang berhasil
	log.Printf("Task with Id %d deleted by user %d", Id, userId)

	return nil
}
