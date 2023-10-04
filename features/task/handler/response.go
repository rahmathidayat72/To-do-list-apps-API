package handler

import (
	"rahmat/to-do-list-app/features/task"
	"time"
)

type TaskResponse struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	UserId      uint      `json:"user_id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
}

func CoreToModel(core task.CoreTask) TaskResponse {
	return TaskResponse{
		ID:          core.ID,
		Name:        core.Name,
		UserId:      core.UserId,
		Description: core.Description,
		Status:      core.Status,
		CreatedAt:   core.CreatedAt,
	}
}