package config

import (
	"log"
	"os"
	"strconv"
	"sync"
	// 	"github.com/joho/godotenv"
)

type AppConfig struct {
	Driver   string
	Name     string
	Address  string
	Port     int
	Username string
	Password string
	Keys3    string
	Secrets3 string
	Regions3 string
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

	// if err != nil {
	// 	log.Fatal("Cannot read configuration")
	// 	return nil
	// }
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

	cnv, err = strconv.Atoi(os.Getenv("DB_PORT"))

	if err != nil {
		log.Fatal("Cannot parse DB Port variable")
		return nil
	}

	defaultConfig.Port = cnv
	defaultConfig.Keys3 = os.Getenv("Keys3")
	defaultConfig.Secrets3 = os.Getenv("Secrets3")
	defaultConfig.Regions3 = os.Getenv("Regions3")

	return &defaultConfig
}
