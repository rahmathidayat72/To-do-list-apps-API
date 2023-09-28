package router

import (
	"net/http"
	"rahmat/to-do-list-app/features/user/data"
	"rahmat/to-do-list-app/features/user/handler"
	"rahmat/to-do-list-app/features/user/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, e *echo.Echo) {
	//testing andpoint

	dataUser := data.NewDataUser(db)
	userService := service.NewServiceUser(dataUser)
	userHandlerApi := handler.NewHandlerUser(userService)

	e.GET("/home", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]any{
			"messages": "Hello,World!",
		})
	})
	e.GET("/users", userHandlerApi.GetAllUser)
	e.POST("/users", userHandlerApi.CreatedUser)
	e.PUT("/users/:user_id", userHandlerApi.UpdateUser)
	e.GET("/users/:user_id", userHandlerApi.GetUserById)
	e.DELETE("/users/:user_id", userHandlerApi.DeleteUser)

}
