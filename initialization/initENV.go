package initialization

import (
	"github.com/joho/godotenv"
	"log"
)

func InitEnv() error {
	if err := godotenv.Load(); err != nil {
		log.Println(err)
		return err
	}
	return nil
}
