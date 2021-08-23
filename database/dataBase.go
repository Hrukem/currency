// Package database contains functions for working with the database
package database

import (
	"currency/errorhandler"
	"strings"
	"time"
)

var currency = "currency"

// Insert function puts the data in the database
func Insert(data string, dateInsert time.Time) {
	query := "INSERT INTO " + currency + " (data, dateInsert) values ($1, $2)"
	_, err := db1.Exec(query, data, dateInsert)
	errorhandler.CheckError(err)
}

// InsertLogAPI function places data about the API request in the log table: currency names, request type, request time
func InsertLogAPI(data string, typeRequest string, dateInsert time.Time) {
	logTable := "logTable"
	query := "INSERT INTO " + logTable + " (data, typeRequest, dateInsert) values ($1, $2, $3)"
	_, err := db1.Exec(query, data, typeRequest, dateInsert)
	errorhandler.CheckError(err)
}

// GetDatabase function takes the saved data on the exchange rate for the selected period
func GetDatabase(start string, end string, s string ) []string {
	var currencySS []string

	s = strings.ReplaceAll(s, " ", "")
	s = strings.ToUpper(s)
	ss := strings.Split(s, ",")
	for _, curr := range ss {
		currLike := "%" + curr + "%"
		query := "SELECT data, dateInsert FROM " + currency + " WHERE dateinsert > $1 AND dateinsert < $2 AND data LIKE $3"
		rows, err := db1.Query(query, start+" 00:00:00", end+" 23:59:59", currLike)
		errorhandler.CheckError(err)
		for rows.Next() {
			var data string
			var dateInsert time.Time

			err1 := rows.Scan(&data, &dateInsert)
			errorhandler.CheckError(err1)
			currencySS = append(currencySS, data + " " + dateInsert.Format("2006-01-02_15:04:05"))
		}
	}
	return currencySS
}