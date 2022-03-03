package config

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Config struct {
	App struct {
		Port string `env-default:"8080"`
		Host string `env-default:"localhost"`
	}
	Database struct {
		DB *sql.DB
	}
	NatsStreaming struct {
		Host string
		Port string
		Time time.Duration // yyyy:mm:dd
	}
}

var once sync.Once
var instance Config

//NewConfig - create config from env
func NewConfig() Config {
	once.Do(func() {
		err := godotenv.Load()
		if err != nil {
			log.Println(err.Error())
		}

		instance.NatsStreaming.Host = os.Getenv("HOST_NS")
		instance.NatsStreaming.Port = os.Getenv("PORT_NS")
		instance.App.Port = os.Getenv("PORT_APP")
		instance.App.Host = os.Getenv("HOST_APP")
		timeSub := os.Getenv("TIME_SUB")

		host := os.Getenv("HOST_DB")
		name := os.Getenv("NAME_DB")
		username := os.Getenv("USERNAME_DB")
		password := os.Getenv("PASSWD_DB")
		port := os.Getenv("PORT_DB")
		instance.Database.DB = connectToDB(port, host, username, password, name)
		instance.NatsStreaming.Time = time.Now().Sub(getTime(timeSub)) // NatsStreaming.Time - это время с которого надо начать принимать сообщения из натса (я предпочитаю твикс)
	})
	return instance
}

func connectToDB(portDB, hostDB, usernameDB, passwordDB, nameDB string) *sql.DB {
	port, err := strconv.Atoi(portDB)
	if err != nil {
		log.Fatal(err)
	}

	psqlconnect := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable",
		hostDB, port, usernameDB, passwordDB, nameDB)

	db, err := sql.Open("postgres", psqlconnect)
	if err != nil {
		log.Fatal("Cannot open database: ", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Connected to Postgresql!")

	return db
}

func getTime(timeSub string) time.Time {
	timeS := strings.Split(timeSub, ":")
	year, err := strconv.Atoi(timeS[0])
	if err != nil {
		return time.Now()
	}

	month, err := strconv.Atoi(timeS[0])
	if err != nil {
		return time.Now()
	}

	day, err := strconv.Atoi(timeS[0])
	if err != nil {
		return time.Now()
	}

	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}
