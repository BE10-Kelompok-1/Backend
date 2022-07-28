package config

import (
	"log"
	"os"
	"strconv"
	"sync"
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
	// 	"github.com/joho/godotenv"
=======

// 	"github.com/joho/godotenv"
>>>>>>> 49057f1 (Update config.go)
=======
	// 	"github.com/joho/godotenv"
>>>>>>> 2e06ba8 (fix conflict)
=======
	// "github.com/joho/godotenv"
>>>>>>> f5449e8 (update semua fitur upload foto)
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
<<<<<<< HEAD
<<<<<<< HEAD
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
=======
	// 	err := godotenv.Load("local.env")

	// 	if err != nil {
	// 		log.Fatal("Cannot read configuration")
	// 		return nil
	// 	}
>>>>>>> 2e06ba8 (fix conflict)
=======
	// err := godotenv.Load("local.env")

	// if err != nil {
	// 	log.Fatal("Cannot read configuration")
	// 	return nil
	// }
>>>>>>> f5449e8 (update semua fitur upload foto)
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
=======
	cnv, err = strconv.Atoi(os.Getenv("DB_PORT"))
>>>>>>> 2e06ba8 (fix conflict)
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
