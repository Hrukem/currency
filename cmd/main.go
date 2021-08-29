package main

import (
	"currency/config"
	"currency/cron"
	"currency/database"
	"currency/server"
	"log"
	"time"
)

func main() {
	err := config.Config()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	err = database.ConnectDB()
	if err != nil {
		log.Fatal("Error initializations database")
	}
	time.Sleep(1 * time.Second)

	cron.UserJob()

	serveruser.Server()
}
