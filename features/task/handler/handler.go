package handler

import (
	"fmt"
	"net/http"
	"rahmat/to-do-list-app/app/middlewares"
	"rahmat/to-do-list-app/features/task"
	"rahmat/to-do-list-app/helper"
	"strconv"
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

func (hendler *TaskHandler) UpdateTask(c echo.Context) error {
	// Mendapatkan ID tugas dari parameter URL
	id := c.Param("task_id")
	idStr, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.WebResponse(http.StatusBadRequest, "error. id should be a number", nil))
	}

	// Mendapatkan ID pengguna dari token (asumsi ExtractTokenUserId bekerja dengan baik)
	userId := middlewares.ExtractTokenUserId(c)

	updateInput := new(UpdateTaskInput)
	err = c.Bind(&updateInput)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.WebResponse(http.StatusBadRequest, "error binding data", nil))
	}
	// Membuat objek CoreTask dari input yang diterima
	taskCore := task.CoreTask{

		Name:        updateInput.Name,
		UserId:      updateInput.UserId,
		Description: updateInput.Description,
	}

	// Memanggil metode Update dari service untuk memperbarui tugas
	err = hendler.taskService.Update(uint(idStr), taskCore, uint(userId))
	if err != nil {
		// Jika terjadi kesalahan saat memperbarui tugas, tangani kesalahan dan kembalikan respons yang sesuai
		return c.JSON(http.StatusInternalServerError, helper.WebResponse(http.StatusInternalServerError, "error updating task", nil))
	}

	// Jika tugas berhasil diperbarui, kembalikan respons sukses
	return c.JSON(http.StatusOK, helper.WebResponse(http.StatusOK, "task updated successfully", nil))
}

func (hendler *TaskHandler) StatusUpdate(c echo.Context) error {
	// Mendapatkan ID tugas dari parameter URL
	id := c.Param("task_id")
	idStr, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.WebResponse(http.StatusBadRequest, "error. id should be a number", nil))
	}
	// Mendapatkan ID pengguna dari token (asumsi ExtractTokenUserId bekerja dengan baik)
	userId := middlewares.ExtractTokenUserId(c)

	updateStatus := new(StatusUpdate)
	err = c.Bind(&updateStatus)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.WebResponse(http.StatusBadRequest, "error binding data", nil))
	}

	statusCore := task.CoreTask{

		Status: updateStatus.Status,
	}
	// Memanggil metode Update dari service untuk memperbarui tugas
	err = hendler.taskService.Status(uint(idStr), statusCore, uint(userId))
	fmt.Println(err)
	if err != nil {
		// Jika terjadi kesalahan saat memperbarui tugas, tangani kesalahan dan kembalikan respons yang sesuai
		return c.JSON(http.StatusInternalServerError, helper.WebResponse(http.StatusInternalServerError, "error updating status task", nil))
	}

	// Jika status tugas berhasil diperbarui, kembalikan respons sukses
	return c.JSON(http.StatusOK, helper.WebResponse(http.StatusOK, "task updated status successfully", nil))
}
