package handler

import (
	"net/http"
	"rahmat/to-do-list-app/app/middlewares"
	"rahmat/to-do-list-app/features/task"
	"rahmat/to-do-list-app/helper"
	"strings"

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

func (hendler *TaskHandler) GetAllTask(c echo.Context) error {
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

// func (hendler *TaskHandler) CreateTask(c echo.Context) error {
// 	taskInput := new(TaskRequest)
// 	err := c.Bind(&taskInput)
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, helper.WebResponse(http.StatusBadRequest, "error bind data", nil))
// 	}

// 	// Mengambil ID pengguna dari token JWT yang terkait dengan permintaan
// 	userID := middlewares.ExtractTokenUserId(c)
// 	fmt.Println(userID)
// 	input := ModelToCore(*taskInput)
// 	err = hendler.taskService.Create(input, input.ID)
// 	if err != nil {
// 		if strings.Contains(err.Error(), "validation") {
// 			return c.JSON(http.StatusBadRequest, helper.WebResponse(http.StatusBadRequest, err.Error(), nil))
// 		}

// 		return c.JSON(http.StatusInternalServerError, helper.WebResponse(http.StatusInternalServerError, "Error internal", nil))
// 	}
// 	return c.JSON(http.StatusOK, helper.WebResponse(200, "Seccess creted task", input))
// }

func (hendler *TaskHandler) CreateTask(c echo.Context) error {
	taskInput := new(TaskRequest)
	err := c.Bind(&taskInput)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.WebResponse(http.StatusBadRequest, "error binding data", nil))
	}

	// Mengambil ID pengguna dari token JWT yang terkait dengan permintaan
	userID := middlewares.ExtractTokenUserId(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, helper.WebResponse(http.StatusUnauthorized, "unauthorized", nil))
	}

	// Pastikan input.UserID sama dengan userID dari token JWT
	if taskInput.UserId != uint(userID) {
		return c.JSON(http.StatusForbidden, helper.WebResponse(http.StatusForbidden, "user does not have access to this project", nil))
	}

	taskCore := ModelToCore(*taskInput)

	err = hendler.taskService.Create(taskCore, uint(userID))
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, helper.WebResponse(http.StatusBadRequest, err.Error(), nil))
		}
		return c.JSON(http.StatusInternalServerError, helper.WebResponse(http.StatusInternalServerError, "internal server error", nil))
	}

	return c.JSON(http.StatusOK, helper.WebResponse(http.StatusOK, "Success created task", nil))
}
