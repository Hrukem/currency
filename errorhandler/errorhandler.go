package errorhandler

import (
	"log"
	"net/http"
)

func CheckError(err error) {
	if err != nil {
		log.Println(err)
	}
}

func CheckErrorBrowser(err error, w http.ResponseWriter) {
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func CheckErrorURL(s string, w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != s {
		http.NotFound(w, r)
		return
	}
}
