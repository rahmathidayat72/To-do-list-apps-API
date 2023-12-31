package router

import (
	"rahmat/to-do-list-app/app/middlewares"
	taskdata "rahmat/to-do-list-app/features/task/data"
	taskhendler "rahmat/to-do-list-app/features/task/handler"
	taskservice "rahmat/to-do-list-app/features/task/service"
	userdata "rahmat/to-do-list-app/features/user/data"
	userhandler "rahmat/to-do-list-app/features/user/handler"
	userservice "rahmat/to-do-list-app/features/user/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, e *echo.Echo) {
	//testing andpoint

	dataUser := userdata.NewDataUser(db)
	userService := userservice.NewServiceUser(dataUser)
	userHandlerApi := userhandler.NewHandlerUser(userService)

	dataTask := taskdata.New(db)
	taskService := taskservice.NewTaskService(dataTask)
	taskHandlerApi := taskhendler.NewHandlerTask(taskService)

	// e.GET("/home", func(c echo.Context) error {
	// 	return c.JSON(http.StatusOK, map[string]any{
	// 		"messages": "Hello,World!",
	// 	})
	// })
	e.POST("/login", userHandlerApi.Login)
	e.GET("/users", userHandlerApi.GetAllUser)
	e.POST("/users", userHandlerApi.CreatedUser)
	e.PUT("/users/:user_id", userHandlerApi.UpdateUser, middlewares.JWTMiddleware())
	e.GET("/users/:user_id", userHandlerApi.GetUserById, middlewares.JWTMiddleware())
	e.DELETE("/users/:user_id", userHandlerApi.DeleteUser, middlewares.JWTMiddleware())

	e.GET("/tasks", taskHandlerApi.GetAllTask, middlewares.JWTMiddleware())
	e.POST("/tasks", taskHandlerApi.CreateTask, middlewares.JWTMiddleware())
	e.PUT("/tasks/:task_id", taskHandlerApi.UpdateTask, middlewares.JWTMiddleware())
	e.PUT("/status/:task_id", taskHandlerApi.StatusUpdate, middlewares.JWTMiddleware())
	e.DELETE("/tasks/:task_id", taskHandlerApi.Delete, middlewares.JWTMiddleware())

}
