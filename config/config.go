package config

import (
	"log"
	"os"
	"strconv"
	"sync"
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
	// err := godotenv.Load("local.env")
	//!! USE "source local.env"!!!!!!!!!!!!!!!
	// if err != nil {
	// 	log.Fatal("Cannot read configuration")
	// 	return nil
	// }
	SECRET = os.Getenv("SECRET")
	cnv, err := strconv.Atoi("80")
	if err != nil {
		log.Fatal("Cannot parse port variable")
		return nil
	}

	SERVERPORT = int16(cnv)
	defaultConfig.Name = os.Getenv("NAME")
	defaultConfig.Username = os.Getenv("USERNAME")
	defaultConfig.Password = os.Getenv("PASSWORD")
	defaultConfig.Address = os.Getenv("ADDRESS")
	cnv, err = strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatal("Cannot parse DB Port variable")
		return nil
	}

	defaultConfig.Port = cnv

	return &defaultConfig
}
