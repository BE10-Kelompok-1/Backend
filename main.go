package main

import (
	"backend/config"
	"backend/factory"
	"backend/infrastructure/database/mysql"
	"fmt"

	"github.com/labstack/echo/v4"
)

func main() {
	cfg := config.GetConfig()
	db := mysql.InitDB(cfg)
	mysql.MigrateData(db)
	e := echo.New()

	factory.InitFactory(e, db)

	fmt.Println("Starting programm ...")
	dsn := fmt.Sprintf(":%d", config.SERVERPORT)
	e.Logger.Fatal(e.Start(dsn))
}
