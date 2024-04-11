package main

import (
	"github.com/palle-404/erp-be/src/app"
	"github.com/palle-404/erp-be/src/config"
	"log"
)

func main() {
	if err := config.Load(); err != nil {
		log.Fatal("Failed to load app config. Error: " + err.Error())
	}
	if err := app.Start(); err != nil {
		log.Fatal("Failed to start app. Error: " + err.Error())
	}
	app.ListenForShutdown()

}
