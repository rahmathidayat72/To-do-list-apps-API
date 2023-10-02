package handler

import (
	"net/http"
	"rahmat/to-do-list-app/app/middlewares"
	"rahmat/to-do-list-app/features/task"
	"rahmat/to-do-list-app/helper"

	"github.com/labstack/echo/v4"
)

type TaskHandler struct {
	taskService task.ServiceTaskInterface
}

func NewHandlerTask(service task.ServiceTaskInterface) *TaskHandler {
	return &TaskHandler{
		taskService: service,
	}
}

func (hendler *TaskHandler) GetAllProjeck(c echo.Context) error {
	userId := middlewares.ExtractTokenUserId(c)
	result, err := hendler.taskService.GetAll(uint(userId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.WebResponse(http.StatusInternalServerError, "error", nil))
	}

	var taskResponse []TaskResponse
	for _, v := range result {
		taskResponse = append(taskResponse, TaskResponse{
			ID:          v.ID,
			Name:        v.Name,
			UserId:      v.UserId,
			Description: v.Description,
			Status:      v.Status,
			CreatedAt:   v.CreatedAt,
		})

	}
	return c.JSON(http.StatusOK, helper.WebResponse(http.StatusOK, "Seccess read data project", taskResponse))

}
