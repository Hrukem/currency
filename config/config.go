// Package config for configures program
package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type cfg struct {
	ApiKey                  string
	Website                 string
	Port                    string
	Table                   string
	DatabaseUser            string
	DatabasePassword        string
	DATABASE                string
	DatabaseDefaultPassword string
	DatabaseDefaultUser     string
}

var Cfg cfg

// Config function configures the server
func Config() error {

	err := godotenv.Load()
	if err != nil {
		log.Println("Error config!", err)
		return err
	}

	Cfg.ApiKey = os.Getenv("API_KEY")
	Cfg.Website = os.Getenv("WEBSITE")
	Cfg.Port = os.Getenv("PORT")
	Cfg.Table = os.Getenv("TABLE")
	Cfg.DatabaseUser = os.Getenv("DATABASE_USER")
	Cfg.DatabasePassword = os.Getenv("DATABASE_PASSWORD")
	Cfg.DATABASE = os.Getenv("DATABASE")
	Cfg.DatabaseDefaultPassword = os.Getenv("DATABASE_DEFAULT_PASSWORD")
	Cfg.DatabaseDefaultUser = os.Getenv("DATABASE_DEFAULT_USER")

	return nil
}
