package service

import "rahmat/to-do-list-app/features/task"

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
