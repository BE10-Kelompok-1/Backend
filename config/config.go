package config

import (
	"log"
	"os"
	"strconv"
	"sync"
<<<<<<< HEAD
	// 	"github.com/joho/godotenv"
=======

// 	"github.com/joho/godotenv"
>>>>>>> 49057f1 (Update config.go)
)

type AppConfig struct {
	Driver   string
	Name     string
	Address  string
	Port     int
	Username string
	Password string
}

var lock = &sync.Mutex{}
var appConfig *AppConfig

func GetConfig() *AppConfig {
	lock.Lock()
	defer lock.Unlock()

	if appConfig == nil {
		appConfig = initConfig()
	}

	return appConfig
}

func initConfig() *AppConfig {
	var defaultConfig AppConfig
<<<<<<< HEAD
	// 	err := godotenv.Load("local.env")

	// 	if err != nil {
	// 		log.Fatal("Cannot read configuration")
	// 		return nil
	// 	}
=======
// 	err := godotenv.Load("local.env")

// 	if err != nil {
// 		log.Fatal("Cannot read configuration")
// 		return nil
// 	}
>>>>>>> 49057f1 (Update config.go)
	SECRET = os.Getenv("SECRET")
	cnv, err := strconv.Atoi(os.Getenv("SERVERPORT"))
	if err != nil {
		log.Fatal("Cannot parse port variable")
		return nil
	}

	SERVERPORT = int16(cnv)
	defaultConfig.Name = os.Getenv("Name")
	defaultConfig.Username = os.Getenv("DB_USERNAME")
	defaultConfig.Password = os.Getenv("DB_PASSWORD")
	defaultConfig.Address = os.Getenv("Address")
<<<<<<< HEAD
	cnv, err = strconv.Atoi(os.Getenv("DB_PORT"))
=======
<<<<<<< HEAD
	cnv, err = strconv.Atoi(os.Getenv("DB_Port"))

	// SERVERPORT = int16(cnv)
	// defaultConfig.Name = os.Getenv("NAME")
	// defaultConfig.Username = os.Getenv("USERNAME")
	// defaultConfig.Password = os.Getenv("PASSWORD")
	// defaultConfig.Address = os.Getenv("ADDRESS")
	// cnv, err = strconv.Atoi(os.Getenv("DB_PORT"))

=======
	cnv, err = strconv.Atoi(os.Getenv("DB_PORT"))
>>>>>>> b1af7c8 (unfinished posttesting)
>>>>>>> aef8658 (unfinished posttesting)
	if err != nil {
		log.Fatal("Cannot parse DB Port variable")
		return nil
	}

	defaultConfig.Port = cnv

	return &defaultConfig
}
