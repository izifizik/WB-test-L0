package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"sync"
)

type Config struct {
	App struct {
		Port string
		IP   string
	}
	Database struct {
		Port     string
		Host     string
		Username string
		Password string
		Name     string
	}
	NatsStreaming struct {
		NatsURL string
	}
}

var once sync.Once
var instance Config

func NewConfig() Config {
	once.Do(func() {
		err := godotenv.Load()
		if err != nil {
			log.Fatal(err.Error())
		}
		instance.App.Port = os.Getenv("PORT_APP")
		instance.App.IP = os.Getenv("IP_APP")
		instance.Database.Host = os.Getenv("HOST_DB")
		instance.Database.Name = os.Getenv("NAME_DB")
		instance.Database.Username = os.Getenv("USERNAME_DB")
		instance.Database.Password = os.Getenv("PASSWD_DB")
		instance.Database.Port = os.Getenv("PORT_DB")
		instance.NatsStreaming.NatsURL = os.Getenv("NATS_URL")
	})
	return instance
}
