// Package cron performs one task - to regularly receive data on the exchange rate
// and put this data in the database
package cron

import (
	"currency/database"
	"currency/errorhandler"
	"currency/request"
	"fmt"
	"github.com/robfig/cron/v3"
	"time"
)

// UserJob function start the task
func UserJob() {
	c := cron.New()
//	_, err := c.AddFunc("1 12 * * *", func() {insertDatabase()})
	_, err := c.AddFunc("*/1 * * * *", func() { insertDatabase() })
	errorhandler.CheckError(err)
	fmt.Println("Cron start!")
	c.Start()
}

// insertDatabase function receive data from site and insert data in database
func insertDatabase() {
	answer, err := request.Request(
		time.Now().Format("2006-1-2"),
		"JPY,CHF,EUR,GBP,RUB",
		"/live",
		)
	errorhandler.CheckError(err)

	data := map[string][]string{"Currency": answer}
	for _, v := range data["Currency"] {
		database.Insert(v, time.Now())
	}
	fmt.Println("Cron job!", time.Now())
}
