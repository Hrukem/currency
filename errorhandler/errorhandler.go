// Package errorhandler
package errorhandler

import (
	"log"
)

// CheckError function prints an error on the screen
func CheckError(err error) {
	if err != nil {
		log.Println(err)
	}
}
