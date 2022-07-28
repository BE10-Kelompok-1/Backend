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

	awsConn := awss3.InitS3("AKIA4OMEXK63GWCQQ4HC", "TDR4qqDP2OwjxNiLcRCrc1pVxuZK4G5DmTEA203P", "ap-southeast-1")
	factory.InitFactory(e, db, awsConn)

	fmt.Println("Starting programm ...")
	dsn := fmt.Sprintf(":%d", config.SERVERPORT)
	e.Logger.Fatal(e.Start(dsn))
}
