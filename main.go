package main

import (
	"rahmat/to-do-list-app/app/config"
	"rahmat/to-do-list-app/app/database"
	"rahmat/to-do-list-app/app/router"

	"github.com/labstack/echo/v4"
)

func main() {
	cfg := config.InitConfig()
	dbMysql := database.InitDbMysql(cfg)
	database.InitMigration(dbMysql)

	e := echo.New()
	router.InitRouter(dbMysql, e)

	e.Logger.Fatal(e.Start(":8000"))
}
