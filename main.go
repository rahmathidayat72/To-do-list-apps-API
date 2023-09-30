package main

import (
	"rahmat/to-do-list-app/app/config"
	"rahmat/to-do-list-app/app/database"
	"rahmat/to-do-list-app/app/router"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	cfg := config.InitConfig()
	dbMysql := database.InitDbMysql(cfg)
	database.InitMigration(dbMysql)

	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	e.Use(middleware.CORS())

	e.Pre(middleware.RemoveTrailingSlash())
	router.InitRouter(dbMysql, e)

	e.Logger.Fatal(e.Start(":8000"))
}
