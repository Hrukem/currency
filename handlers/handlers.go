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
	errorhandler.CheckErrorURL("/", w, r)

	outputHTML(w, "./templates/home.html", nil)
}

func InputDay(w http.ResponseWriter, r *http.Request) {
	errorhandler.CheckErrorURL("/inputday/", w, r)
	outputHTML(w, "./templates/inputday.html", nil)
}

// ShowCurrencyDay function receives data from request
// and displays this data on the browser page (use function outputHTML(w, f, d) )
func ShowCurrencyDay(w http.ResponseWriter, r *http.Request) {
	errorhandler.CheckErrorURL("/showcurrencyday/", w, r)

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
			outputHTML(w,"./templates/inputday.html", alert)
		} else {
			currency := r.FormValue("currency")
			answer, err2 := request.Request(date, currency, "/historical")
			if err2  != nil {
				log.Println("Error in request.Request(w, r)", time.Now())
				alert := map[string]string{"Alert": "No connection, try again"}
				outputHTML(w, "./templates/inputday.html", alert)
			}else {
				data := map[string]interface{}{"Currency": answer}
				outputHTML(w, "./templates/showcurrencyday.html", data)
			}
		}
	}else{
		outputHTML(w, "./templates/home.html", nil)
	}
}

func InputPeriod(w http.ResponseWriter, r *http.Request) {
	errorhandler.CheckErrorURL("/inputperiod/", w, r)
	outputHTML(w, "./templates/inputperiod.html", nil)
}

func ShowCurrencyPeriod(w http.ResponseWriter, r *http.Request) {
	errorhandler.CheckErrorURL("/showcurrencyperiod/", w, r)

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
			outputHTML(w,"./templates/inputperiod.html", alert)
		} else {
			currency := r.FormValue("currency")

			currencySS := database.GetDatabase(dateStart, dateEnd, currency)

			data := map[string]interface{}{"Currency": currencySS}
			outputHTML(w, "./templates/showcurrencyperiod.html", data)
		}
	}else{
		outputHTML(w, "./templates/home.html", nil)
	}
}

// outputHTML function displays data on the browser page
func outputHTML(w http.ResponseWriter, filename string, data interface{}) {
	t, err := template.ParseFiles(filename)
	errorhandler.CheckErrorBrowser(err, w)

	err = t.Execute(w, data)
	errorhandler.CheckErrorBrowser(err, w)
}
