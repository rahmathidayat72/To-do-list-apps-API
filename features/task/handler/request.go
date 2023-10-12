package handler

import "rahmat/to-do-list-app/features/task"

type TaskRequest struct {
	Name        string `json:"name" form:"name"`
	UserId      uint   `json:"user_id" form:"user_id"`
	Description string `json:"description" form:"description"`
	Status      string `json:"status" form:"status"`
}

type UpdateTaskInput struct {
	Name        string `json:"name" form:"name"`
	UserId      uint   `json:"user_id" form:"user_id"`
	Description string `json:"description" form:"description"`
}

type StatusUpdate struct {
	Status string `json:"status" form:"status"`
}

func ModelToCore(r TaskRequest) task.CoreTask {
	return task.CoreTask{

		Name:        r.Name,
		UserId:      r.UserId,
		Description: r.Description,
		Status:      r.Status,
	}
}
