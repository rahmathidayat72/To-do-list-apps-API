package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func InitRouter(e *echo.Echo) {
	//testing andpoint
	e.GET("/home", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]any{
			"messages": "Hello,World!",
		})
	})
}
