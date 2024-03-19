package main

import (
	"github.com/Dubjay18/green-lit/app"
	"github.com/Dubjay18/green-lit/logger"
)

func main() {
	//log.Println("Starting the application...")
	app.GetEnvVar()
	logger.Info("Starting the application...")
	app.SanityCheck()
	app.Start()
}
