package handler

import "time"

type TaskResponse struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	UserId      uint      `json:"user_id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
}
