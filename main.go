package main

import (
	"backend/config"
	"backend/factory"
	awss3 "backend/infrastructure/database/aws"
	"backend/infrastructure/database/mysql"
	"fmt"

	"github.com/labstack/echo/v4"
)

func main() {
	cfg := config.GetConfig()
	db := mysql.InitDB(cfg)
	mysql.MigrateData(db)
	e := echo.New()

	awsConn := awss3.InitS3(cfg.Keys3, cfg.Secrets3, cfg.Regions3)
	factory.InitFactory(e, db, awsConn)

	fmt.Println("Starting programm ...")
	dsn := fmt.Sprintf(":%d", config.SERVERPORT)
	e.Logger.Fatal(e.Start(dsn))
}
