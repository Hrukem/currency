// Package serveruser contain function for work server
package serveruser

import (
	"currency/config"
	"currency/handlers"
	"log"
	"net/http"
)

// Server function launch server on specified port
func Server() {
	port := config.Cfg.Port

	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.Home)
	mux.HandleFunc("/inputday/", handlers.InputDay)
	mux.HandleFunc("/showcurrencyday/", handlers.ShowCurrencyDay)
	mux.HandleFunc("/inputperiod/", handlers.InputPeriod)
	mux.HandleFunc("/showcurrencyperiod/", handlers.ShowCurrencyPeriod)

	log.Println("запуск сервера на " + port)
	err := http.ListenAndServe(port, mux)
	log.Fatal(err)
}
