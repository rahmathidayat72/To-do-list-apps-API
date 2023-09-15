package main

import (
	"rahmat/to-do-list-app/app/config"
	"rahmat/to-do-list-app/app/database"
	"rahmat/to-do-list-app/app/router"

	"github.com/labstack/echo/v4"
)

func main() {
	cfg := config.InitConfig()
	database.InitDbMysql(cfg)

	e := echo.New()
	router.InitRouter(e)

	e.Logger.Fatal(e.Start(":8000"))
}
