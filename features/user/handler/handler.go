package handler

import (
	"fmt"
	"net/http"
	"rahmat/to-do-list-app/app/middlewares"
	"rahmat/to-do-list-app/features/task"
	"rahmat/to-do-list-app/features/user"
	"rahmat/to-do-list-app/helper"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService user.ServiceUserInterface
}

func NewHandlerUser(service user.ServiceUserInterface) *UserHandler {
	return &UserHandler{
		userService: service,
	}
}

func (handler *UserHandler) Login(c echo.Context) error {
	inputLogin := new(LoginRequest)
	if err := c.Bind(inputLogin); err != nil {
		return c.JSON(http.StatusBadRequest, helper.WebResponse(http.StatusBadRequest, "error bind data, data not valid", nil))
	}

	login, token, err := handler.userService.Login(inputLogin.Email, inputLogin.Password)
	if err != nil {
		if strings.Contains(err.Error(), "invalid password") {
			return c.JSON(http.StatusUnauthorized, helper.WebResponse(http.StatusUnauthorized, "Invalid credentials", nil))
		}
		return c.JSON(http.StatusInternalServerError, helper.WebResponse(http.StatusInternalServerError, "Internal Server Error", nil))

	}
	// responseData := map[string]any{
	// 	"token": token,
	// 	"user":  login,
	// }
	var response = LoginResponse{
		ID:          login.ID,
		Name:        login.Name,
		Email:       login.Email,
		Address:     login.Address,
		PhoneNumber: login.PhoneNumber,
		Token:       token,
	}
	return c.JSON(http.StatusOK, helper.WebResponse(http.StatusOK, "success login", response))
}

func (handler *UserHandler) GetAllUser(c echo.Context) error {
	// middlewares.ExtractTokenUserId(c)
	result, err := handler.userService.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.WebResponse(http.StatusInternalServerError, "Error", nil))
	}
	var usersResponse []UserResponse
	//mapping dari truck core ke response
	for _, value := range result {
		usersResponse = append(usersResponse, UserResponse{
			ID:          value.ID,
			Name:        value.Name,
			Email:       value.Email,
			Address:     value.Address,
			PhoneNumber: value.PhoneNumber,
			CreatedAt:   value.CreatedAt,
			Task:        []task.CoreTask{},
		})
	}
	return c.JSON(http.StatusOK, helper.WebResponse(http.StatusOK, "seccess get All user", usersResponse))
}

func (handler *UserHandler) CreatedUser(c echo.Context) error {
	userInput := new(UserRequest)

	err := c.Bind(&userInput)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.WebResponse(http.StatusBadRequest, "error bind data , data not valid", nil))

	}

	//mapping dari struct request to struct core

	userCore := RequestToCore(*userInput)

	err = handler.userService.CreateUser(userCore)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, helper.WebResponse(http.StatusBadRequest, err.Error(), nil))
		}
		return c.JSON(http.StatusInternalServerError, helper.WebResponse(http.StatusInternalServerError, "Error", nil))

	}
	return c.JSON(http.StatusCreated, helper.WebResponse(http.StatusCreated, "Seccess creted data", nil))

}

func (handler *UserHandler) UpdateUser(c echo.Context) error {
	id := c.Param("user_id")
	idStr, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.WebResponse(http.StatusBadRequest, "error. id should be a number", nil))
	}
	roleId := middlewares.ExtractTokenUserId(c)
	fmt.Println(roleId)
	if idStr != roleId {
		return c.JSON(http.StatusBadRequest, helper.WebResponse(http.StatusBadRequest, "is not your id", nil))
	}
	userUpdate := new(UserRequest)
	err = c.Bind(&userUpdate)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.WebResponse(http.StatusBadRequest, "error binding data", nil))
	}

	// userCore := user.CoreUser{

	// 	Name:        userUpdate.Name,
	// 	Email:       userUpdate.Email,
	// 	Password:    userUpdate.Password,
	// 	Address:     userUpdate.Address,
	// 	PhoneNumber: userUpdate.PhoneNumber,
	// }

	userCore := RequestToCore(*userUpdate)
	err = handler.userService.Update(userCore, uint(idStr))
	if err != nil {
		// mengecek ada inputan sudah sesuai
		if strings.Contains(err.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, helper.WebResponse(http.StatusBadRequest, err.Error(), nil))
		}
		return c.JSON(http.StatusInternalServerError, helper.WebResponse(http.StatusInternalServerError, "Error", nil))

	}
	return c.JSON(http.StatusOK, helper.WebResponse(http.StatusOK, "User updated successfully", nil))

}

func (handler *UserHandler) GetUserById(c echo.Context) error {
	id := c.Param("user_id")
	idparam, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.WebResponse(http.StatusBadRequest, "error. id should be a number", nil))
	}
	user, err := handler.userService.SelectById(uint(idparam))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.WebResponse(http.StatusInternalServerError, "Error", nil))
	}
	userById := UserResponse{
		ID:          user.ID,
		Name:        user.Name,
		Email:       user.Email,
		Address:     user.Address,
		PhoneNumber: user.PhoneNumber,
		CreatedAt:   user.CreatedAt,
		Task:        []task.CoreTask{},
	}
	return c.JSON(http.StatusOK, helper.WebResponse(200, "success get user by id", userById))
}

func (handler *UserHandler) DeleteUser(c echo.Context) error {
	id := c.Param("user_id")
	idparam, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.WebResponse(http.StatusBadRequest, "error. id should be a number", nil))
	}
	roleId := middlewares.ExtractTokenUserId(c)
	if idparam != roleId {
		return c.JSON(http.StatusBadRequest, helper.WebResponse(http.StatusBadRequest, "is not your id", nil))

	}
	err = handler.userService.Delete(uint(idparam))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.WebResponse(http.StatusInternalServerError, "Error", nil))
	}
	return c.JSON(http.StatusOK, helper.WebResponse(http.StatusOK, "successfully delete user", nil))

}
