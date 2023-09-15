package database

import (
	"fmt"
	"rahmat/to-do-list-app/app/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDbMysql(cfg *config.AppConfig) *gorm.DB {
	// var err error
	dbString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		cfg.DB_USERNAME, cfg.DB_PASSWORD, cfg.DB_HOSTNAME, cfg.DB_PORT, cfg.DB_NAME)
		
	db, err := gorm.Open(mysql.Open(dbString), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	fmt.Println("Conncetion to db")
	return db

}
