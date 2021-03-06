// Package database contains functions for working with the database
package database

import (
	"currency/config"
	"currency/errorhandler"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

var db1 *sql.DB = nil

// ConnectDB function connects the user to the database
func ConnectDB() error {

	err := initDB()
	if err != nil {
		log.Println(err)
		return err
	}

	s := "host=" + "localhost" +
		" port=" + "5432" +
		" user=" + config.Cfg.DatabaseUser +
		" password=" + config.Cfg.DatabasePassword +
		" dbname=" + config.Cfg.DATABASE +
		" sslmode=disable"

	db, err1 := sql.Open("postgres", s)
	if err1 != nil {
		log.Println(err1)
		return err1
	}

	db1 = db
	fmt.Println("connected database user!")

	return nil
}

// initDB function initializes the database for user (create user, table, privileges)
func initDB() error {

	s := "host=" + "localhost" +
		" port=" + "5432" +
		" user=" + config.Cfg.DatabaseDefaultUser +
		" password=" + config.Cfg.DatabaseDefaultPassword +
		" dbname=" + config.Cfg.DATABASE +
		" sslmode=disable"

	db, err := sql.Open("postgres", s)
	if err != nil {
		log.Println(err)
		return err
	}

	fmt.Println("connected database default!")

	logTable := "logTable"

	s = "CREATE TABLE " + config.Cfg.Table + " (id integer PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY, data varchar(100) NOT NULL, dateInsert timestamp(0) without time zone);"
	_, err = db.Exec(s)
	errorhandler.CheckError(err)
	s = "CREATE TABLE " + logTable + " (id integer PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY, data varchar(300), typeRequest varchar(100), dateInsert timestamp(0) without time zone);"
	_, err = db.Exec(s)
	errorhandler.CheckError(err)
	s = "CREATE USER " + config.Cfg.DatabaseUser + " WITH PASSWORD '" + config.Cfg.DatabasePassword + "';"
	_, err = db.Exec(s)
	errorhandler.CheckError(err)
	s = "GRANT ALL PRIVILEGES ON DATABASE " + config.Cfg.DATABASE + " to " + config.Cfg.DatabaseUser + ";"
	_, err = db.Exec(s)
	errorhandler.CheckError(err)
	s = "GRANT All PRIVILEGES ON TABLE " + config.Cfg.Table + ", " + logTable + " TO " + config.Cfg.DatabaseUser + ";"
	_, err = db.Exec(s)
	errorhandler.CheckError(err)

	fmt.Println("create Table and user")
	err = db.Close()
	errorhandler.CheckError(err)
	fmt.Println("base default close")

	return nil
}
