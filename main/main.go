package main

import (
	"currency/cron"
	"currency/database"
	"currency/initialization"
	"currency/server"
	"log"
	"time"
)

func main() {
	err := initialization.InitEnv()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	time.Sleep(500 * time.Millisecond)

	err = database.InitDataBase()
	if err != nil {
		log.Fatal("Error initializations database")
	}
	time.Sleep(1 * time.Second)

	cron.UserJob()

	serveruser.Server()
}
