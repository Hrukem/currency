// Package handlers contain function for handler URL
package handlers

import (
	"currency/database"
	"currency/errorhandler"
	"currency/request"
	"html/template"
	"log"
	"net/http"
	"time"
)

// Home function open home page (use function outputHTML() )
func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	} else {
		outputHTML(w, "./web/templates/home.html", nil)
	}
}

// InputDay function open page for input date
// to view the exchange rate on the specified date
func InputDay(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/inputday/" {
		outputHTML(w, "./web/templates/home.html", nil)
	} else {
		outputHTML(w, "./web/templates/inputday.html", nil)
	}
}

// ShowCurrencyDay function receives data from request
// and displays this data on the browser page (use function outputHTML(w, f, d) )
func ShowCurrencyDay(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/showcurrencyday/" {
		outputHTML(w, "./web/templates/home.html", nil)
	} else {
		if r.Method == "POST" {
			err := r.ParseForm()
			errorhandler.CheckError(err)

			date := r.FormValue("day")
			dateTime, err1 := time.Parse("2006-01-02", date)
			errorhandler.CheckError(err1)

			// if date from user > time.Now() open page with alert
			// else send request
			if dateTime.After(time.Now()) {
				alert := map[string]string{"Alert": "You enter wrong date"}
				outputHTML(w, "./web/templates/inputday.html", alert)
			} else {
				currency := r.FormValue("currency")
				answer, err2 := request.Request(date, currency, "/historical")
				if err2 != nil {
					log.Println("Error in request.Request(w, r)", time.Now())
					alert := map[string]string{"Alert": "No connection, try again"}
					outputHTML(w, "./web/templates/inputday.html", alert)
				} else {
					data := map[string]interface{}{"Currency": answer}
					outputHTML(w, "./web/templates/showcurrencyday.html", data)
				}
			}
		} else {
			outputHTML(w, "./web/templates/home.html", nil)
		}
	}
}

// InputPeriod function open page for input two dates
// to view the exchange rate on the specified period
func InputPeriod(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/inputperiod/" {
		outputHTML(w, "./web/templates/home.html", nil)
	} else {
		outputHTML(w, "./web/templates/inputperiod.html", nil)
	}
}

// ShowCurrencyPeriod function receives data from request
// and displays this data on the browser page (use function outputHTML(w, f, d) )
func ShowCurrencyPeriod(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/showcurrencyperiod/" {
		outputHTML(w, "./web/templates/home.html", nil)
	} else {
		if r.Method == "POST" {
			err := r.ParseForm()
			errorhandler.CheckError(err)

			dateStart := r.FormValue("daystart")
			dateEnd := r.FormValue("dayend")
			dateStartTime, err1 := time.Parse("2006-01-02", dateStart)
			errorhandler.CheckError(err1)
			dateEndTime, err2 := time.Parse("2006-01-02", dateEnd)
			errorhandler.CheckError(err2)

			// if dateStart from user > dateEnd or dateEnd > time.Now() open page with alert
			// else take data from database
			if dateStartTime.After(dateEndTime) || dateEndTime.After(time.Now()) {
				alert := map[string]string{"Alert": "You enter wrong date"}
				outputHTML(w, "./web/templates/inputperiod.html", alert)
			} else {
				currency := r.FormValue("currency")

				currencySS := database.GetDatabase(dateStart, dateEnd, currency)

				data := map[string]interface{}{"Currency": currencySS}
				outputHTML(w, "./web/templates/showcurrencyperiod.html", data)
			}
		} else {
			outputHTML(w, "./web/templates/home.html", nil)
		}
	}
}

// outputHTML function displays data on the browser page
func outputHTML(w http.ResponseWriter, filename string, data interface{}) {
	t, err := template.ParseFiles(filename)
	if err != nil {
		log.Println("Error in outputHTML(w, f, d) in parse file", filename)
		http.Error(w, err.Error(), 500)
		return
	}

	err = t.Execute(w, data)
	if err != nil {
		log.Println("Error in outputHTML(w, f, d) in t.Execute(w, d)")
		http.Error(w, err.Error(), 500)
		return
	}
}
