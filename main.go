package main

import (
	"github.com/Dubjay18/green-lit/app"
	"github.com/Dubjay18/green-lit/logger"
	"os"
)

func main() {
	//log.Println("Starting the application...")
	if os.Getenv("APP_ENV") != "production" {

		app.GetEnvVar()
	}
	logger.Info("Starting the application...")
	app.SanityCheck()
	app.Start()
}
