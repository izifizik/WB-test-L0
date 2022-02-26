package main

import (
	"WB-test-L0/internal/app"
	"log"
)

func main() {
	err := app.Run()
	if err != nil {
		log.Fatal(err.Error())
	}
}
